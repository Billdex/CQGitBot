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

	msg := "[" + eventInfo.Repository.Name + " | star]\n"
	if eventInfo.Action == "created" {
		msg += eventInfo.Sender.Login + "点了个star\n"
		msg += fmt.Sprintf("该项目现在有%v个star", eventInfo.Repository.StargazersCount)

	} else if eventInfo.Action == "deleted" {
		msg += eventInfo.Sender.Login + "取消了star"
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
