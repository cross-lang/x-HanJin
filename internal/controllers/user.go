// Package controllers 实现 HTTP 请求的控制器层，
// 负责参数绑定、请求校验和响应格式化。
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"x-HanJin/internal/models/user"
	userReq "x-HanJin/internal/models/user/request"
	"x-HanJin/internal/services"
	"x-HanJin/pkg/log"
)

// UserController 用户模块控制器，处理用户相关的 HTTP 请求
type UserController struct {
	userService *services.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body user.UserModel true "用户信息"
// @Success 201 {object} map[string]interface{} "创建成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /api/v1/users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var userModel user.UserModel
	if err := c.ShouldBindJSON(&userModel); err != nil {
		log.Warn("创建用户参数绑定失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := ctrl.userService.CreateUser(userModel)
	if err != nil {
		log.Error("创建用户失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Info("创建用户成功", zap.Uint("id", newUser.ID))
	c.JSON(http.StatusCreated, gin.H{"message": "ok", "data": newUser})
}

// GetAllUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 500 {object} map[string]interface{} "内部错误"
// @Router /api/v1/users [get]
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	users, err := ctrl.userService.GetAllUsers()
	if err != nil {
		log.Error("获取用户列表失败", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": users})
}

// GetUserById 根据 ID 获取用户信息
// @Summary 根据ID获取用户
// @Description 根据用户ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 404 {object} map[string]interface{} "用户不存在"
// @Router /api/v1/users/{id} [get]
func (ctrl *UserController) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warn("用户ID参数无效", zap.String("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	dstUser, err := ctrl.userService.GetUserById(uint(id))
	if err != nil {
		log.Error("获取用户信息失败", zap.Int("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": dstUser})
}

// UpdateUser 更新用户信息
// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body userReq.UpdateUserRequest true "用户信息"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 404 {object} map[string]interface{} "用户不存在"
// @Router /api/v1/users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warn("用户ID参数无效", zap.String("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var updateUserReq userReq.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserReq); err != nil {
		log.Warn("更新用户参数绑定失败", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := ctrl.userService.UpdateUserById(uint(id), updateUserReq)
	if err != nil {
		log.Error("更新用户失败", zap.Int("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Info("更新用户成功", zap.Int("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": newUser})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]interface{} "成功"
// @Failure 400 {object} map[string]interface{} "请求参数错误"
// @Failure 404 {object} map[string]interface{} "用户不存在"
// @Router /api/v1/users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Warn("用户ID参数无效", zap.String("id", c.Param("id")))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	isSuccess, err := ctrl.userService.DeleteUserById(uint(id))
	if err != nil {
		log.Error("删除用户失败", zap.Int("id", id), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	log.Info("删除用户成功", zap.Int("id", id))
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": isSuccess})
}
