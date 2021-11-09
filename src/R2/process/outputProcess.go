package process

import (
	config2 "RIP_UDP/src/R1/config"
	"RIP_UDP/src/R2/config"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func Transmit() {
	conn, err := net.Dial("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	for {
		routingTable, err := json.Marshal(config.RoutingTable)
		if err != nil {
			fmt.Println("json.Marshal err:", err)
			return
		}
		conn.Write(routingTable)
		fmt.Printf("%v=>%v: %v\n", config.RouterName, config2.RouterName, config.RoutingTable)
		time.Sleep(time.Second * 3)
	}
}
