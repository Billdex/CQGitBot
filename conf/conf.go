package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type config struct {
	Port int64 `json:"port"`
	Git  git   `json:"git"`
	QQ   qq    `json:"qq"`
}
type git struct {
	Uri    string `json:"uri"`
	Secret string `json:"secret"`
}
type qq struct {
	Uri     string   `json:"uri"`
	CQPort  int64    `json:"cq_port"`
	GroupId []string `json:"group_id"`
}

var Cfg config

//Load configs from json file
func LoadConfig(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("load config ", path, " failed: ", err)
		return err
	}
	err = json.Unmarshal(buf, &Cfg)
	if err != nil {
		log.Println("config unmarshal json failed: ", err)
		return err
	}
	return nil
}
