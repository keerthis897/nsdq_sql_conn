// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/sqlserver"
// 	"gorm.io/gorm"
// )

// type StudentNew1 struct {
// 	ID   int    `gorm:"primaryKey"`
// 	Name string `gorm:"size:50"`
// }

// func main() {
// 	dsn := "sqlserver://test5:1111@localhost:1433?database=test5"

// 	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatalf("Error getting database connection: %v", err)
// 	}
// 	defer sqlDB.Close()

// 	ctx := context.Background()

// 	err = sqlDB.PingContext(ctx)
// 	if err != nil {
// 		log.Fatalf("Error pinging database: %v", err)
// 	}

// 	fmt.Println("Connected to SQL Server!")
// }

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type StudentNew1 struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:50"`
}

func main() {
	dsn := "sqlserver://test5:1111@localhost:1433?database=test5"

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to SQL Server!")

	if !db.Migrator().HasTable(&StudentNew1{}) {
		if err := db.AutoMigrate(&StudentNew1{}); err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}

	newStudent := StudentNew1{Name: "John9"}
	result := db.Create(&newStudent)
	if result.Error != nil {
		log.Fatalf("Error inserting data: %v", result.Error)
	}
	fmt.Printf("%d row(s) inserted.\n", result.RowsAffected)

	deleteResult := db.Where("Name = ?", "Johna").Delete(&StudentNew1{})
	if deleteResult.Error != nil {
		log.Fatalf("Error deleting data: %v", deleteResult.Error)
	}
	fmt.Printf("%d row(s) deleted.\n", deleteResult.RowsAffected)

	updateResult := db.Model(&StudentNew1{}).Where("ID = ?", "2").Update("Name", "John8")
	if updateResult.Error != nil {
		log.Fatalf("Error updating data: %v", updateResult.Error)
	}
	fmt.Printf("%d row(s) updated.\n", updateResult.RowsAffected)
}
