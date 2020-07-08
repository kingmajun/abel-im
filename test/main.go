package test

import "fmt"

//匿名函数
func user(name string) func(int2 int) string {
	return func(age int) string {
		return fmt.Sprint(name, age, "岁了")
	}
}

//闭包=函数+外层变量的引用
func main() {
	ageMethod := user("马俊")
	info := ageMethod(18)
	fmt.Println(info)
}
