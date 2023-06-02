package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 顾客处理
func CreateCustomer(c *gin.Context) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&customer)
	c.JSON(http.StatusOK, customer)
}
func GetCustomers(c *gin.Context) {
	var customers []Customer
	if err := db.Preload("Order").Preload("VIPCard").Find(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, customers)
}
func GetCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := db.Preload("Order").Preload("VIPCard").First(&customer, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, customer)
}
func UpdateCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(customer.VIPCard)
	db.Save(&customer)
	c.JSON(http.StatusOK, customer)
}
func DeleteCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := db.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	db.Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 产品处理
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}
func GetProducts(c *gin.Context) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, products)
}
func GetProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, product)
}
func UpdateProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&product)
	c.JSON(http.StatusOK, product)
}
func DeleteProduct(c *gin.Context) {
	var product Product
	id := c.Param("id")
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 订单处理
func CreateOrder(c *gin.Context) {
	var order Order
	var product Product
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析失败：" + err.Error()})
		return
	}
	if err := db.First(&product, order.Product.Id).Error; err != nil { // 同时更新产品库存
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if order.Payment == "会员卡" {
		var vipCard VIPCard
		if err := db.First(&vipCard, order.CustomerId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
			return
		}
		if vipCard.Balance < order.Paid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "会员卡余额不足"})
			return
		}
		vipCard.Balance -= order.Paid
		db.Save(&vipCard)
	}
	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "库存不足"})
		return
	}
	product.Quantity -= order.Quantity
	product.Sale += order.Quantity
	db.Save(&product)
	db.Create(&order)
	c.JSON(http.StatusOK, order)
}
func GetOrders(c *gin.Context) {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func GetTodayOrders(c *gin.Context) {
	//获取今日订单
	var orders []Order
	if err := db.Where("created_at > ?", time.Now().Format("2006-01-02")).Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
func GetOrder(c *gin.Context) {
	var order Order
	id := c.Param("id")
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, order)
}
func UpdateOrder(c *gin.Context) { //这个函数不应使用，因为订单不应该修改
	var order Order
	id := c.Param("id")
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if order.Recall {
		c.JSON(http.StatusBadRequest, gin.H{"error": "订单已退货"})
		return
	}
	order.Recall = true
	db.Save(&order)
	//创建一个退货订单，数量为原订单数量的负数，总价为原订单总价的负数
	db.Create(&Order{
		CustomerId: order.CustomerId,
		Product:    order.Product,
		Quantity:   -order.Quantity,
		UnitPrice:  order.UnitPrice,
		TotalPrice: -order.TotalPrice,
		Discount:   order.Discount,
		Paid:       -order.Paid,
		Payment:    order.Payment,
		Note:       "退货订单",
		Recall:     true,
	})
	//回退产品库存
	var product Product
	if err := db.First(&product, order.Product.Id).Error; err != nil { // 同时更新产品库存
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	//如果支付方式是会员卡，退款到会员卡
	if order.Payment == "会员卡" {
		var vipCard VIPCard
		if err := db.First(&vipCard, order.CustomerId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
			return
		}
		vipCard.Balance += order.TotalPrice
		db.Save(&vipCard)
	}
	product.Quantity += order.Quantity
	product.Sale -= order.Quantity
	db.Save(&product)
	c.JSON(http.StatusOK, order)
}
func DeleteOrder(c *gin.Context) {
	var order Order
	id := c.Param("id")
	if err := db.First(&order, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	db.Delete(&order)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 进货处理
func CreatePurchase(c *gin.Context) {
	var purchase Purchase
	var product Product
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.First(&product, purchase.Product.Id).Error; err != nil { // 同时更新产品库存
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	product.Quantity += purchase.Quantity
	product.Purchase += purchase.Quantity
	db.Create(&purchase)
	db.Save(&product)
	c.JSON(http.StatusOK, purchase)
}
func GetPurchases(c *gin.Context) {
	var purchases []Purchase
	if err := db.Find(&purchases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, purchases)
}
func GetPurchase(c *gin.Context) {
	var purchase Purchase
	id := c.Param("id")
	if err := db.First(&purchase, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, purchase)
}
func UpdatePurchase(c *gin.Context) {
	var purchase Purchase
	id := c.Param("id")
	if err := db.First(&purchase, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if purchase.Recall {
		c.JSON(http.StatusBadRequest, gin.H{"error": "订单已退货"})
		return
	}
	purchase.Recall = true
	db.Save(&purchase)
	//新建一个退货订单，数量为原订单数量的负数，总价为原订单总价的负数，清除id
	db.Create(&Purchase{
		Product:    purchase.Product,
		Supplier:   purchase.Supplier,
		Quantity:   -purchase.Quantity,
		UnitPrice:  purchase.UnitPrice,
		TotalPrice: -purchase.TotalPrice,
		Paid:       -purchase.Paid,
		Debt:       -purchase.Debt,
		Note:       "撤单订单",
		Recall:     true,
	})
	//回退产品库存
	var product Product
	if err := db.First(&product, purchase.Product.Id).Error; err != nil { // 同时更新产品库存
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	product.Quantity -= purchase.Quantity
	product.Purchase -= purchase.Quantity
	db.Save(&product)
	c.JSON(http.StatusOK, purchase)
}
func DeletePurchase(c *gin.Context) { //这个函数不应使用，因为进货记录不应该删除
	var purchase Purchase
	id := c.Param("id")
	if err := db.First(&purchase, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	db.Delete(&purchase)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// VIP卡处理
func CreateVIPCard(c *gin.Context) {
	var vipCard VIPCard
	if err := c.ShouldBindJSON(&vipCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&vipCard)
	c.JSON(http.StatusOK, vipCard)
}
func GetVIPCards(c *gin.Context) {
	var vipCards []VIPCard
	if err := db.Find(&vipCards).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, vipCards)
}
func GetVIPCard(c *gin.Context) {
	var vipCard VIPCard
	id := c.Param("id")
	if err := db.First(&vipCard, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, vipCard)
}
func UpdateVIPCard(c *gin.Context) {
	var vipCard VIPCard
	id := c.Param("id")
	if err := db.First(&vipCard, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	if err := c.ShouldBindJSON(&vipCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&vipCard)
	c.JSON(http.StatusOK, vipCard)
}
func DeleteVIPCard(c *gin.Context) {
	var vipCard VIPCard
	id := c.Param("id")
	if err := db.First(&vipCard, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	db.Delete(&vipCard)
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// 系统管理
func Login(c *gin.Context) {
	// get admin info from viper and compare with the request body
	var admin Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if admin.Username != config.Admin.Username || admin.Password != config.Admin.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "token": genToken(admin.Username, "secret")})
}
func BackupDB(c *gin.Context) {
	c.File("./shop-manager.db")
	c.JSON(http.StatusOK, gin.H{"message": "数据库备份成功"})
}
func RestoreDB(c *gin.Context) {
	var file struct {
		File string `form:"file" binding:"required"`
	}
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SaveUploadedFile(c.Request.MultipartForm.File["file"][0], "./shop-manager.db")
	c.JSON(http.StatusOK, gin.H{"message": "数据库恢复成功"})
}
