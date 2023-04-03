package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"k8s.io/klog/v2/klogr"
)

const dsn = "root:rockman@mysql@tcp(127.0.0.1:3306)/test?parseTime=true"

type Person struct {
	gorm.Model

	Name string
	Age  int
}

func main() {
	var logger = klogr.NewWithOptions(klogr.WithFormat(klogr.FormatKlog))
	var p Person

	logger.Info("open database")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err, "open connection fail")
		return
	}

	// logger.Info("auto migrate")
	// if err := db.AutoMigrate(&Person{}); err != nil {
	// 	logger.Error(err, "auto migrate fail")
	// 	return
	// }

	// logger.Info("create person")
	// if err := db.Create(&Person{Name: "hello", Age: 10}).Error; err != nil {
	// 	logger.Error(err, "create person fail")
	// 	return
	// }

	logger.Info("get first")
	if err := db.First(&p).Error; err != nil {
		logger.Error(err, "get first fail")
	}
	logger.Info("get first", "name", p.Name, "age", p.Age)

	var d = time.Minute
	logger.Info("sleep", "duration", d)
	time.Sleep(d)

	// logger.Info("create person")
	// if err := db.Create(&Person{Name: "hello", Age: 10}).Error; err != nil {
	// 	logger.Error(err, "create person fail")
	// 	return
	// }

	logger.Info("get first")
	if err := db.First(&p).Error; err != nil {
		logger.Error(err, "get first fail")
	}
	logger.Info("get first", "name", p.Name, "age", p.Age)
}
