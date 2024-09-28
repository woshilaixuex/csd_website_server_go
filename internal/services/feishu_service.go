package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/csd-world/csd_webstie_server_go/app/config"
	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/csd-world/csd_webstie_server_go/pkg"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 飞书
 * @Date: 2024-09-28 03:33
 */

func FeiShuServiceLisen(cfg *config.Config, resultCh <-chan interface{}) {
	// 如果这个函数出异常不要影响返回（感觉其实可以删掉）
	defer func() {
		if r := recover(); r != nil {
			logx.Error("recovered panic", r)
		}
	}()
	for {
		time.Sleep(time.Second) // 每秒检查一次
		select {
		case data, ok := <-resultCh:
			if !ok {
				logx.Info("received close signal. Exiting...")
				return
			}
			if signal, ok := data.(string); ok { // 传输关闭信号后接收并关闭
				if signal == pkg.ClosedSignal {
					logx.Info("received close signal. Normal shutdown")
					return
				}
			}
			logx.Debug(data)
			if enroll, ok := data.(*models.EnrollTable); ok {
				PushEnroll(cfg, enroll)
			}
		default:
			continue
		}
	}
}
func PushEnroll(cfg *config.Config, enroll *models.EnrollTable) {
	if enroll == nil {
		return
	}
	token, _, err := fetchTenantAccessToken(cfg)
	if err != nil {
		return
	}
	var hadExperience string
	if enroll.IsFresh {
		hadExperience = "T"
	} else {
		hadExperience = "F"
	}
	err = createRecord(token, cfg, &models.RecordFields{
		CreatedAt:     int(time.Now().Unix() * 1000),
		StudentID:     enroll.StudentNumber,
		Name:          enroll.StudentNumber,
		QQ:            enroll.QQNumber,
		Email:         enroll.Email,
		HadExperience: hadExperience,
		ApplyReason:   enroll.Reason,
		Grade:         int(enroll.Grade),
		Experience:    enroll.Experience,
		Orientation:   enroll.Orientation,
	})
	if err != nil {
		logx.Error(err)
		return
	}
}

// FetchTenantAccessToken 获取租户访问令牌
func fetchTenantAccessToken(cfg *config.Config) (string, int, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

	// 创建请求体
	requestBody := models.TenantAccessTokenRequest{
		AppID:     cfg.FeiShuServer.AppId,
		AppSecret: cfg.FeiShuServer.AppSecret,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", 0, fmt.Errorf("请求体序列化错误: %w", err)
	}

	// 发送 POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", 0, fmt.Errorf("发送请求错误: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("读取响应错误: %w", err)
	}

	// 解析响应
	var response models.TenantAccessTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", 0, fmt.Errorf("响应体解析错误: %w", err)
	}

	// 检查返回代码
	if response.Code != 0 {
		return "", 0, fmt.Errorf("错误代码: %d, 消息: %s", response.Code, response.Msg)
	}

	return response.TenantAccessToken, response.Expire, nil
}

// CreateRecord 创建 Bitable 记录
func createRecord(userToken string, cfg *config.Config, fields *models.RecordFields) error {

	url := fmt.Sprintf("https://open.feishu.cn/open-apis/bitable/v1/apps/%s/tables/%s/records", cfg.FeiShuServer.AppToken, cfg.FeiShuServer.TableId)

	// 创建请求体
	requestBody := models.CreateRecordRequest{
		Fields: *fields,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("请求体序列化错误: %w", err)
	}

	// 创建请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求错误: %w", err)
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+userToken)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送请求错误: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取响应错误: %w", err)
		}
		logx.Info(body)
		return fmt.Errorf("请求失败，状态码: %d", resp.StatusCode)
	}

	fmt.Println("记录创建成功")
	return nil
}
