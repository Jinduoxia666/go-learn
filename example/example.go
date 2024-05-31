package example

import "fmt"

const name string = "金多虾"

func SayJinduoxia() {
	student := Student{
		Name: name,
		Age:  18,
		Sex:  true,
	}
	fmt.Println(student)
}

type Student struct {
	Name string
	Age  int
	Sex  bool
}

func PlusAge(stu *Student) {
	stu.Age++
}
