package main

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Id      int     `json:"id"`
	Name    string  `json:"name"`                                  //姓名
	Gender  string  `json:"gender"`                                //性别
	Phone   string  `json:"phone"`                                 //电话
	Order   []Order `json:"order" gorm:"foreignKey:CustomerId"`    //订单
	Wallet  float64 `json:"wallet"`                                //钱包
	IsVIP   bool    `json:"is_vip"`                                //是否是VIP
	VIPCard VIPCard `json:"vip_card" gorm:"foreignKey:CustomerId"` //会员信息
	Note    string  `json:"note"`                                  //备注
}

type Product struct {
	gorm.Model
	Id          int     `json:"id"`
	Name        string  `json:"name"`        //产品名称
	Brand       string  `json:"brand"`       //品牌
	Location    string  `json:"location"`    //产地
	Description string  `json:"description"` //描述
	Price       float64 `json:"price"`       //价格
	Purchase    float64 `json:"purchase"`    //进货量
	Sale        float64 `json:"sale"`        //销售量
	Quantity    float64 `json:"quantity"`    //库存
	Note        string  `json:"note"`        //备注
}

type Purchase struct {
	gorm.Model
	Id         int     `json:"id"`
	Product    Product `json:"product" gorm:"embedded;embeddedPrefix:product_"` // 产品信息
	Supplier   string  `json:"supplier"`                                        //供应商
	Quantity   float64 `json:"quantity"`                                        //数量
	UnitPrice  float64 `json:"unit_price"`                                      //单价
	TotalPrice float64 `json:"total_price"`                                     //总价
	Paid       float64 `json:"paid"`                                            //已付
	Debt       float64 `json:"debt"`                                            //欠款
	Note       string  `json:"note"`                                            //备注
}

type Order struct {
	gorm.Model
	Id         int     `json:"id"`
	CustomerId int     `json:"customer_id"`                                     //顾客id
	Product    Product `json:"product" gorm:"embedded;embeddedPrefix:product_"` // 产品信息
	Quantity   float64 `json:"quantity"`                                        //数量
	UnitPrice  float64 `json:"unit_price"`                                      //单价
	TotalPrice float64 `json:"total_price"`                                     //总价
	Discount   float64 `json:"discount"`                                        //折扣
	Paid       float64 `json:"paid"`                                            //已付
	Payment    string  `json:"payment"`                                         //付款方式
	Note       string  `json:"note"`                                            //备注
}

type VIPCard struct {
	gorm.Model
	Id         int       `json:"id"`
	CustomerId int       `json:"customer_id"` //顾客id
	Date       time.Time `json:"date"`        //办理日期
	Balance    float64   `json:"balance"`     //余额
	Discount   float64   `json:"discount"`    //折扣
	Note       string    `json:"note"`        //备注
}
