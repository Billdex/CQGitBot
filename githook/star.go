package githook

import (
	"CQGitBot/conf"
	"CQGitBot/models"
	"CQGitBot/qqbot"
	"encoding/json"
	"fmt"
	"github.com/lunny/log"
)

//handle star event
func StarHandler(payload []byte) (err error) {
	var eventInfo models.StarPayload
	err = json.Unmarshal(payload, &eventInfo)
	if err != nil {
		log.Error("Unmarshal payload fail!", err)
		return err
	}
	log.Println(eventInfo)

	msg := "项目：" + eventInfo.Repository.Name + "\n"
	if eventInfo.Action == "created" {
		msg += "事件：@" + eventInfo.Sender.Login + " 点了个 ★\n"
		msg += fmt.Sprintf("该项目有 %v 个 ★ 啦！加油鸭！", eventInfo.Repository.StargazersCount)

	} else if eventInfo.Action == "deleted" {
		msg += "事件：@" + eventInfo.Sender.Login + " 取消了 ☆\n"
		msg += "再接再厉鸭！"
	}
	log.Println(msg)
	for _, groupId := range conf.Cfg.QQ.GroupId {
		err = qqbot.SendGroupMsg(qqbot.GroupMsg{
			GroupId:    groupId,
			Message:    msg,
			AutoEscape: true,
		})
		if err != nil {
			log.Error("Send Group Msg Fail!", err)
		}
	}
	return nil
}
