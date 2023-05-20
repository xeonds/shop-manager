package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//go:embed frontend/dist/*
var f embed.FS
var r *gin.Engine
var db *gorm.DB
var config *Config

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Admin struct {
		Username string
		Password string
	}
}

// router handlers
func InitRouter() {
	r = gin.Default()
	apiRouter := r.Group("/api")

	apiRouter.GET("/customers", GetCustomers)
	apiRouter.GET("/customers/:id", GetCustomer)
	apiRouter.POST("/customers", CreateCustomer)
	apiRouter.PUT("/customers/:id", UpdateCustomer)
	apiRouter.DELETE("/customers/:id", DeleteCustomer)
	subFS, err := fs.Sub(f, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.FS(subFS))))
}

// database actions
func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败：%v", err)
	}
	db.AutoMigrate(&Customer{}, &Product{}, &Order{})
	return nil
}
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
	if err := db.First(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "记录不存在"})
		return
	}
	c.JSON(http.StatusOK, customers)
}
func GetCustomer(c *gin.Context) {
	var customer Customer
	id := c.Param("id")
	if err := db.First(&customer, id).Error; err != nil {
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
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
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
func CreateOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&order)
	c.JSON(http.StatusOK, order)
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
func UpdateOrder(c *gin.Context) {
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
	db.Save(&order)
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

// utilities
func InitConfig() {
	conf, err := LoadConfig("config")
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = CreateDefaultConfig("config.yaml")
			if err != nil {
				log.Fatalln(err)
			}
			panic("请修改配置文件后重启程序")
		} else {
			panic(err)
		}
	}
	config = conf
}
func LoadConfig(filename string) (*Config, error) {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
func CreateDefaultConfig(filename string) error {
	viper.SetDefault("server.host", "127.0.0.1")
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("admin.username", "admin")
	viper.SetDefault("admin.password", "admin")
	if err := viper.WriteConfigAs(filename); err != nil {
		return err
	}
	return nil
}
func ReplaceInvalidChars(name string) string {
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	return re.ReplaceAllString(name, "_")
}

// main
func main() {
	InitConfig()
	InitRouter()
	InitDB()

	panic(r.Run(config.Server.Host + ":" + config.Server.Port)) // run the server on the specified host and port
}
