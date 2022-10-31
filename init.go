package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DataHandler 数据库
	DataHandler *gorm.DB
	// err 订阅错误
	err error
)

// InitDataBase 初始化数据库
func InitDataBase(dsn string) error {
	var err error
	DataHandler, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 注册model
	// 自动同步数据库模型
	// 用户管理(用于管理当前系统的用户权限)
	// 商户号申请
	DataHandler.AutoMigrate(&MerchantApply{})
	DataHandler.AutoMigrate(&AdminAccount{})
	// 增加会员帐户信息
	DataHandler.AutoMigrate(&MemberInfo{})

	// 店铺信息
	DataHandler.AutoMigrate(&StoreInfo{})

	DataHandler.AutoMigrate(&User{})
	// 品牌管理 (超级管理员权限)
	DataHandler.AutoMigrate(&Brand{})
	// 门店管理 (超级管理员权限)
	DataHandler.AutoMigrate(&StoreModel{})
	// 部门管理 (超级管理员权限)
	DataHandler.AutoMigrate(&Department{})
	// 菜品分类
	DataHandler.AutoMigrate(&Category{})
	// 规格
	DataHandler.AutoMigrate(&Unit{})
	// 菜品库
	DataHandler.AutoMigrate(&Dishes{})
	// 菜品分组
	DataHandler.AutoMigrate(&MenuGroup{})
	// 菜品做法
	DataHandler.AutoMigrate(&FormulaGroup{})
	// 账号
	DataHandler.AutoMigrate(&AccountInfo{})
	// 账号
	// DataHandler.AutoMigrate(&Cart{})
	// 账号
	DataHandler.AutoMigrate(&PaymentFlow{})
	// 账号
	DataHandler.AutoMigrate(&AddressModel{})

	DataHandler.AutoMigrate(&Finance{})
	// 账号
	// DataHandler.AutoMigrate(&Order{})
	return nil
}
