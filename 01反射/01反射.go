package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Life int
}

type People struct {
	Animal
	Name string "姓名"
	Age  int    "年龄"
}

func (p People) GetName() string {
	fmt.Println("GetName called", p.Name)
	return p.Name
}

func (p *People) SetName(name string) {
	fmt.Println("SetName called:", name)
	p.Name = name
}

//func(string,int)int
func (p *People) Eat(food string, grams int) (power int) {
	fmt.Println(p.Name,"Eat called:", food, grams)
	return 5 * grams
}

func main() {
	p := People{Name: "张三", Age: 20, Animal: Animal{100}}
	//typeAPI(p)
	valueAPI(p)
}

/*
访问和修改对象的任意属性的值
访问对象的任意方法
*/
func valueAPI(o People) {

	oValue := reflect.ValueOf(o)
	fmt.Println(oValue)//{{100} 张三 20}

	fmt.Println("查看oValue的原始类型")
	fmt.Println(oValue.Kind())//struct

	fmt.Println("查看p的所有属性的值")
	for i := 0; i < oValue.NumField(); i++ {
		fValue := oValue.Field(i)
		fmt.Println(fValue.Interface())
	}

	fmt.Println("获得父类的属性值")
	fatherFieldValue := oValue.FieldByIndex([]int{0, 0})
	fmt.Println(fatherFieldValue.Interface())

	fmt.Println("获得指针value的内容进而获得成员的值")
	opValue := reflect.ValueOf(&o)
	//opValue是一个People指针的Value，oValue是一个People值的Value
	oValue = opValue.Elem()
	fmt.Println(oValue.Field(1).Interface())

	fmt.Println("修改对象的属性")
	oValue.FieldByName("Age").SetInt(60)
	fmt.Println(oValue)

	fmt.Println("调用对象的方法")
	mValue := opValue.MethodByName("SetName")
	mValue.Call([]reflect.Value{reflect.ValueOf("张三疯")})
	fmt.Println(oValue)

	mValue = opValue.MethodByName("Eat")
	retValues := mValue.Call([]reflect.Value{reflect.ValueOf("猫屎咖啡"), reflect.ValueOf(100)})
	fmt.Println("吃完的热量是", retValues[0].Interface())
}

/*对任意对象的细致入微的类型检测*/
func typeAPI(obj interface{}) {
	oType := reflect.TypeOf(obj)
	fmt.Println(oType) //main.People

	//原始类型
	oKind := oType.Kind()
	fmt.Println(oKind) //struct

	//类型名称
	fmt.Println(oType.Name()) //People

	//属性和方法的个数
	fmt.Println(oType.NumField())  //3
	fmt.Println(oType.NumMethod()) //2

	fmt.Println("全部属性：")
	for i := 0; i < oType.NumField(); i++ {
		structField := oType.Field(i)
		fmt.Println(structField.Name, structField.Type)
	}

	fmt.Println("全部方法：")
	for i := 0; i < oType.NumMethod(); i++ {
		method := oType.Method(i)
		fmt.Println(method.Name, method.Type)
	}

	fmt.Println("获得父类属性")
	fmt.Println(oType.FieldByIndex([]int{0, 0}).Name)
}
