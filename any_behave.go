package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

func replaceSpace(s string) string {
	strByte := make([]byte, 0)
	replStr := "%20"
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			strByte = append(strByte, s[i])
		} else {
			strByte = append(strByte, *(*[]byte)(unsafe.Pointer(&replStr))...)
		}
	}
	return *(*string)(unsafe.Pointer(&strByte))
}

type testStruct struct {
}

type testExtra struct {
	Module   string  `json:"module"`
	GameList []int32 `json:"gameList"`
}

type TCCBannerBetterConf struct {
	BannerDefaultDays int64                `json:"banner_default_days"`
	BannerCountLimit  int32                `json:"banner_count_limit"`
	BannerList        []ResourceBetterItem `json:"banner_list"`
}

type ResourceBetterItem struct {
	// 必填信息
	BannerId      int64           `json:"banner_id"`      // banner唯一id
	MainTitle     string          `json:"main_title"`     // 主标题 游戏名称
	SubTitle      string          `json:"sub_title"`      // 副标题 运营推荐语
	Url           string          `json:"url"`            // 图片链接
	Schema        string          `json:"schema"`         // 图片跳转游戏链接
	Order         int64           `json:"order"`          // 轮换顺序
	BeginTime     int64           `json:"begin_time"`     // 开始时间，当天配置
	VisibleDays   int64           `json:"visible_days"`   // 弹窗可见天数
	VisibleAid    map[int32]bool  `json:"visible_aid"`    // 可见端
	VisibleSystem map[string]bool `json:"visible_system"` // 可见系统 ios/安卓
	Status        bool            `json:"status"`         // 状态 1-可用 0-不可用
	// 选填信息
	GameId   string   `json:"game_id"`   // 游戏id
	GameName string   `json:"game_name"` // 游戏名字
	GameTag  []string `json:"game_tag"`  // - 标签（上限两个）：新游上线、编辑精选
	GameType string   `json:"game_type"` // - 游戏类型：悬疑、模拟经营
	GameSpot string   `json:"game_spot"` // - 游戏亮点：新游、头部cp、活动
	CoinNum  int64    `json:"coin_num"`  // 金币发放数量
	// 运营信息
	Operator    string `json:"operator"`     // 填写人
	ResourceDes string `json:"resource_des"` // 其他备注
}

func StructToJson(newStruct interface{}) {
	str, _ := json.Marshal(newStruct)
	fmt.Println(string(str))
}

type LotteryItem struct {
	Id    string `json:"id"`
	Count int64  `json:"count"`
	Type  string `json:"type"`
	Index int    `json:"index"` //转盘的位置号
	Desc  string `json:"desc"`
	Name  string `json:"name"`
}

func NullToJson() {
	result := &LotteryItem{}
	dataStr := "null"
	err := json.Unmarshal([]byte(dataStr), result)
	fmt.Println(err)
	fmt.Println(result)
}

const (
	Askq = "askq"
	Skq  = "skq"
)

type ParalStruct struct {
	Name string
}

// 配置列表
type TCCNewGameConfList []NewGamePopConf

type ProgressRewardConfItem struct {
	TaskFinishCnt     int32 `json:"task_finish_cnt"`     // 任务完成数量
	ProgressRewardNum int64 `json:"progress_reward_num"` // 对应完成任务奖励
}

type PopTaskConfItem struct {
	TaskId          int64  `json:"task_id"`          // 任务id
	TaskKey         string `json:"task_key"`         // 任务key
	Weight          int32  `json:"weight"`           // 任务权重
	MainTitle       string `json:"main_title"`       // 主标题
	SubTitle        string `json:"sub_title"`        // 次标题
	RewardNum       int64  `json:"reward_num"`       // 奖励数量
	TaskLimit       int32  `json:"task_limit"`       // 任务完成次数限制（目前一般都是一次）
	RefreshTime     int32  `json:"refresh_time"`     // 任务刷新时间，(天级，模块级)
	TaskType        int32  `json:"task_type"`        // 任务类型（用来区分任务完成模式，GamePush PlaytimePull）
	TaskRequirement int64  `json:"task_requirement"` // 任务完成条件
}

