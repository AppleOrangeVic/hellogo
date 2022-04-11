package main

import "fmt"

func main() {
	// v1.0.0 未引用任何第三方包，也未使用go module
	// fmt.Println("v1.0.0 未引用任何第三方包，也未使用go module")

	// v1.1.0 未引用任何第三方包，已开始使用go module，但没有任何外部依赖
	fmt.Println("v1.1.0 未引用任何第三方包，已开始使用go module，但没有任何外部依赖")
}
