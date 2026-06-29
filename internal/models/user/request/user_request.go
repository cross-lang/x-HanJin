// Package request 定义用户模块的请求参数 DTO（Data Transfer Object）。
package request

// CreateUserRequest 创建用户请求参数
type CreateUserRequest struct {
	Name   string `json:"name" example:"张三"`       // 用户姓名
	Age    uint   `json:"age" example:"18"`          // 用户年龄
	Email  string `json:"email" example:"test@test.com"` // 用户邮箱
	Gender string `json:"gender" example:"male"`     // 用户性别
}

// UpdateUserRequest 更新用户请求参数
type UpdateUserRequest struct {
	Name   string `json:"name" example:"张三"`       // 用户姓名
	Age    uint   `json:"age" example:"18"`          // 用户年龄
	Email  string `json:"email" example:"test@test.com"` // 用户邮箱
	Gender string `json:"gender" example:"male"`     // 用户性别
}
