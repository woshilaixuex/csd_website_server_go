package engine

import (
	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/zeromicro/go-zero/core/logx"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 数据库操作
 * @Date: 2024-09-28 02:45
 */
// 插入新记录
func (engine *MysqlEngine) InsertEnrollTable(enrollTable *models.EnrollTable) error {
	if err := engine.Create(enrollTable).Error; err != nil {
		logx.Error("failed to insert record:", err)
		return err
	}
	logx.Info("new record inserted successfully.")
	return nil
}

// 查询记录
func (engine *MysqlEngine) QueryEnrollTables() ([]models.EnrollTable, error) {
	var enrollTables []models.EnrollTable
	if err := engine.Find(&enrollTables).Error; err != nil {
		logx.Error("failed to query records:", err)
		return nil, err
	}
	logx.Info("records queried successfully.")
	return enrollTables, nil
}
