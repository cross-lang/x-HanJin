// Package constants 定义应用级别的公共常量，
// 包括上下文 Key、应用标识、匹配模式等。
package constants

// 上下文 Key 常量，用于在 gin.Context 中传递请求元信息
const (
	CtxKeyXOrigin    = "X_ORIGIN"          // 请求来源信息 key
	CtxKeyXReqID     = "X_REQUEST_ID"      // 请求 ID key
	CtxKeyXReqRealIP = "X_REQUEST_REAL_IP" // 请求真实 IP key
)

// 应用基本属性
const (
	AppId   = "x-HanJin"      // 应用 ID
	AppName = "汉津"      // 应用名称
)

// 匹配模式常量
const (
	MatchModeExact = "exact" // 精确匹配
	MatchModeFuzzy = "fuzzy" // 模糊匹配
)
