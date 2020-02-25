package githook

import (
	"CQGitBot/models"
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
		msg += fmt.Sprintf("该项目现在有%v个star\n", eventInfo.Repository.StargazersCount)

	} else if eventInfo.Action == "deleted" {
		msg += eventInfo.Sender.Login + "取消了star\n"
	}
	log.Println(msg)
	return nil
}
