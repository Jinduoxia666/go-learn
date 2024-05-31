package example_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/Jinduoxia666/go-learn/example"
)

func Test1(t *testing.T) {
	os.Stdout.WriteString("hello 世界!")
	fmt.Println("hello 2")
}

func Test2(t *testing.T) {
	fmt.Printf("%s\n", "hello world")
}

func Test3(t *testing.T) {
	// 读
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

func Test4(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("执行到第%d个数\n", (i + 1))
	}
}

func Test5(t *testing.T) {
	var count int = 1
	for count < 11 {
		t.Logf("执行到第%d个数\n", count)
		count++
	}
}

func Test6(t *testing.T) {
	seq := "jdiasjdioas"
	for index, value := range seq {
		if index%2 != 0 {
			continue
		}
		fmt.Printf("index:%d, value:%c\n", index, value)
	}
}

/*
数组测试
@author: 金多虾
@since: 2024-05-30
*/
func Test7(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	for _, v := range arr[1:3] {
		fmt.Printf("v:%d\n", v)
	}
}

/*
数组切片测试
*/
func Test8(t *testing.T) {
	arr := []int{1, 2}
	for _, v := range arr {
		fmt.Printf("v:%d\n", v)
	}
}

func Test9(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // cap = 9
	s2 := s1[1:4:5]                        // cap = 9 - 3 = 6
	fmt.Println(s1, s2)
}

func Test10(t *testing.T) {
	p1 := new([]int)
	fmt.Println(*p1)
	pointer := make([]int, 5, 100) // 长度为10，容量100的整型切片
	fmt.Println(pointer)
}

func Test12(t *testing.T) {
	student := example.Student{
		Name: "金多虾",
		Age:  18,
		Sex:  true,
	}
	fmt.Println(student)
	example.PlusAge(&student)
	fmt.Println(student)
}
