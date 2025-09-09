package api

// ---------------- Request Models ----------------

// UserRegisterRequest 用户注册请求模型
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required" example:"admin" validate:"min=3,max=50"` // 用户名，长度3-50字符
	Password string `json:"password" binding:"required" example:"123456" validate:"min=6"`       // 密码，最少6位字符
}

// UserLoginRequest 用户登录请求模型
type UserLoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin" validate:"min=3,max=50"` // 用户名，长度3-50字符
	Password string `json:"password" binding:"required" example:"123456" validate:"min=6"`       // 密码，最少6位字符
}

// ---------------- Response Models ----------------

// User 用户信息响应模型
type User struct {
	Username string `json:"username" example:"admin"`       // 用户名
	Password string `json:"password,omitempty" example:"-"` // 密码（通常不返回）
	Id       int64  `json:"id" example:"1"`                 // 用户ID
}

// HealthResponse 健康检查响应模型
type HealthResponse struct {
	Status string `json:"status" example:"ok"` // 服务状态
}

// ---------------- Common Response Wrappers ----------------

// ApiResponse API统一响应结构（用于Swagger文档）
type ApiResponse struct {
	Data interface{} `json:"data,omitempty" swaggertype:"object"` // 响应数据
	Msg  string      `json:"msg" example:"success"`               // 响应消息
	Code int         `json:"code" example:"10000"`                // 响应代码，10000表示成功，10001表示失败
}

// SuccessResponse 成功响应示例（用于Swagger文档）
type SuccessResponse struct {
	Data interface{} `json:"data,omitempty"`        // 响应数据
	Msg  string      `json:"msg" example:"success"` // 响应消息
	Code int         `json:"code" example:"10000"`  // 响应代码
}

// ErrorResponse 错误响应示例（用于Swagger文档）
type ErrorResponse struct {
	Data interface{} `json:"data,omitempty"`       // 错误详情
	Msg  string      `json:"msg" example:"fail"`   // 错误消息
	Code int         `json:"code" example:"10001"` // 错误代码
}
