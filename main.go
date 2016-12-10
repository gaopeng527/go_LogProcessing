// Log Processing project main.go
package main

import (
	"fmt"
	"logs"
	"strconv"
)

func main() {
	i := "12.3"
	logs.Logger.Info("开始将字符串转换成整数")
	j, err := strconv.Atoi(i)
	fmt.Println(j)
	logs.Logger.Critical("转换出错:", err)
}
