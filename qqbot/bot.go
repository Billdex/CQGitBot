package qqbot

import (
	"CQGitBot/conf"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//QQ user info
type User struct {
	//private info
	Id       int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Age      int64  `json:"age"`
	Sex      string `json:"sex"`

	//group member info
	Area  string `json:"area"`
	Card  string `json:"card"`
	Level string `json:"level"`
	Role  string `json:"role"`
	Title string `json:"title"`

	//anonymous user info in group
	AnonymousId int64 `json:"anonymousId"`
}

//Message data from cqhttp
type CQMsg struct {
}

type GroupMsg struct {
	GroupId    string `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape bool   `json:"auto_escape"`
}

func SendMsg(byteMsg []byte, url string) (err error) {
	request, err := http.NewRequest("POST", url, bytes.NewReader(byteMsg))
	if err != nil {
		return nil
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("CoolQ Response Info:", string(body))
	return nil
}

func SendGroupMsg(msg GroupMsg) (err error) {
	byteMsg, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	url := "http://localhost:" + strconv.FormatInt(conf.Cfg.QQ.CQPort, 10) + "/send_group_msg"
	err = SendMsg(byteMsg, url)
	if err != nil {
		return err
	}
	log.Println("Send Group Msg Success!", msg)
	return nil
}
