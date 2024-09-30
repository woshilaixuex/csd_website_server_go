package services

import (
	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/csd-world/csd_webstie_server_go/internal/models/engine"
	"github.com/csd-world/csd_webstie_server_go/pkg"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: api服务层(数据操作)
 * @Date: 2024-09-28 03:26
 */
type ApiService struct {
	Engine   *engine.MysqlEngine
	ResultCh *pkg.MssChan
}

func NewApiService(engine *engine.MysqlEngine, resultCh *pkg.MssChan) *ApiService {
	return &ApiService{
		Engine:   engine,
		ResultCh: resultCh,
	}
}
func (s *ApiService) InsertEnroll(enroll *models.EnrollTable) error {
	if s.ResultCh.CheckIsOpen() {
		s.ResultCh.Send(enroll)
	}
	return s.Engine.InsertEnrollTable(enroll)
}

func (s *ApiService) QueryEnrolls() ([]models.EnrollTable, error) {
	return s.Engine.QueryEnrollTables()
}
