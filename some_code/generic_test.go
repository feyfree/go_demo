package some_code

// 新增的接口
type Animal interface {
	GetName() string
	GetAge() int
}

type Person struct {
	name string
	age  int
}

func (p *Person) GetName() string {
	return p.name
}

func (p Person) GetAge() int {
	return p.age
}

func main() {
	// 定义的接口变量
	var ani Animal

	// person 实现了 Animal 接口，赋值给了 ani 变量
	// 但是，这里编译会通不过，错误如下：
	// Cannot use 'Person{ name: "DaYu", age: int(28), }' (type Person) as the type Animal Type does not implement 'Animal' as the 'GetName' method has a pointer receiver
	//ani = Person{
	//	name: "DaYu",
	//	age:  int(28),
	//}

	ani = &Person{
		name: "DaYu",
		age:  int(28),
	}

	ani.GetName()
	ani.GetAge()
}
