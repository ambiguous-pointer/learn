package main

import (
	"fmt"
	"github.com/ambiguous-pointer/learn-go/three/socket"
)

func main() {
	//[Linux根目录分区扩容_reallywish的博客-CSDN博客](https://blog.csdn.net/YDJ18234492659/article/details/120887652)
	//
	//[Linux挂载磁盘（扩展根目录）_磁盘挂载根目录_没点is、的博客-CSDN博客](https://blog.csdn.net/qq_42730163/article/details/115460401#:~:text=Linux%E6%8C%82%E8%BD%BD%E7%A3%81%E7%9B%98%EF%BC%88%E6%89%A9%E5%B1%95%E6%A0%B9%E7%9B%AE%E5%BD%95%EF%BC%89%201%201.df%E5%91%BD%E4%BB%A4%E6%9F%A5%E7%9C%8B%E8%87%AA%E5%B7%B1%E7%9A%84%2F%E7%9B%AE%E5%BD%95%E5%B1%9E%E4%BA%8E%E5%93%AA%E4%B8%AA%E9%80%BB%E8%BE%91%E5%8D%B7%EF%BC%88%E6%88%91%E4%BB%AC%E8%A6%81%E5%81%9A%E7%9A%84%E6%98%AF%E5%AF%B9%E9%80%BB%E8%BE%91%E5%8D%B7%E6%89%A9%E5%B1%95%EF%BC%89%202%202.%E6%96%B0%E5%8A%A0%E4%B8%80%E5%9D%97%E7%A1%AC%E7%9B%98%EF%BC%8C%E6%88%91%E7%9A%84%E6%98%AFvdb%20%E5%85%88%E5%88%9B%E5%BB%BA%E7%89%A9%E7%90%86%E5%8D%B7%203%203.%E5%86%8D%E6%AC%A1%E6%9F%A5%E7%9C%8B%E7%A3%81%E7%9B%98%E6%83%85%E5%86%B5,5.%E6%9F%A5%E7%9C%8B%E6%A0%B9%E7%9B%AE%E5%BD%95%E6%89%80%E5%B1%9E%E7%9A%84%E5%8D%B7%E7%BB%84%E5%92%8C%E9%80%BB%E8%BE%91%E5%8D%B7%206%206.%E5%B0%86%E5%88%9A%E5%88%9B%E5%BB%BA%E7%9A%84%E7%89%A9%E7%90%86%E5%8D%B7%2Fdev%2Fvdb%E6%B7%BB%E5%8A%A0%E5%88%B0%E5%8D%B7%E7%BB%84%E4%B8%AD%207%207.%E5%B0%86%E5%8D%B7%E7%BB%84%E7%9A%84900GB%E6%89%A9%E5%B1%95%E7%BB%99%E6%A0%B9%E6%89%80%E5%9C%A8%E7%9A%84%E9%80%BB%E8%BE%91%E5%8D%B7%208%208.%E6%BF%80%E6%B4%BB%E4%BF%AE%E6%94%B9%E7%9A%84%E9%85%8D%E7%BD%AE%EF%BC%88%E6%AD%A4%E6%AD%A5%E9%AA%A4%E5%8F%AF%E8%83%BD%E4%BC%9A%E5%A4%B1%E8%B4%A5%EF%BC%89%20%E6%9B%B4%E5%A4%9A%E9%A1%B9%E7%9B%AE)
	//
	//[linux ln命令: 链接文件或目录_ln 文件夹_gsl68的博客-CSDN博客](https://blog.csdn.net/gsl68/article/details/6956842)
	//
	//[Linux中挂载详解以及mount命令用法_daydayup654的博客-CSDN博客](https://blog.csdn.net/daydayup654/article/details/78788310)

	callback := func(a int, b int) int {
		return a * b
	}
	callback1 := func(a int, b int) {
		c := a * b
		fmt.Println(c)
	}
	fun1 := func(f func(a int, b int) int) int {
		return f(1, 1) * f(1, 10)
	}
	fun2 := func(f func(a int, b int)) string {
		f(100, 100)
		return "OK"
	}
	// c
	go socket.RunSocketServer()
	fmt.Println(fun1(callback))
	fmt.Println(fun2(callback1))
	//go socket.RunSocketClient()
	socket.RunSocketClient()

}
