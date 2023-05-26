package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	Admin  Admin
	Secret string
}
type Admin struct {
	Username string
	Password string
}

// router handlers
func InitRouter() {
	r = gin.Default()

	apiRouter := r.Group("/api")

	customerRouter := apiRouter.Group("/customers")
	customerRouter.GET("", GetCustomers)
	customerRouter.GET("/:id", GetCustomer)
	customerRouter.POST("", CreateCustomer)
	customerRouter.PUT("/:id", UpdateCustomer)
	customerRouter.DELETE("/:id", DeleteCustomer)

	productRouter := apiRouter.Group("/products")
	productRouter.GET("", GetProducts)
	productRouter.GET("/:id", GetProduct)
	productRouter.POST("", CreateProduct)
	productRouter.PUT("/:id", UpdateProduct)
	productRouter.DELETE("/:id", DeleteProduct)

	orderRouter := apiRouter.Group("/orders")
	orderRouter.GET("", GetOrders)
	orderRouter.GET("/:id", GetOrder)
	orderRouter.POST("", CreateOrder)
	orderRouter.PUT("/:id", UpdateOrder)
	orderRouter.DELETE("/:id", DeleteOrder)

	purchaseRouter := apiRouter.Group("/purchases")
	purchaseRouter.GET("", GetPurchases)
	purchaseRouter.GET("/:id", GetPurchase)
	purchaseRouter.POST("", CreatePurchase)
	purchaseRouter.PUT("/:id", UpdatePurchase)
	purchaseRouter.DELETE("/:id", DeletePurchase)

	vipCardRouter := apiRouter.Group("/vip-cards")
	vipCardRouter.GET("", GetVIPCards)
	vipCardRouter.GET("/:id", GetVIPCard)
	vipCardRouter.POST("", CreateVIPCard)
	vipCardRouter.PUT("/:id", UpdateVIPCard)
	vipCardRouter.DELETE("/:id", DeleteVIPCard)

	apiRouter.Group("/auth").POST("/login", Login)
	apiRouter.Group("/database").POST("/backup", BackupDB)
	apiRouter.Group("/database").POST("/restore", RestoreDB)

	subFS, err := fs.Sub(f, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}
	r.NoRoute(gin.WrapH(http.FileServer(http.FS(subFS))))
}

// database actions
func InitDB() error {
	var err error
	db, err = gorm.Open(sqlite.Open("shop-manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败：%v", err)
	}
	db.AutoMigrate(&Customer{}, &Product{}, &Order{}, &Purchase{}, &VIPCard{})
	return nil
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
	viper.SetDefault("secret", "super-secret-key")
	if err := viper.WriteConfigAs(filename); err != nil {
		return err
	}
	return nil
}
func ReplaceInvalidChars(name string) string {
	re := regexp.MustCompile(`[\\/:*?"<>|]`)
	return re.ReplaceAllString(name, "_")
}
func genToken(userID string, secretKey string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token.Claims = claims

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return ""
	}

	return tokenString
}

// main
func main() {
	InitConfig()
	InitRouter()
	InitDB()

	panic(r.Run(config.Server.Host + ":" + config.Server.Port)) // run the server on the specified host and port
}
