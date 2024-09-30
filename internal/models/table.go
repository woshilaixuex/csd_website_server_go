package models

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 表单模型
 * @Date: 2024-09-27 23:12
 */
var table_name = "enroll_table"

// 人员信息
type PersonalInfo struct {
	StudentNumber string `gorm:"column:student_number;not null"`             // 学号
	Name          string `gorm:"column:name;type:varchar(20);not null"`      // 姓名
	QQNumber      string `gorm:"column:qq_number;type:varchar(20);not null"` // QQ号
	Email         string `gorm:"column:email;type:varchar(50);not null"`     // 邮箱
	Reason        string `gorm:"column:reason;type:varchar(50);not null"`    // 申请理由
}

// 学习状况考察
type Inspection struct {
	Grade         int8   `gorm:"column:grade;not null;check:grade in (1, 2)"`
	HadExperience bool   `gorm:"column:had_experience;not null"`     // 是否为无开发经验
	Orientation   string `gorm:"column:orientation;type:char(4)"`    // 方向
	Experience    string `gorm:"column:experience;type:varchar(50)"` // 经验
}

// 提交表单
type EnrollTable struct {
	gorm.Model `gorm:"<-:create"` // 数据库默认模型
	PersonalInfo
	Inspection
}

func (EnrollTable) TableName() string {
	return table_name
}
func CreateEnrollTableIfNotExists(db *gorm.DB, cfgTableName string) error {
	table_name = cfgTableName
	// 检查表是否存在
	if !db.Migrator().HasTable(&EnrollTable{}) {
		logx.Info("table EnrollTable does not exist, creating...")

		// 使用 AutoMigrate 创建表
		if err := db.AutoMigrate(&EnrollTable{}); err != nil {
			return err
		}
		logx.Info("table EnrollTable created successfully.")
	} else {
		logx.Info("table EnrollTable already exists.")
	}
	return nil
}
func MustCreateEnrollTableIfNotExists(db *gorm.DB, cfgTableName string) {
	if err := CreateEnrollTableIfNotExists(db, cfgTableName); err != nil {
		panic(fmt.Sprintf("failed to create table EnrollTable: %v", err))
	}
}
