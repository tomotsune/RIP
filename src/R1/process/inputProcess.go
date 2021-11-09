package process

import (
	"RIP_UDP/src/R1/config"
	config2 "RIP_UDP/src/R2/config"
	"RIP_UDP/src/common/common"
	"encoding/json"
	"fmt"
	"net"
)

func Sense() {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP err:", err)
			return
		}
		// fmt.Printf("接收到客户端[%s]：%s", raddr, string(buf[:n]))
		var routingTable []common.Message
		err = json.Unmarshal(buf[:n], &routingTable)
		if err != nil {
			fmt.Println("json.Unmarshal err:", err)
			return
		}
		fmt.Printf("%v<=%v: %v\n", config.RouterName, config2.RouterName, routingTable)
		for _, newR := range routingTable {
			if newR.Dis++; newR.Dis == 17 {
				newR.Dis--
			}
			newR.Next = config2.RouterName
			flag := true
			for idx, oldR := range config.RoutingTable {

				if newR.Net == oldR.Net {
					flag = false
					if newR.Next == oldR.Next {
						config.RoutingTable[idx].Net = newR.Net
						config.RoutingTable[idx].Dis = newR.Dis
						config.RoutingTable[idx].Next = newR.Next
					} else {
						if newR.Dis < oldR.Dis {
							config.RoutingTable[idx].Net = newR.Net
							config.RoutingTable[idx].Dis = newR.Dis
							config.RoutingTable[idx].Next = newR.Next
						}
					}
				}
			}
			if flag {
				config.RoutingTable = append(config.RoutingTable, newR)
			}
		}

		fmt.Printf("%v更新后的路由表:%v\n", config.RouterName, config.RoutingTable)
	}
}
