package githook

import (
	"CQGitBot/conf"
	"CQGitBot/models"
	"CQGitBot/qqbot"
	"encoding/json"
	"github.com/lunny/log"
	"strings"
)

//handle push event
func PushHandle(payload []byte) (err error) {
	var eventInfo models.PushPayload
	err = json.Unmarshal(payload, &eventInfo)
	if err != nil {
		log.Error("Unmarshal payload fail!", err)
		return err
	}
	log.Println(eventInfo)

	msg := "项目：" + eventInfo.Repository.Name + "\n"
	msg += "操作：@" + eventInfo.Pusher.Name + " 提交了代码"
	ref := strings.Split(eventInfo.Ref, "/")
	if len(ref) == 3 {
		msg += "到 " + ref[2] + " 分支"
	}
	msg += "\n"
	msg += "最新提交：" + eventInfo.HeadCommit.Url + "\n"
	msg += "\n"
	for pos, commit := range eventInfo.Commits {
		msg += "[" + commit.Timestamp.Format("01-02 15:04:05") + "] " + commit.Id[0:7] + "\n"
		msg += commit.Message
		if pos != len(eventInfo.Commits)-1 {
			msg += "\n"
		}
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
