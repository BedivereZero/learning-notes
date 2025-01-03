package main

import (
	"encoding/json"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn string = "root:rockman_mysql@/dev?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("create database session fail: %v", err)
	}

	// if err := db.AutoMigrate(&Example{}, &Scene{}, &SceneResourceBinding{}); err != nil {
	// 	log.Fatalf("auto migrate fail: %v", err)
	// }

	db = db.Debug()

	// log.Print("create example")
	// if err := create(db.Debug()); err != nil {
	// 	log.Fatalf("create example fail: %v", err)
	// }

	log.Print("list examples")
	if err := list(db); err != nil {
		log.Fatalf("list examples fail: %v", err)
	}

	// log.Print("get example")
	// if err := get(db); err != nil {
	// 	log.Fatalf("get example fail: %v", err)
	// }

	// log.Print("create example")
	// if err := createExample(db); err != nil {
	// 	log.Fatalf("create example fail: %v", err)
	// }
}

func list(tx *gorm.DB) error {
	var users []User
	if err := tx.Where("name like ?", "%Scarlet%").Find(&users).Error; err != nil {
		return err
	}
	log.Printf("found %d users", len(users))
	for i, u := range users {
		log.Printf("users[%d]: %s", i, u.Name)
	}
	return nil
}

// func get(tx *gorm.DB) error {
// 	var example Example

// 	example.ID = 6

// 	tx = tx.Take(&example)
// 	log.Printf("example: %s", encodeAsJSON(&example))

// 	if err := tx.Model(&example).Association("Scenes").Find(&example.Scenes); err != nil {
// 		return fmt.Errorf("find scenes fail: %w", err)
// 	}
// 	log.Printf("example: %s", encodeAsJSON(&example))

// 	// tx = tx.Preload("Scenes")
// 	// tx = tx.Take(&example, 3)
// 	// log.Printf("example: %s", encodeAsJSON(&example))

// 	return tx.Error
// }

// func createExample(tx *gorm.DB) error {
// 	example := &Example{
// 		Name: "example-0001",
// 		Scenes: []Scene{
// 			{Name: "scene-0010"},
// 			{Name: "scene-0011"},
// 			{Name: "scene-0012"},
// 		},
// 	}

// 	tx = tx.Create(example)
// 	return tx.Error
// }

func encodeAsJSON(v any) []byte {
	b, _ := json.MarshalIndent(v, "", "  ")
	return b
}
