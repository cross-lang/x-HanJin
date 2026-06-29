// Package services 实现业务逻辑层，
// 封装数据库操作和消息队列调用，为控制器提供业务接口。
package services

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"x-HanJin/internal/databases/mysql"
	"x-HanJin/internal/message_queues/rabbitmq"
	"x-HanJin/internal/message_queues/rabbitmq/producer"
	"x-HanJin/internal/models/user"
	userReq "x-HanJin/internal/models/user/request"
	"x-HanJin/pkg/log"
)

// UserService 用户业务逻辑服务
type UserService struct {
	DB *gorm.DB       // 数据库实例
	MQ *producer.Producer // 消息队列生产者
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{
		DB: mysql.MySQLdb,
		MQ: rabbitmq.RMQdp,
	}
}

// CreateUser 创建用户，写入数据库并发布消息通知
func (us *UserService) CreateUser(userData user.UserModel) (user.UserModel, error) {
	result := us.DB.Create(&userData)
	if result.Error != nil {
		log.Error("创建用户-数据库写入失败", zap.Error(result.Error))
		return userData, result.Error
	}

	// 发送 RabbitMQ 消息通知
	if us.MQ != nil {
		err := us.MQ.Publish(fmt.Sprintf("Create User Success, User ID: %d", userData.ID))
		if err != nil {
			log.Error("创建用户-消息发送失败", zap.Error(err))
			return userData, err
		}
	}

	return userData, nil
}

// GetAllUsers 获取所有用户列表
func (us *UserService) GetAllUsers() ([]user.UserModel, error) {
	var users []user.UserModel
	result := us.DB.Find(&users)
	if result.Error != nil {
		log.Error("获取用户列表失败", zap.Error(result.Error))
		return users, result.Error
	}
	return users, nil
}

// GetUserById 根据 ID 获取单个用户
func (us *UserService) GetUserById(userId uint) (*user.UserModel, error) {
	var userModel user.UserModel
	result := us.DB.Where("id = ?", userId).First(&userModel)
	if result.Error != nil {
		log.Error("获取用户信息失败", zap.Uint("id", userId), zap.Error(result.Error))
		return nil, result.Error
	}
	return &userModel, nil
}

// UpdateUserById 根据 ID 更新用户信息
func (us *UserService) UpdateUserById(id uint, updateUser userReq.UpdateUserRequest) (user.UserModel, error) {
	var userModel user.UserModel
	if err := us.DB.First(&userModel, id).Error; err != nil {
		log.Error("更新用户-查询失败", zap.Uint("id", id), zap.Error(err))
		return userModel, err
	}

	userModel.Email = updateUser.Email
	userModel.Name = updateUser.Name
	userModel.Gender = updateUser.Gender
	userModel.Age = updateUser.Age

	result := us.DB.Save(&userModel)
	if result.Error != nil {
		log.Error("更新用户-保存失败", zap.Uint("id", id), zap.Error(result.Error))
		return userModel, result.Error
	}
	return userModel, nil
}

// DeleteUserById 根据 ID 删除用户
func (us *UserService) DeleteUserById(id uint) (bool, error) {
	result := us.DB.Delete(&user.UserModel{}, id)
	if result.Error != nil {
		log.Error("删除用户失败", zap.Uint("id", id), zap.Error(result.Error))
		return false, result.Error
	}
	return true, nil
}
