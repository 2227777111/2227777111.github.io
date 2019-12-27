package main

import "fmt"

/*
类型断言(type assertion)
判断一个接口实例的具体类型

类型断言的两种方式：
方式一，判断接口实例的具体类型：
switch xxx.(type){
case *Coder:
//...
case *Boss:
//...
}

方式二：判断接口实例是不是程序猿
if coder,ok:=xxx.(*Coder);ok{
	//确实是程序猿
	//此时的coder是程序猿指针类型
}else{
	//xxx压根不是程序猿
}
*/

/*
劳动者父类接口
内含两个抽象方法：工作、休息
*/
type Worker interface {
	//每天工作多少小时，产出何种产品
	Work(hour int) (product string)
	//休息
	Rest()
}

/*定义码农结构体*/
type Coder struct {
	skill string
}
/*码农指针实现Worker接口*/
func (c *Coder)Work(hour int) (product string){
	fmt.Printf("码农一天工作%d小时\n",hour)
	fmt.Printf("码农正在%s\n",c.skill)
	return "BUG"
}
func (c *Coder)Rest(){
	fmt.Println("休息是什么？？")
}
/*码农特有的方法*/
func (c *Coder)WorkHome()  {
	fmt.Println("程序员在家工作")
}



/*定义产品经理结构体*/
type ProductManager struct {
	skill string
}
/*ProductManager指针实现Worker接口*/
func (pm *ProductManager)Work(hour int) (product string){
	fmt.Printf("产品一天工作%d小时\n",hour)
	fmt.Printf("产品正在%s\n",pm.skill)
	return "无逻辑的需求"
}
func (pm *ProductManager)Rest(){
	fmt.Println("看程序员撸代码")
}



/*定义老板结构体*/
type Boss struct {
	skill string
}
/*Boss指针实现Worker接口*/
func (b *Boss)Work(hour int) (product string){
	fmt.Printf("老板一天工作%d小时\n",hour)
	fmt.Printf("产品正在%s\n",b.skill)
	return "梦想"
}
func (c *Boss)Rest(){
	fmt.Println("一天到晚都在休息")
}

func main041() {
	//创建一支劳动者大军
	workers := make([]Worker, 0)

	//添加不同类型的劳动者
	workers = append(workers, &Coder{"撸代码"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Boss{"吹牛逼"})

	for _,worker := range workers{
		switch worker.(type) {
		case *Coder:
			fmt.Println("伦家是个撸代码滴")
		case *ProductManager:
			fmt.Println("伦家是搞创意的，不要跟我聊逻辑")
		case *Boss:
			fmt.Println("我是给员工打工的")
		default:
			fmt.Println("不知道什么鸟")
		}
	}
}

func main() {
	workers := make([]Worker, 0)
	workers = append(workers, &Boss{"吹牛逼"})
	workers = append(workers, &ProductManager{"拍脑门"})
	workers = append(workers, &Coder{"撸代码"})

	for _,worker := range workers{
		if coderPtr,ok := worker.(*Coder);ok{
			fmt.Println("发现一只程序猿在", coderPtr.skill)
			coderPtr.WorkHome()
		}else {
			fmt.Println(worker,"不是程序猿")
		}
	}

}
