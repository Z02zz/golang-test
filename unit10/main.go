package main	
import (
	"fmt"
)

//定义动物结构体
type Animal struct{
	Age int
	Weight float32
}

//绑定喊叫方法
func (an *Animal) Shout(){
	fmt.Println("aaaaaaaaa")
}

func (an *Animal) ShowSelf(){
	fmt.Printf("体重是：%v,年龄是：%v\n",an.Weight,an.Age)
}

//为了体现复用性，体现继承，内嵌一个匿名构体
type Cat struct{
	Animal
}

func (c *Cat) scratch(){
	fmt.Println("naoren")
}

func main(){
	cat := &Cat{}
	cat.Animal.Age = 7
	cat.Animal.Weight = 10.5
	cat.Animal.Shout()
	cat.Animal.ShowSelf()
	cat.scratch()
}