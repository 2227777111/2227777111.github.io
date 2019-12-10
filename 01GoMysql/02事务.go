package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

func main() {
	db, _ := sqlx.Connect("mysql", "root:123456@tcp(localhost:3306)/mydb")
	defer db.Close()

	//开启事务
	tx, _ := db.Begin()

	//执行系列增删改方法
	//sex列并不存在，性别列叫做gender
	//ret1, e1 := tx.Exec("insert into person(name,age,sex) values(?,?,?)", "咸鸭蛋", 20, true)
	ret1, e1 := tx.Exec("insert into person(name,age,gender) values(?,?,?)", "咸鸭蛋", 20, true)
	ret2, e2 := tx.Exec("delete from person where name not like ?", "%蛋")
	ret3, e3 := tx.Exec("update person set name = ? where name=?", "卤蛋", "双黄蛋")

	//有任何错误都回滚事务，否则提交
	if e1 != nil || e2 != nil || e3 != nil {
		fmt.Println("事务执行失败，e1/e2/e3=", e1, e2, e3)

		//回滚事务
		tx.Rollback()
	} else {

		//提交事务
		tx.Commit()

		ra1, _ := ret1.RowsAffected()
		ra2, _ := ret2.RowsAffected()
		ra3, _ := ret3.RowsAffected()
		fmt.Println("事务执行成功，受影响的行=", ra1+ra2+ra3)
	}
}
