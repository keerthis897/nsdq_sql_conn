// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {

// 	db, err := sql.Open("mysql", "root:keerthi1234@tcp(localhost:3306)/testdb")

// 	if err != nil {
// 		fmt.Println("error validating sql.Open arguments")
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		fmt.Println("error verifying connection with db.Ping")
// 		panic(err.Error())
// 	}

// 	insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('3', 'keerthi', 'shekar');")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer insert.Close()
// 	fmt.Println("Successful Connection to Database!")
// }
