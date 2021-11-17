package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/galaxy-future/BridgX/internal/constants"
	jsoniter "github.com/json-iterator/go"
)

const (
	LarkMsgType_Text       = "text"        //富文本
	LarkMsgType_Post       = "post"        //富文本
	LarkMsgType_Image      = "image"       //图片
	LarkMsgTyp_ShareChat   = "share_chat"  //分享群名片
	LarkMsgTyp_Interactive = "interactive" //消息卡片
)

/*
// err response
{
    "code": 19002,
    "msg": "params error, msg_type need"
}
// success response
{
    "Extra": null,
    "StatusCode": 0,
    "StatusMessage": "success"
}
*/
type LarkMsgResponse struct {
	Code          int64  `json:"code"`
	Msg           string `json:"msg"`
	StatusCode    int64  `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

//func LarkAlarmForAllEnv(ctx context.Context, hookID string, title string, text string) (err error) {
//	//note := "线上"
//	//if !env.IsProduct() {
//	//	note = "测试"
//	//}
//	note := "-"
//	return LarkAlarmBase(ctx, hookID, title, note, text)
//}

func LarkAlarm(ctx context.Context, hookID string, title string, text string) (err error) {
	//if !env.IsProduct() {
	//	return nil
	//	//hookID = constant.LARK_ALARM_PAY_TEST
	//}
	//note := "线上"
	//note, _ := LocalIp()
	return LarkAlarmBase(ctx, hookID, title, text)
}

func LarkAlarmBase(ctx context.Context, hookID string, title string, text string) (err error) {
	//logId, _ := kitutil.GetCtxLogID(ctx)
	url := initLarkUrl(hookID)
	textMap := make(map[string]interface{})
	//headerText := fmt.Sprintf("[%s]", note)
	ip, _ := LocalIp()
	gapLine := "——————————————————————————"
	batchNo := ctx.Value(constants.BatchNo)
	footerText := fmt.Sprintf("%s\n任务批次[%v]\nts: %v\nip: %v", gapLine, batchNo, CurrentTime(), ip)
	textMap["text"] = fmt.Sprintf("%s\n%s\n%s", title, text, footerText)

	postMap := make(map[string]interface{})
	postMap["msg_type"] = LarkMsgType_Text
	postMap["content"] = textMap

	jsonBytes, err := jsoniter.Marshal(postMap)
	if err != nil {
		err = fmt.Errorf("func:LarkAlarm | json marshal:%v", err.Error())
		return err
	}
	httpRes, err := HttpPostJsonDataT(url, jsonBytes, 2)
	if err != nil {
		err = fmt.Errorf("func:LarkAlarm | http err:%v", err.Error())
		return err
	}
	msgRes := LarkMsgResponse{
		StatusCode: -1,
	}
	err = json.Unmarshal(httpRes, &msgRes)
	if err != nil {
		err = fmt.Errorf("func:LarkAlarm | http unmarshal:%v", err.Error())
		return err
	}
	if msgRes.StatusCode != 0 {
		err = fmt.Errorf("func:LarkAlarm | lark msg failed [%+v]| %s", msgRes, jsonBytes)
		return err
	}
	return nil
}

func initLarkUrl(hookID string) string {
	return fmt.Sprintf("https://open.feishu.cn/open-apis/bot/v2/hook/%s", hookID)
}

func LocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	var ip string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}

	return ip, nil
}
