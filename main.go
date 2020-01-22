package main	// import "CQGitBot"

import (
	"CQGitBot/conf"
	"github.com/gin-gonic/gin"
	. "CQGitBot/githook"
	"log"
	"strconv"
)

func main(){
	//Load Config
	confPath := "./conf.json"
	err := conf.LoadConfig(confPath)
	if err != nil{
		log.Fatalln("Load Config Fail!")
		return
	}

	//Create a Gin router and run
	r := gin.Default()
	gitUri := conf.Cfg.Git.Uri
	r.POST(gitUri, EventParsing)

	port := strconv.FormatInt(conf.Cfg.Port, 10)
	err = r.Run(":" + port)
	if err != nil{
		log.Fatalln("gin.Run ERROR:", err)
	}

}
