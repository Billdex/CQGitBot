package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	_, err := os.Lstat(path)
	if os.IsNotExist(err) {
		log.Println("配置文件不存在！")
		var confFile = `{
		  "port": 7920,
		  "git": {
			"uri": "/git",
			"secret": ""
		  },
		  "qq": {
			"uri": "/qq",
			"cq_port": 5700,
			"group_id": [
			  ""
			]
		  }
		}`
		if ioutil.WriteFile(path, []byte(confFile), 0644) != nil {
			log.Fatalln("创建配置文件失败！")
			return err
		} else {
			log.Println("创建配置文件成功")
		}
	} else {
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
	}
	return nil
}
