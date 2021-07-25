package base

import (
	"fmt"
	"time"
)

//StartTimer 定时器测试
func StartTimer(){
	//1.timer基本使用 两秒后执行 可用于延时功能
	timer1 := time.NewTimer(2 * time.Second)
	t1 := time.Now()
	fmt.Printf("t1:%v\n", t1)
	t2 := <-timer1.C
	fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	timer2 := time.NewTimer(time.Second)
	a := <-timer2.C;
	fmt.Printf("t2:%v\n", a);
	//for o := range timer2.C {
	//	fmt.Println("时间到",o)
	//}


	// 3.timer实现延时的功能
	//第一种方式
	time.Sleep(time.Second)
	//第二种方式
	timer3 := time.NewTimer(2 * time.Second)
	<-timer3.C
	fmt.Println("2秒到")
	//第三种方式
	<-time.After(2*time.Second)
	fmt.Println("2秒到")

	// 4.停止定时器
	timer4 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer4.C
		fmt.Println("定时器执行了")
	}()
	b := timer4.Stop()
	if b {
		fmt.Println("timer4已经关闭")
	}

	// 5.重置定时器 定时器设置的时间  可以重置
	timer5 := time.NewTimer(3 * time.Second)
	timer5.Reset(1 * time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)

	//Ticker 例子
	// 1.获取ticker对象
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//<-ticker.C
			i++
			fmt.Println(<-ticker.C)
			fmt.Println("data")
			//if i == 5 {
			//	//停止
			//	ticker.Stop()
			//}

			//for t := range ticker.C{
			//	fmt.Println("时间：",t);
			//}
		}
	}()
	for {
	}

}
