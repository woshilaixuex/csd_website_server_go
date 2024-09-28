package engine

import (
	"fmt"
	"strings"

	"github.com/csd-world/csd_webstie_server_go/app/config"
	"github.com/csd-world/csd_webstie_server_go/internal/models"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: 数据库初始化引擎(mysql)
 * @Date: 2024-09-27 23:44
 */
type MysqlEngine struct {
	*gorm.DB
}

func InitDB(cfg *config.Config) *MysqlEngine {
	BuilderDsn := func() string {
		var builder strings.Builder
		// 拼接 DSN 字符串
		builder.WriteString(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.PassWord,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.DatabaseName,
		))
		return builder.String()
	}
	dsn := BuilderDsn()
	var err error
	// 尝试连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		// 如果连接失败，检查错误类型
		if isDatabaseNotFoundError(err) {
			// 数据库不存在，尝试创建数据库
			if err := createDatabase(cfg); err != nil {
				panic(fmt.Sprintf("failed to create database: %v", err))
			}
			// 再次尝试连接数据库
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
			if err != nil {
				panic(fmt.Sprintf("database connect failed after creation: %v", err))
			}
		} else {
			panic(fmt.Sprintf("database connect failed: %v", err))
		}
	}
	models.MustCreateEnrollTableIfNotExists(db, cfg.Database.TableName)
	logx.Info("database is connected")
	return &MysqlEngine{
		db,
	}
}
func isDatabaseNotFoundError(err error) bool {
	// 根据具体的错误信息判断，下面的示例是针对 MySQL 的常见错误代码
	return strings.Contains(err.Error(), "1049") // MySQL's "Unknown database" error code
}
func createDatabase(cfg *config.Config) error {
	BuilderDsn := func() string {
		var builder strings.Builder
		// 拼接 DSN 字符串
		builder.WriteString(fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.Database.User,
			cfg.Database.PassWord,
			cfg.Database.Host,
			cfg.Database.Port,
		))
		return builder.String()
	}
	// 拼接 DSN 但不包括数据库名
	dsn := BuilderDsn()
	// 连接到 MySQL 服务器
	tempDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logx.Errorf("failed to connect to MySQL server: %w", err)
		return err
	}

	// 创建数据库
	if err := tempDb.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`;", cfg.Database.DatabaseName)).Error; err != nil {
		logx.Errorf("failed to connect to MySQL server: %w", err)
		return err
	}

	return nil
}
