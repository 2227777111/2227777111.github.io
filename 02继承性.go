package main

import "fmt"

type ghost struct {
	magic string
}

func (g *ghost)doMagic()  {
	fmt.Printf("一只鬼正在%s\n",g.magic)
}

type person struct {
	//封装结构体的属性
	name string
	age int
	sex bool
	hobby []string
}
func (p *person)Eat()  {
	fmt.Printf("%s爱饕餮\n",p.name)
}
func (p *person)Drink()  {
	fmt.Printf("%s爱喝酒\n",p.name)
}
func (p *person)Love()  {
	fmt.Printf("%s是有感情的\n",p.name)
	p.age -= 1
}

/*死堆代码的*/
type coder struct {
	//持有一个父类声明——继承了person
	person

	//会的语言
	langs []string
}

func (c *coder)Code()  {
	fmt.Printf("%s会%v，它正在堆代码",c.name,c.langs)
}

type driver struct {
	person
	ghost

	/*司机的特有属性*/
	jiazhaoID string
	//是不是在开车
	isDriving bool
}

/*司机的特有方法*/
func (d *driver)drive()  {
	fmt.Printf("%s一言不合就发车",d.name)
}

/*覆写父类方法*/
func (d *driver)Drink()  {
	if !d.isDriving{
		fmt.Printf("%s爱喝酒\n",d.name)
	}else {
		fmt.Println("fuckoff，司机一滴酒亲人两行泪")
	}
}


//子类默认拥有父类的全部属性和方法
func main021() {
	//c := coder{person{"西门阿明", 50, true, []string{"腼腆地笑", "猥琐地笑", "尴尬而不失礼貌的微笑"}}}
	c := new(coder)
	c.name = "西门阿明"
	c.langs = []string{"Go","汉语","安徽话"}
	c.Drink()
	c.Code()
}

//访问子类扩展的属性和方法
func main022() {
	dPtr := new(driver)
	//访问父类的属性和方法
	dPtr.name = "拍黄片的朴哥"
	dPtr.Eat()

	//访问特有的属性和方法
	dPtr.jiazhaoID = "PHP许可证京A1024"
	dPtr.drive()
	dPtr.isDriving = true
	dPtr.Drink()
}

//访问子类覆写的父类方法
func main023() {
	p := person{name: "普通人"}
	p.Drink()

	d := new(driver)
	d.name = "老司机"
	d.isDriving = true
	d.Drink()
}

func main() {
	d:= driver{person{name: "女司机"}, ghost{"马路杀"}, "007", true}
	d.Love()
	d.doMagic()
}
