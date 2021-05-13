package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: 100 * time.Millisecond,
				LogLevel:      logger.Info,
				Colorful:      true,
			}),
		},
	)
	if err != nil {
		panic(err)
	}
	content := content{
		Name: "Quang23",
		Age:  24,
	}

	contentJson, _ := json.Marshal(content)

	data := Data{
		ID:      8,
		Content: string(contentJson),
		Value:   new(int),
	}
	err = db.Save(&data).Error
	if err != nil {
		print("err")
		panic(err)
	}
	print("ok")
}

type Data struct {
	ID        int `gorm:"primaryKey"`
	Content   string
	Value     *int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type content struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *Data) TableName() string {
	return "test"
}
