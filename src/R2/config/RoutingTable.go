// @Title  main
// @Description
// @Author  haipinHu  09/11/2021 10:53
// @Update  haipinHu  09/11/2021 10:53
package config

import "RIP_UDP/src/common/common"

var (
	RouterName   = "C"
	RoutingTable = []common.Message{
		{Net: "N1", Dis: 5, Next: "-"},
		{Net: "N2", Dis: 4, Next: "-"},
		// {Net: "N2", Dis: 16, Next: "B"},
		{Net: "N3", Dis: 8, Next: "-"},
		{Net: "N6", Dis: 4, Next: "-"},
		{Net: "N8", Dis: 3, Next: "-"},
	}
	//RoutingTable = []common.Message{
	//	{Net: "N1", Dis: 3, Next: "R1"},
	//	{Net: "N2", Dis: 4, Next: "R2"},
	//	{Net: "N3", Dis: 1, Next: "-"},
	//}
)
