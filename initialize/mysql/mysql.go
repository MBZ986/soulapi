package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"soulapi/global"
	"soulapi/models"
	"time"
)

func init() {
	sqldsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		global.GLOBAL_CONFIG.Mysql.Username,
		global.GLOBAL_CONFIG.Mysql.Password,
		global.GLOBAL_CONFIG.Mysql.Host,
		global.GLOBAL_CONFIG.Mysql.Port,
		global.GLOBAL_CONFIG.Mysql.Database)
	var err error

	if global.DB, err = gorm.Open(mysql.Open(sqldsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	}); err != nil {
		panic(err)
	}

	sqlDB, err := global.DB.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		global.Logger.Panicf("数据库连接失败：%v", err)
		return
	}

	global.Logger.Infof("数据库链接成功")

	if global.GLOBAL_CONFIG.Mysql.InitTables {
		initModels()
	}
}

func initModels() {
	global.Logger.Infof("初始化数据库表结构")
	err := global.DB.AutoMigrate(models.User{}, models.Title{}, models.Music{}, models.Video{}, models.Partition{}, models.Comment{},
		models.Label{}, models.Poster{}, models.Browsed{}, models.Like{}, models.LabelMedia{}, models.PartitionType{}, models.Collection{}, models.Message{})
	if err != nil {
		global.Logger.Errorf("初始化数据库表失败：%v", err)
		return
	}
	global.Logger.Infof("数据库表初始化成功")
}
