package models

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: http返回和请求模型
 * @Date: 2024-09-28 03:58
 */
// 请求体
type ApiReqBody struct {
	StudentNumber string `json:"student_number" binding:"required"`
	Name          string `json:"name" binding:"required"`
	QQNumber      string `json:"qq_number" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Reason        string `json:"reason" binding:"required"`
	Grade         int8   `json:"grade" binding:"required"`
	HadExperience bool   `json:"had_experience" `
	Orientation   string `json:"orientation"`
	Experience    string `json:"experience"`
}
type ApiRequest struct {
	Data *ApiReqBody `json:"data"`
}

// 返回体
type ApiRespBody struct {
	Info string `json:"info"`
}
type ApiResponse struct {
	Code int32        `json:"code"`
	Msg  string       `json:"msg"`
	Data *ApiRespBody `json:"data"`
}

// TenantAccessTokenRequest 请求租户访问令牌的结构体
type TenantAccessTokenRequest struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

// TenantAccessTokenResponse 响应结构体
type TenantAccessTokenResponse struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}
