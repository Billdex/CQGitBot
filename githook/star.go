package githook

import (
	"CQGitBot/models"
	"encoding/json"
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
	if eventInfo.Action == "created" {
		log.Println(eventInfo.Sender.Login, "给", eventInfo.Repository.Name, "点了个star")
		log.Println("该项目现在有", eventInfo.Repository.StargazersCount, "个star")
	} else if eventInfo.Action == "deleted" {
		log.Println(eventInfo.Sender.Login, "取消了对", eventInfo.Repository.Name, "的star")
	}
	return nil
}
