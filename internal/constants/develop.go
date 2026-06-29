package constants

// HTTP 内容类型常量
const (
	JsonContentType              = "application/json"              // JSON
	FileContentType              = "application/octet-stream"      // 二进制流
	FormUrlEncodedContentType    = "application/x-www-form-urlencoded" // 表单
	MultipartFormDataContentType = "multipart/form-data"           // 多部分表单
)

// 加密算法名称常量
const (
	// 单向加密算法（哈希算法）
	MD5    = "MD5"
	SHA1   = "SHA1"
	SHA256 = "SHA256"
	SHA512 = "SHA512"
	SM3    = "SM3"

	// 对称加密算法
	AES      = "AES"
	SM4      = "SM4"
	DES      = "DES"
	ThreeDES = "3DES"
	ChaCha20 = "ChaCha20"
	RC4      = "RC4"

	// 非对称加密算法
	RSA = "RSA"
	ECC = "ECC"
	DSA = "DSA"
	SM2 = "SM2"
)

// 加密工作模式常量
const (
	ModeECB = "ECB" // 电子密码本模式
	ModeCBC = "CBC" // 密码分组链接模式
	ModeGCM = "GCM" // 伽罗瓦/计数器模式
)

// 填充方式常量
const (
	// 对称加密中的块加密填充方式
	PKCS7Padding    = "PKCS7"    // PKCS#7 填充
	ISO10126Padding = "ISO10126" // ISO 10126 填充
	NoPadding       = "NoPadding" // 无填充
	ZeroPadding     = "ZeroPadding" // 零填充

	// 非对称加密填充方式
	PKCS1v15 = "PKCS1v15" // PKCS#1 v1.5 填充
	OAEP     = "OAEP"     // 最优非对称加密填充
)

// 编码方式常量
const (
	EncodingBase64 = "base64" // Base64 编码
	EncodingHex    = "hex"    // 十六进制编码
)
