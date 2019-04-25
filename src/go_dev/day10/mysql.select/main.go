package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var DB *sqlx.DB

func init()  {
	database,err = DB.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open sql failed,err:",err)
		return
	}
	DB = database
}

func main()  {
	var person = []Person
	r,err := DB.Select(&person,"select user_id,username,sex,email from person where user_id=?",1)
	if err != nil {
		fmt.Println("exec failed ,er:",err)
		return
	}
	fmt.Println("select succ:",person)
}