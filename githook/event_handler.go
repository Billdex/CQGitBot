package githook

import (
	"CQGitBot/conf"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
)

//Verify X-GitHub-Signature and identify X-GitHub-Event
func EventParsing(c *gin.Context){
	event := c.GetHeader("X-GitHub-Event")
	signature := c.GetHeader("X-Hub-Signature")
	log.Println("X-GitHub-Event: ", event)
	log.Println("X-Hub-Signature: ", signature)

	//Verify X-GitHub-Signature
	fromGitHub, err := verifySignature(c)
	if err != nil{
		log.Println("Verify signature fail: ", err)
		c.JSON(200, gin.H{
			"status_code": 1001,
			"message": "Verify Signature Fail!",
		})
		return
	}
	if fromGitHub != true{
		c.JSON(200, gin.H{
			"status_code": 1002,
			"message": "Signature Wrong!",
		})
		return
	}

	//Identify X-GitHub-Event
	switch event {
	case "star":
		StarHandler(c)
	default:
		log.Println("Other event, do not handle it")
	}

	c.JSON(200, gin.H{
		"status_code": 200,
		"message": "ok",
	})
}

// Verify X-GitHub-Signature
func verifySignature(c *gin.Context) (bool, error) {
	payLoadBody, err := c.GetRawData()
	if err != nil {
		return false, err
	}
	headerSignature := c.GetHeader("X-Hub-Signature")
	bodySignature := hmacSha1(payLoadBody ,conf.Cfg.Git.Secret)
	return (headerSignature == bodySignature), nil
}

// hmac-sha1
func hmacSha1(text []byte, secret string) (string) {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(text)
	return "sha1=" + hex.EncodeToString(h.Sum(nil))
}