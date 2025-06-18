package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Record defines the data model
type Record struct {
	gorm.Model
	Name   string
	Value1 string
	Value2 string
	Value3 string
}

var db *gorm.DB

func initDB() {
	dsn := "host=192.168.1.182 user=user_jnNjMC password=password_NtxCWa dbname=gin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Record{})
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

func submitHandler(c *gin.Context) {
	name := c.PostForm("name")
	value1 := c.PostForm("value1")
	value2 := c.PostForm("value2")
	value3 := c.PostForm("value3")

	db.Create(&Record{
		Name:   name,
		Value1: value1,
		Value2: value2,
		Value3: value3,
	})

	c.Redirect(http.StatusMovedPermanently, "/")
}

func viewHandler(c *gin.Context) {
	var records []Record
	db.Find(&records)

	c.HTML(http.StatusOK, "records.html", gin.H{
		"records": records,
	})
}

func main() {
	initDB()

	r := gin.Default()

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	r.GET("/", indexHandler)
	r.POST("/submit", submitHandler)
	r.GET("/view", viewHandler)

	r.Run("127.0.0.1:8080")
}
