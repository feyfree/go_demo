package go20230211

import (
	"fmt"
)

// 简单说来，
//类型struct{T}和*struct{T}均将获取类型T的所有方法。
//类型*struct{T}、struct{*T}和*struct{*T}都将获取类型*T的所有方法。

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person // 通过内嵌Person类型来扩展之
	works  []string
}
