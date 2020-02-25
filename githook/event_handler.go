package githook

import (
	"CQGitBot/conf"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/lunny/log"
)

//Verify X-GitHub-Signature and identify X-GitHub-Event
func EventParsing(c *gin.Context) {
	event := c.GetHeader("X-GitHub-Event")
	signature := c.GetHeader("X-Hub-Signature")
	log.Println("X-GitHub-Event: ", event)
	log.Println("X-Hub-Signature: ", signature)

	payloadBody, err := c.GetRawData()
	//Verify X-GitHub-Signature
	fromGitHub, err := verifySignature(payloadBody, signature, conf.Cfg.Git.Secret)
	if err != nil {
		log.Println("Verify signature fail: ", err)
		c.JSON(200, gin.H{
			"status_code": 1001,
			"message":     "Verify Signature Fail!",
		})
		return
	}
	if fromGitHub != true {
		c.JSON(200, gin.H{
			"status_code": 1002,
			"message":     "Signature Wrong!",
		})
		return
	}

	//Identify X-GitHub-Event
	switch event {
	case "ping":
		err = PingHandle(payloadBody)
	case "star":
		err = StarHandler(payloadBody)
	case "push":
		err = PushHandle(payloadBody)
	default:
		log.Println("Other event, do not handle it")
	}

	if err != nil {
		log.Error("Handler Event Error!", err)
		c.JSON(200, gin.H{
			"status_code": 1001,
			"message":     "Handler Event Error",
		})
	}
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     "ok",
	})
}

// Verify X-GitHub-Signature
func verifySignature(text []byte, signature string, secret string) (bool, error) {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write(text)
	bodySignature := "sha1=" + hex.EncodeToString(h.Sum(nil))
	return (signature == bodySignature), nil
}
