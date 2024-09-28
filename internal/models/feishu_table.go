package models

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description:
 * @Date: 2024-09-28 14:39
 */
// RecordFields 创建记录时使用的字段
type RecordFields struct {
	CreatedAt     int    `json:"createdAt"`
	StudentID     string `json:"studentId"`
	Name          string `json:"name"`
	QQ            string `json:"qq"`
	Email         string `json:"email"`
	HadExperience string `json:"hadExperience"`
	ApplyReason   string `json:"applyReason"`
	Grade         int    `json:"grade"`
	Experience    string `json:"experience"`
	Orientation   string `json:"orientation"`
}

// CreateRecordRequest 请求体
type CreateRecordRequest struct {
	Fields RecordFields `json:"fields"`
}
