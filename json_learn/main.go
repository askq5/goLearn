package main

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

const (
	jsonTest1 = iota + 1
	jsonTest2
)

type testStruct struct {
	AmountType int32 `thrift:"AmountType,1,required" json:"AmountType"`
	Amount     int64 `thrift:"Amount,2,required" json:"Amount"`
}

type SixtySecOneOfThreeResp struct {
	Type         int32       `thrift:"Type,1,required" json:"Type"`
	TaskStrategy *testStruct `thrift:"TaskStrategy,2" json:"TaskStrategy,omitempty"`
	PassCard     *testStruct `thrift:"PassCard,3" json:"PassCard,omitempty"`
	AmountReward *testStruct `thrift:"AmountReward,4" json:"AmountReward,omitempty"`
	BaseResp     *testStruct `thrift:"BaseResp,255" json:"BaseResp,omitempty"`
}

func main() {
	//testStr := "{\"AmountReward\":{\"Amount\":10,\"AmountType\":1001},\"PassCard\":{},\"TaskStrategy\":\"{}\",\"Type\":4}"
	//testStru := &SixtySecOneOfThreeResp{}
	//json.Unmarshal([]byte(testStr), testStru)
	//fmt.Printf("printf1:%+v\n", testStru)
	//printStr, _ := json.Marshal(testStru)
	//fmt.Printf("printf2:%+v\n", string(printStr))

	// str2 := "{\"ttgame\":{\"find_tab_recommend\":\"1\"},\"ttgame_data_find_more_game\":{\"test_info\":\"find_more_game_test_info\",\"post_rank\":{\"module_groups_order\":{\"enable_module_groups_order\":true,\"req_module_game_count\":7,\"req_module_count\":8,\"module_rank_count\":8,\"module_rank_weight\":[1,1,1,1,1,1,1,1]}}}}\n"
	str2 := "{\"final_game_list\":[\"tt604816f5ff6d8581\",\"ttaf00d9bd303a8ddb\",\"ttb416c2ab2b2d7549\",\"tta539d3843a134f3d\",\"tta796d12b38ae2c5002\",\"ttb7a9c11428830de2\",\"tt8b26fd9b141ff83a02\",\"tt115f67e28e645f19\",\"ttb8efce5227c59661\",\"tte7f4914978cc31d302\",\"tt96bee0030985821d\",\"tt442b3296db10857b\",\"tta3d5d5683cff178f\",\"tt49d23f6fd6e8b047\",\"tt841ff2a21b254572\",\"tt3bea4d57a4d1e47d\"]}"
	dataValue := gjson.Get(str2, "ttgame_data_find_more_game").Value()
	dataMap := gjson.Get(str2, "ttgame_data_find_more_game").Map()
	dataStr := gjson.Get(str2, "ttgame_data_find_more_game").String()
	strMap := make(map[string][]string)
	json.Unmarshal([]byte(str2), &strMap)
	fmt.Println(strMap)
	testStr, _ := json.Marshal(strMap["final_game_list"])
	fmt.Printf("%s\n", testStr)
	fmt.Printf("dataValue:%s\n dataMap:%s\n dataStr1:%s\n dataStr2:%s\n", ToString(dataValue), ToString(dataMap), ToString(dataStr), dataStr)
	//a := jsonTest2
	//aJson, err := json.Marshal(jsonTest2)
	//if err != nil {
	//	fmt.Errorf("marshall a err")
	//}
	//fmt.Printf("%+v", aJson)
}

func ToString(val interface{}) string {
	if v, ok := val.(string); ok {
		return v
	}
	printStr, _ := json.Marshal(val)
	return string(printStr)
}
