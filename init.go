package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	// DataHandler 数据库
	DataHandler *gorm.DB
	// err 订阅错误
	err error

	DeBugLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Microsecond, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,            // Disable color
		},
	)
)

// InitDataBase 初始化数据库
func InitDataBase(dsn string, p logger.Interface, debug bool) error {
	var err error
	DataHandler, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: p,
	})
	if err != nil {
		panic("failed to connect database")
	}

	if debug {
		// debug 模式下会创建表
		// 首次调用（可以使用debug模式运行后产生表后，切换为release)
		DataHandler = DataHandler.Debug()

		// 注册model
		// 自动同步数据库模型

		// 客户相关
		DataHandler.AutoMigrate(&CustomerGroups{})
		DataHandler.AutoMigrate(&Customers{})

		// 订单相关
		DataHandler.AutoMigrate(&Commands{})
		DataHandler.AutoMigrate(&Invoices{})
		DataHandler.AutoMigrate(&Reviews{})

		// 菜品相关
		DataHandler.AutoMigrate(&Categories{})
		DataHandler.AutoMigrate(&Menus{})

		// 门店相关
		DataHandler.AutoMigrate(&Brands{})
		DataHandler.AutoMigrate(&Finances{})
		DataHandler.AutoMigrate(&Departments{})
		DataHandler.AutoMigrate(&StoreGroup{})
		DataHandler.AutoMigrate(&Incomes{})
		DataHandler.AutoMigrate(&Stores{})
		DataHandler.AutoMigrate(&MenuVariants{})

		// 用户管理(用于管理当前系统的用户权限)
		//// 商户号申请
		//DataHandler.AutoMigrate(&MerchantApply{})
		//DataHandler.AutoMigrate(&AdminAccount{})
		//// 增加会员帐户信息
		//DataHandler.AutoMigrate(&MemberInfo{})
		//DataHandler.AutoMigrate(&Products{})
		//
		//// 店铺信息
		//DataHandler.AutoMigrate(&StoreGroupInfo{})
		//
		//DataHandler.AutoMigrate(&User{})
		//// 品牌管理 (超级管理员权限)
		////DataHandler.AutoMigrate(&Brand{})
		//// 门店管理 (超级管理员权限)
		//DataHandler.AutoMigrate(&StoreModel{})
		//// 部门管理 (超级管理员权限)
		//DataHandler.AutoMigrate(&Department{})
		//// 菜品分类
		//DataHandler.AutoMigrate(&Category{})
		//// 规格
		//DataHandler.AutoMigrate(&Unit{})
		//// 菜品库
		//DataHandler.AutoMigrate(&Dishes{})
		//// 菜品分组
		//DataHandler.AutoMigrate(&MenuGroup{})
		//// 菜品做法
		//DataHandler.AutoMigrate(&FormulaGroup{})
		//// 账号
		//DataHandler.AutoMigrate(&AccountInfo{})
		//// 账号
		//// DataHandler.AutoMigrate(&Cart{})
		//// 账号
		//DataHandler.AutoMigrate(&PaymentFlow{})
		//// 账号
		//DataHandler.AutoMigrate(&AddressModel{})
		//
		//DataHandler.AutoMigrate(&Finance{})
		// 账号
		// DataHandler.AutoMigrate(&Order{})

	}
	return nil
}
