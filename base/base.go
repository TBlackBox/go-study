package base

import (
	"fmt"
	"math"
)

//单个定义常量
const pi = 3.1415

//批量定义常量 如果省略值则和上面值相同  即 v1 = 1 v2 = 1
const (
	v = 1
	v1
	v2
)

//iota 可以理解为常量计数器
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

// //使用_跳过某些值
// const (
// 		n1 = iota //0
// 		n2        //1
// 		_
// 		n4        //3
// 	)
// //iota声明中间插队
// const (
// 		n1 = iota //0
// 		n2 = 100  //100
// 		n3 = iota //2
// 		n4        //3
// 	)
// const n5 = iota //0

// 内存大小单位
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func init() {
	fmt.Printf("基础函数初始化。。。\n")
}



//Test 基础测试方法
func Test() {

	fmt.Printf("开始基本方法测试\n")

	//StructTest()

	sliceDemo()

	// arrayDemo()

	// sqrtDemo()

	// changeString()

	// constTest()

	// base()
}

//通过make 创建切片
var slice7 []int = make([]int, 10)
var slice5 = make([]int, 10)
var slice6 = make([]int, 10, 10)

//切片的初始化
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
//切片
func sliceDemo() {

	//创建的方式

	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr1 := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr1[1:4]
	fmt.Println(s6)

	//切片的初始化
	fmt.Printf("全局变量：arr %v\n", arr)
	fmt.Printf("全局变量：slice0 %v\n", slice0)
	fmt.Printf("全局变量：slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr[2:8]
	slice6 := arr[0:6]         //可以简写为 slice := arr[:end]
	slice7 := arr[5:10]        //可以简写为 slice := arr[start:]
	slice8 := arr[0:len(arr)]  //slice := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素
	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)

	fmt.Printf("make全局slice0 ：%v\n", slice7)
	fmt.Printf("make全局slice1 ：%v\n", slice5)
	fmt.Printf("make全局slice2 ：%v\n", slice6)
	fmt.Println("--------------------------------------")
	slice12 := make([]int, 10)
	slice13 := make([]int, 10)
	slice14 := make([]int, 10, 10)
	fmt.Printf("make局部slice3 ：%v\n", slice12)
	fmt.Printf("make局部slice4 ：%v\n", slice13)
	fmt.Printf("make局部slice5 ：%v\n", slice14)

	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	b := a[1:2]
	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))

}

//一维数组
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

//多维数组
var arr10 [5][3]int
var arr11 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func arrayDemo() {
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用引号初始化元素。
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)

	//多维数组
	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	f := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。
	fmt.Println(arr10, arr11)
	fmt.Println(e, f)
}

//强制类型转换  go没有隐式类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//改变字符串
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

func constTest() {

	fmt.Printf("MB:%d,%b", MB, MB)
	fmt.Printf("KB:%d,%b\n", KB, KB)

	fmt.Printf("int64,最小值：%d,最大值：%d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("int8,最小值：%d,最大值：%d\n", math.MinInt8, math.MaxInt8)
}

func base() {

	//单个变量声明
	var name string = "老王"
	var age int
	var year int8
	var isSuccess bool

	//批量变量声明 字符串默认值值""
	var (
		sex  string
		code int
	)

	fmt.Printf("定义的基本变量默认值：%s ,%d ,%d,%t,%s,%d\n", name, age, year, isSuccess, sex, code)

	color := "red"

	fmt.Printf("短变量:%s\n", color)

	//最后一个为匿名变量
	var height, width, _ = 12, 13, "mmd"

	fmt.Printf("高：%d,宽：%d\n", height, width)
}
