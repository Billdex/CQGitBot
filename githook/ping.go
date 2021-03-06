package githook

import (
	"CQGitBot/conf"
	"CQGitBot/models"
	"CQGitBot/qqbot"
	"encoding/json"
	"github.com/lunny/log"
	"strconv"
)

func PingHandle(payload []byte) (err error) {
	var eventInfo models.PingPayload
	err = json.Unmarshal(payload, &eventInfo)
	if err != nil {
		log.Error("Unmarshal payload fail!", err)
		return err
	}
	log.Println(eventInfo)

	msg := "项目：" + eventInfo.Repository.Name + "\n"
	msg += "事件：@" + eventInfo.Sender.Login + " 添加了一个WebHook\n"
	msg += "Hook ID:" + strconv.FormatInt(eventInfo.HookId, 10) + "\n"
	msg += "仓库地址: " + eventInfo.Repository.HtmlUrl

	log.Println(msg)
	for _, groupId := range conf.Cfg.QQ.GroupId {
		log.Println("尝试发消息至QQ群：", groupId)
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
