package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	// v1.0.0 未引用任何第三方包，也未使用go module
	// fmt.Println("v1.0.0 未引用任何第三方包，也未使用go module")

	// v1.1.0 未引用任何第三方包，已开始使用go module，但没有任何外部依赖
	// fmt.Println("v1.1.0 未引用任何第三方包，已开始使用go module，但没有任何外部依赖")

	// v1.2.0 引用了第三方包，并更新了项目依赖
	fmt.Println("v1.2.0 引用了第三方包，并更新了项目依赖")

	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}