// TCC配置
type NewGamePopConf struct {
	ConfId            string                    `json:"conf_id"` // 配置id
	GameName          string                    `json:"game_name"`
	GameId            string                    `json:"game_id"`
	GameIcon          string                    `json:"game_icon"`          // 游戏icon
	ModuleUrl         string                    `json:"module_url"`         // 模块图片
	SubTitle          string                    `json:"sub_title"`          // 模块副标题
	BeginTime         int64                     `json:"begin_time"`         // 上线日期  需要确定格式
	VisibleDays       int64                     `json:"visible_days"`       // 上线时长
	VisibleApp        map[int32]map[string]bool `json:"visible_app"`        // 可见端 是否要做成针对某个appid的某个操作系统 key:appId  val:key:ios/android val:bool
	TaskShowLimit     int32                     `json:"task_show_limit"`    // 任务横滑展示数量
	SilenceDaysLimit  int64                     `json:"silence_days_limit"` // 曝光多少天后不点击下掉该模块
	ProgressTaskTitle string                    `json:"progress_task_title"`
	TaskList          []PopTaskConfItem         `json:"task_list"` // 新游配置的任务列表
}

func PopJson() {
	popItem := PopTaskConfItem{}
	visibleApp := map[int32]map[string]bool{
		2329: map[string]bool{
			"android": true,
		},
	}
	confList := TCCNewGameConfList{NewGamePopConf{
		VisibleApp: visibleApp,
		TaskList:   []PopTaskConfItem{popItem},
	}}
	str, _ := json.Marshal(confList)
	fmt.Println(string(str))
}

func main() {
	PopJson()
	testStruct1 := ParalStruct{
		Name: "b",
	}

	go func() {
		for {
			testStruct1.Name = Askq
			time.Sleep(1)
		}
	}()

	go func() {
		for {
			testStruct1.Name = Skq
			time.Sleep(1)
		}
	}()
	for {
		tempStr, _ := json.Marshal(testStruct1)
		fmt.Println(tempStr)
	}
	//var nilPointer *int64
	dataStr, _ := json.Marshal(nil)
	fmt.Println(dataStr)

	tccStruct := TCCBannerBetterConf{
		BannerList: []ResourceBetterItem{ResourceBetterItem{
			VisibleAid:    map[int32]bool{2329: true},
			VisibleSystem: map[string]bool{"android": true},
			GameTag:       []string{"新游上线", "编辑精选"},
		}},
	}

	StructToJson(tccStruct)
	//strArray1 := []string{"1", "2", "3", "4"}
	//strArray2 := []string{"3", "2", "4", "1", "5"}
	//_, newUsedGames := funk.DifferenceString(strArray1, strArray2)
	//joinedGames := funk.InnerJoinString(strArray1, strArray2)
	//newUsedGames = append(newUsedGames, joinedGames...)
	//fmt.Println(newUsedGames)

	//module := "gold"
	//gameList := []int32{1, 2, 3, 4}
	//extra := make(map[string]interface{})
	//extra["module"] = module
	//extra["game_list"] = gameList
	//extra1 := testExtra{
	//	Module:   module,
	//	GameList: gameList,
	//}
	//bytes1, _ := json.Marshal(extra)
	//bytes2, _ := json.Marshal(extra1)
	//fmt.Println(string(bytes1))
	//fmt.Println(string(bytes2))

	var test interface{}
	test = &testStruct{}
	fmt.Println(reflect.TypeOf(test), reflect.TypeOf(test).String(), reflect.TypeOf(test).Elem().Name())
	fmt.Println(reflect.ValueOf(replaceSpace), reflect.ValueOf(replaceSpace).String(), reflect.ValueOf(replaceSpace).Pointer())

	a := "foo"
	rV := reflect.ValueOf(&a)
	rV = rV.Elem()
	rV.SetString("new foo")
	fmt.Println(a)

	var v *int
	bytes, _ := json.Marshal(v)
	fmt.Println(string(bytes))
	map1 := make(map[int]int)
	map2 := map1
	for i := 0; i < 10000; i++ {
		map2[i] = i
	}
	fmt.Println(map1)
	fmt.Println(map2)
	i, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		log.Panicln(i)
	} else {
		log.Println(i, err)
	}
	constFeature1()
	constFeature2()
}

func constFeature1() {
	const test1 = "test1"
	fmt.Println(test1)
}

func constFeature2() {
	const test1 = "test2"
	fmt.Println(test1)
}
