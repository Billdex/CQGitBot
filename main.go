package main // import "CQGitBot"

import (
	"CQGitBot/conf"
	. "CQGitBot/githook"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func main() {
	//Load Config
	confPath := "./conf.json"
	err := conf.LoadConfig(confPath)
	if err != nil {
		log.Fatalln("Load Config Fail!")
	}
	log.Println("配置文件加载完成！", conf.Cfg)

	//Create a Gin router and run
	r := gin.Default()
	gitUri := conf.Cfg.Git.Uri
	r.POST(gitUri, EventParsing)

	port := strconv.FormatInt(conf.Cfg.Port, 10)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalln("gin.Run ERROR:", err)
	}

}
