package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"code.byted.org/overpass/data_abtest_vm_framed/kitex_gen/abtest"
	"code.byted.org/overpass/data_abtest_vm_framed/rpc/data_abtest_vm_framed"
)

func main() {
	str := "{\n    \"3093234\":[\n        {\n            \"game_id\": \"ttb8efce5227c59661\",\n            \"game_name\": \"成语大富翁\",\n            \"introduction\": \"开店赚金币，致富街等你\"\n        },\n        {\n            \"game_id\": \"tta539d3843a134f3d\",\n            \"game_name\": \"金币农场\",\n            \"introduction\": \"种田种菜，金币满袋\"\n        },\n        {\n            \"game_id\": \"ttb416c2ab2b2d7549\",\n            \"game_name\": \"芒果斗地主\",\n            \"introduction\": \"天天打牌赚金币\"\n        },\n        {\n            \"game_id\": \"ttaf00d9bd303a8ddb\",\n            \"game_name\": \"当皇上金币版\",\n            \"introduction\": \"升官发财当皇帝啦\"\n        },\n        {\n            \"game_id\": \"tt604816f5ff6d8581\",\n            \"game_name\": \"金币消消乐\",\n            \"introduction\": \"消除水果赚金币\"\n        },\n        {\n            \"game_id\": \"tt96bee0030985821d\",\n            \"game_name\": \"水果爱消除\",\n            \"introduction\": \"畅快水果消除，金币多多\"\n        },\n        {\n            \"game_id\": \"tt442b3296db10857b\",\n            \"game_name\": \"金币连连看\",\n            \"introduction\": \"让麻将子碰撞消除吧！\"\n        },\n        {\n            \"game_id\": \"tte7f4914978cc31d302\",\n            \"game_name\": \"金币乐园\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tta3d5d5683cff178f\",\n            \"game_name\": \"找茬高高手\",\n            \"introduction\": \"找茬也能轻松赚金币！\"\n        },\n        {\n            \"game_id\": \"tt49d23f6fd6e8b047\",\n            \"game_name\": \"金币建筑师\",\n            \"introduction\": \"建房子，领取海量金币\"\n        },\n        {\n            \"game_id\": \"tt841ff2a21b254572\",\n            \"game_name\": \"天天CS金币版\",\n            \"introduction\": \"刺激枪战等你体验\"\n        }\n    ],\n    \"3093235\":[\n        {\n            \"game_id\": \"ttb8efce5227c59661\",\n            \"game_name\": \"成语大富翁\",\n            \"introduction\": \"开店赚金币，致富街等你\"\n        },\n        {\n            \"game_id\": \"tta539d3843a134f3d\",\n            \"game_name\": \"金币农场\",\n            \"introduction\": \"种田种菜，金币满袋\"\n        },\n        {\n            \"game_id\": \"ttb416c2ab2b2d7549\",\n            \"game_name\": \"芒果斗地主\",\n            \"introduction\": \"天天打牌赚金币\"\n        },\n        {\n            \"game_id\": \"ttaf00d9bd303a8ddb\",\n            \"game_name\": \"当皇上金币版\",\n            \"introduction\": \"升官发财当皇帝啦\"\n        },\n        {\n            \"game_id\": \"tt604816f5ff6d8581\",\n            \"game_name\": \"金币消消乐\",\n            \"introduction\": \"消除水果赚金币\"\n        },\n        {\n            \"game_id\": \"tt115f67e28e645f19\",\n            \"game_name\": \"金币点点消\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tt96bee0030985821d\",\n            \"game_name\": \"水果爱消除\",\n            \"introduction\": \"畅快水果消除，金币多多\"\n        },\n        {\n            \"game_id\": \"tt442b3296db10857b\",\n            \"game_name\": \"金币连连看\",\n            \"introduction\": \"让麻将子碰撞消除吧！\"\n        },\n        {\n            \"game_id\": \"tte7f4914978cc31d302\",\n            \"game_name\": \"金币乐园\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tta3d5d5683cff178f\",\n            \"game_name\": \"找茬高高手\",\n            \"introduction\": \"找茬也能轻松赚金币！\"\n        },\n        {\n            \"game_id\": \"tt49d23f6fd6e8b047\",\n            \"game_name\": \"金币建筑师\",\n            \"introduction\": \"建房子，领取海量金币\"\n        },\n        {\n            \"game_id\": \"tt841ff2a21b254572\",\n            \"game_name\": \"天天CS金币版\",\n            \"introduction\": \"刺激枪战等你体验\"\n        }\n    ],\n    \"3842610\":[\n        {\n            \"game_id\": \"ttb8efce5227c59661\",\n            \"game_name\": \"成语大富翁\",\n            \"introduction\": \"开店赚金币，致富街等你\"\n        },\n        {\n            \"game_id\": \"tta539d3843a134f3d\",\n            \"game_name\": \"金币农场\",\n            \"introduction\": \"种田种菜，金币满袋\"\n        },\n        {\n            \"game_id\": \"ttb416c2ab2b2d7549\",\n            \"game_name\": \"芒果斗地主\",\n            \"introduction\": \"天天打牌赚金币\"\n        },\n        {\n            \"game_id\": \"ttaf00d9bd303a8ddb\",\n            \"game_name\": \"当皇上金币版\",\n            \"introduction\": \"升官发财当皇帝啦\"\n        },\n        {\n            \"game_id\": \"tt604816f5ff6d8581\",\n            \"game_name\": \"金币消消乐\",\n            \"introduction\": \"消除水果赚金币\"\n        },\n        {\n            \"game_id\": \"tt96bee0030985821d\",\n            \"game_name\": \"水果爱消除\",\n            \"introduction\": \"畅快水果消除，金币多多\"\n        },\n        {\n            \"game_id\": \"tt442b3296db10857b\",\n            \"game_name\": \"金币连连看\",\n            \"introduction\": \"让麻将子碰撞消除吧！\"\n        },\n        {\n            \"game_id\": \"tte7f4914978cc31d302\",\n            \"game_name\": \"金币乐园\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tta3d5d5683cff178f\",\n            \"game_name\": \"找茬高高手\",\n            \"introduction\": \"找茬也能轻松赚金币！\"\n        },\n        {\n            \"game_id\": \"tt49d23f6fd6e8b047\",\n            \"game_name\": \"金币建筑师\",\n            \"introduction\": \"建房子，领取海量金币\"\n        },\n        {\n            \"game_id\": \"tt841ff2a21b254572\",\n            \"game_name\": \"天天CS金币版\",\n            \"introduction\": \"刺激枪战等你体验\"\n        }\n    ],\n    \"3842611\":[\n        {\n            \"game_id\": \"ttffff894260c7700b02\",\n            \"game_name\": \"终极钓手\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tta796d12b38ae2c5002\",\n            \"game_name\": \"金酷麻将\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tt85a4f3191a9d164402\",\n            \"game_name\": \"垂直火力\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tt72fc0441edc5f8c602\",\n            \"game_name\": \"水煮篮球\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"ttb8efce5227c59661\",\n            \"game_name\": \"成语大富翁\",\n            \"introduction\": \"开店赚金币，致富街等你\"\n        },\n        {\n            \"game_id\": \"tta539d3843a134f3d\",\n            \"game_name\": \"金币农场\",\n            \"introduction\": \"种田种菜，金币满袋\"\n        },\n        {\n            \"game_id\": \"ttb416c2ab2b2d7549\",\n            \"game_name\": \"芒果斗地主\",\n            \"introduction\": \"天天打牌赚金币\"\n        },\n        {\n            \"game_id\": \"ttaf00d9bd303a8ddb\",\n            \"game_name\": \"当皇上金币版\",\n            \"introduction\": \"升官发财当皇帝啦\"\n        },\n        {\n            \"game_id\": \"tt604816f5ff6d8581\",\n            \"game_name\": \"金币消消乐\",\n            \"introduction\": \"消除水果赚金币\"\n        },\n        {\n            \"game_id\": \"tt96bee0030985821d\",\n            \"game_name\": \"水果爱消除\",\n            \"introduction\": \"畅快水果消除，金币多多\"\n        },\n        {\n            \"game_id\": \"tt442b3296db10857b\",\n            \"game_name\": \"金币连连看\",\n            \"introduction\": \"让麻将子碰撞消除吧！\"\n        },\n        {\n            \"game_id\": \"tte7f4914978cc31d302\",\n            \"game_name\": \"金币乐园\",\n            \"introduction\": \"\"\n        },\n        {\n            \"game_id\": \"tta3d5d5683cff178f\",\n            \"game_name\": \"找茬高高手\",\n            \"introduction\": \"找茬也能轻松赚金币！\"\n        },\n        {\n            \"game_id\": \"tt49d23f6fd6e8b047\",\n            \"game_name\": \"金币建筑师\",\n            \"introduction\": \"建房子，领取海量金币\"\n        },\n        {\n            \"game_id\": \"tt841ff2a21b254572\",\n            \"game_name\": \"天天CS金币版\",\n            \"introduction\": \"刺激枪战等你体验\"\n        }\n    ]\n}"

	println(len(str))
	StructToString := func(v interface{}) string {
		byteStr, _ := json.Marshal(v)
		return string(byteStr)
	}
	ctx := context.Background()
	properties := configProperties{
		Token:    "aweme_lite",
		UID:      7167531,
		UIDType:  12,
		DeviceID: 4567897654567,
		AppID:    2329,
		//Channel:       "money",
		//Did:           105,
		//ClientVersion: "9.1.12",
		//ABLayerIds:        []int64{33153},
		//ABLayerIds:        []int64{33153, 25604, 25569, 25400, 25397},
		//ABLayerNamespaces: []string{"e_game"},
	}
	propertiesBytes, err := json.Marshal(properties)
	if err != nil {
		panic(err)
	}

	resp, err := data_abtest_vm_framed.GetAbVersions(ctx, string(propertiesBytes))
	if err != nil || resp == nil {
		log.Printf("service or parameters err\n")
		return
	}
	//log.Printf("resp:%s\n", StructToString(resp))
	for _, versionInfo := range resp.Info {
		if versionInfo.GetParameters() != "" && versionInfo.GetRspType() == abtest.VersionRspType_SUCCEED {
			ret := map[string]bool{}

			oriVersionStr := versionInfo.GetVersionName()
			if oriVersionStr != "" {
				arr := strings.Split(oriVersionStr, ",")
				for _, a := range arr {
					ret[a] = true
				}
			}

			oriClientStr := versionInfo.GetClientVersionName()
			if oriClientStr != "" {
				arr := strings.Split(oriClientStr, ",")
				for _, a := range arr {
					ret[a] = true
				}
			}
			//log.Printf("sversionStr:%s clientVersionStr:%s\n", StructToString(versionInfo.VersionName), StructToString(versionInfo.ClientVersionName))
			log.Printf("oriVersiontr:%v\n oriClientStr:%v\n", oriVersionStr, oriClientStr)
		} else {
			log.Printf("resp:%s\n", StructToString(resp))
		}
	}

	//parameters := &ABGrayReleaseResp{}
	//err = json.Unmarshal([]byte(version.GetParameters()), parameters)
	//if err != nil {
	//	fmt.Printf("err:%+v", err)
	//}
	//log.Printf("paramerters:%+v\n", parameters)
}

type configProperties struct {
	Token             string   `form:"token" json:"token"`
	UID               int64    `form:"uid" json:"uid"`
	UIDType           int      `form:"uid_type" json:"uid_type"`
	DeviceID          int64    `form:"device_id" json:"device_id"`
	AppID             int      `form:"app_id" json:"app_id"`
	Channel           string   `form:"channel" json:"channel"`
	Did               int64    `form:"did" json:"did"`
	ClientVersion     string   `form:"client_version" json:"client_version"`
	ABLayerIds        []int64  `form:"_ab_layer_ids" json:"_ab_layer_ids"`
	ABLayerNamespaces []string `form:"_ab_layer_namespaces" json:"_ab_layer_namespaces"`
}

type ABGrayReleaseResp struct {
	EGame GameGrayVersionInfo `form:"e_game" json:"e_game"`
}

type GameGrayVersionInfo struct {
	AppID   string `json:"app_id"`
	Version string `json:"version"`
}
