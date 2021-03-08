package handler
const (
	ErrCodeOK               = 0
	ErrCodeApiGatewayFormat = 1001
	ErrCodeUnknownAction    = 1002
	ErrCodeInvalidRequest   = 1003
	ErrCodeImportFail       = 1004
	ErrCodeReadFile         = 1005
	ErrCodeJsonMarshal      = 1006
	ErrCodeReadDB           = 1007
	ErrCodeCloudIdMissed    = 1008
	ErrCodeUploadFile       = 1009
	ErrCodeMd5Check         = 1010
	ErrCodeDataNotFound     = 1011
	ErrCodeUnsupportedFunc  = 1012
	ErrCodeShellCmd         = 1013
	ErrCodeMkdir            = 1014
	ErrCodeActivatePlan     = 1015
)

const (
	ErrMsgOK               = "OK"
	ErrMsgApiGatewayFormat = "API Gateway格式解析错误"
	ErrMsgUnknownAction    = "未知的Action"
	ErrMsgInvalidRequest   = "请求格式错误"
	ErrMsgImportFail       = "导入物料失败"
	ErrMsgReadFile         = "读取文件失败"
	ErrMsgJsonMarshal      = "JsonMarshal失败"
	ErrMsgReadDB           = "读取数据库失败"
	ErrMsgCloudIdMissed    = "获取不到CloudId信息"
	ErrMsgUploadFile       = "上传文件失败"
	ErrMsgMd5Check         = "Md5校验失败"
	ErrMsgDataNotFound     = "未查到对应数据"
	ErrMsgUnsupportedFunc  = "功能暂时不支持"
	ErrMsgShellCmd         = "执行CMD命令失败"
	ErrMsgMkDir            = "创建目录失败"
	ErrMsgActivatePlan     = "激活规划失败"
)
type ErrorMsg struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

//错误码返回体
func (e ErrorMsg) WithErrRsp(code int, msg string) interface{} {
	e.Code = code
	e.Message = msg
	return struct {
		Error ErrorMsg `json:Error`
	}{e}
}
