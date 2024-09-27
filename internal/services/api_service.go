package services

import "github.com/csd-world/csd_webstie_server_go/internal/models"

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: api服务层(数据操作)
 * @Date: 2024-09-28 03:26
 */
type ApiService struct {
	Engine *models.MysqlEngine
}

func (s *ApiService) InsertEnroll(enroll *models.EnrollTable) error {
	return s.Engine.InsertEnrollTable(enroll)
}

func (s *ApiService) QueryEnrolls() ([]models.EnrollTable, error) {
	return s.Engine.QueryEnrollTables()
}
