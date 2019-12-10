package main

import (
	//并不需要使用其API,只需要执行其包的init方法(事实上加载MySQL的驱动程序)
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
)

/*该结构体对应着数据库里的person表*/
type Person struct {
	//对应name表字段
	Name string `db:"name"`
	//对应age表字段
	Age int `db:"age"`
	//对应rmb表字段
	Money float64 `db:"rmb"`
}

/*执行增删改*/
func main011() {
	//链接数据库
	db, _ := sqlx.Open("mysql", "root:123456@tcp(localhost:3306)/mydb")
	defer db.Close()

	//执行增删改
	//result, e := db.Exec("insert into person(name,age,rmb,gender,birthday) values(?,?,?,?,?);","张巨蛋",50,123,false,19680101)
	//result, e := db.Exec("delete from person where name not like ?;","%蛋")
	result, e := db.Exec("update person set name = ? where id = ?","张一蛋",1)
	if e!=nil{
		fmt.Println("err=",e)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	lastInsertId, _ := result.LastInsertId()
	fmt.Println("受影响的行数=",rowsAffected)
	fmt.Println("最后一行的ID=",lastInsertId)
}

/*执行分查询*/
func main() {
	database, _ := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb")
	defer database.Close()

	//预定义Person切片用于接收查询结果
	var ps []Person

	//执行查询,得到Person对象的集合,丢入预定义的ps地址
	e := database.Select(&ps, "select name,age,rmb from person where name like ?", "%蛋")
	if e!=nil{
		fmt.Println("err=",e)
		return
	}

	fmt.Println("查询成功",ps)

}
