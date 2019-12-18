package main

import (
	"os"
	"encoding/json"
	"fmt"
	"reflect"
)

/*
所有商品有一些共性：品名、价格
个性无千无万，自行封装出三种商品（以模拟30万种商品）
随意给出一个商品的切片，将每件商品的所有属性信息输出到JSON文件(品名.json)
用反射做
*/
type Computer struct {
	Name string
	Price float64
	Cpu string
	Memory int
	Disk int
}

type TShirt struct {
	Name string
	Price float64
	Color string
	Size int
	Sex bool
}

type Car struct {
	Name string
	Price float64
	Cap int
	Power string
}

func main() {
	products := make([]interface{}, 0)
	products = append(products, Computer{"未来人类",10000,"英特尔i7",16,1024})
	products = append(products, TShirt{"爆款T恤",10000,"红色",40,false})
	products = append(products, Car{"兰博比基尼",10000,5,"油电混合"})

	//EncodeObj2JsonFile(products,"D:/BJBlockChain1803/demos/W2/day4/file/商品总表.json")
	for _,p := range products{

		name := reflect.ValueOf(p).FieldByName("Name").Interface().(string)
		fmt.Println(name)

/*		//先拿到p的值（不在乎类型）
		pValue := reflect.ValueOf(p)

		//再通过p值，拿到p下面的名为Name的属性值
		nameValue := pValue.FieldByName("Name")

		//Name属性的值，反射形式化为正射形式(interface{}类型)
		nameInterface := nameValue.Interface()

		//将空接口断言为string
		name := nameInterface.(string)
		fmt.Println(name)*/

		EncodeObj2JsonFile(p,"D:/BJBlockChain1803/demos/W2/day4/file/"+name+".json")
	}

}

func EncodeObj2JsonFile(obj interface{},filename string) bool {
	dstFile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer dstFile.Close()

	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(obj)
	if err != nil{
		fmt.Println("编码失败，err=",err)
		return false
	}else{
		fmt.Println("编码成功!")
		return true
	}
}
