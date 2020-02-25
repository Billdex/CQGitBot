package githook

import (
	"CQGitBot/models"
	"encoding/json"
	"github.com/lunny/log"
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

	msg := "[" + eventInfo.Repository.Name + " | push]\n"
	msg += eventInfo.Pusher.Name + "提交了代码\n"
	for _, commit := range eventInfo.Commits {
		msg += commit.Message + " [" + commit.Timestamp.Format("2006-01-02 15:04:05") + "]\n"
	}
	log.Println(msg)
	return nil
}
