// @Title  main
// @Description
// @Author  haipinHu  09/11/2021 10:45
// @Update  haipinHu  09/11/2021 10:45
package main

import (
	"RIP_UDP/src/R2/config"
	"RIP_UDP/src/R2/process"
	"fmt"
)

func main() {
	fmt.Printf("%v路由器初始路由表:%v\n", config.RouterName, config.RoutingTable)
	go process.Sense()
	process.Transmit()
}
