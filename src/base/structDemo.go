package base

import "fmt"

// StructTest 结构体测试入口
func StructTest() {
	fmt.Println("进入structTest 方法。。。。")

	user := User{"老王", 12}

	p := &user

	username := user.getName()
	//下面两种写法都可以
	fmt.Printf("用户名：%s,年龄：%d \n", username, user.age)

	fmt.Printf("用户名：%s,年龄：%d \n", p.getName(), p.age)

	n := p.getName

	fmt.Printf("用户名：%s\n", n())

	//接口测试
	var peo People

	peo = &Student{}

	think := "fack"

	fmt.Printf("说话的值:%s", peo.Speaking(think))

}

//User 定义用户结构体
type User struct {
	name string
	age  int
}

func (u *User) getName() string {
	return u.name
}

//People 接口
type People interface {
	Speaking(string) string
}

//Student 学生结构体
type Student struct{}

//Speaking 方法
func (s *Student) Speaking(think string) (talk string) {
	if think == "你好" {
		talk = "你好  老师"
	} else {
		talk = "你为什么不说你好"
	}
	return
}
