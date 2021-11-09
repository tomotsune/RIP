// @Title  main
// @Description
// @Author  haipinHu  09/11/2021 10:53
// @Update  haipinHu  09/11/2021 10:53
package config

import "RIP_UDP/src/common/common"

var (
	RouterName   = "B"
	RoutingTable = []common.Message{
		{Net: "N1", Dis: 5, Next: "A"},
		{Net: "N2", Dis: 3, Next: "C"},
		{Net: "N6", Dis: 6, Next: "F"},
		{Net: "N8", Dis: 4, Next: "E"},
	}
	//RoutingTable = []common.Message{
	//	{Net: "N1", Dis: 3, Next: "R4"},
	//	{Net: "N3", Dis: 4, Next: "R5"},
	//}
)
