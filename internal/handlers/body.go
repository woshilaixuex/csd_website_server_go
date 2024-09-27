package handlers

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
	IsFresh       bool   `json:"is_fresh" `
	Orientation   string `json:"orientation"`
	Experience    string `json:"experience"`
}
type ApiRequest struct {
	Data interface{} `json:"data"`
}

// 返回体
type ApiRspBody string
type ApiResponse struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
