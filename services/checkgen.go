package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/soubhagya")
	if err != nil {
		fmt.Println("error ", err)
	}
	//defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(version)
}
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
func GenerateId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// Generate a snowflake ID.
	id := node.Generate().Int64()
	fmt.Println("Generated id is : ", id)
	// rand.Seed(time.Now().UnixNano())
	// min := 1111000111
	// max := 9999999999
	// fmt.Println(rand.Intn(max-min+1) + min)
	// a := rand.Intn(max-min+1) + min
	return id
}
