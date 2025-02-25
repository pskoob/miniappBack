// // models.go
package models

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID       uint   `json:"id" gorm:"primaryKey"`
// 	Username string `json:"username" gorm:"unique"`
// 	Password string `json:"password"`
// }

// var db *gorm.DB

// func InitDB() {
// 	var err error
// 	dsn := "host=localhost user=pap password=123 dbname=minik port=5432 sslmode=disable"
// 	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to the database")
// 	}
// 	db.AutoMigrate(&User{})
// }
