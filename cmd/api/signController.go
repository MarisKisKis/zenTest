package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"zenTest/cmd/model"
)

func postSha512(c *gin.Context) {
	var sign model.Sign
	if err := c.ShouldBindJSON(&sign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	hexString := makeHexString(sign)
	c.JSON(http.StatusOK, hexString)
	return
}

func makeHexString(sign model.Sign) string {
	// sha512 encoding
	inputText := []byte(sign.Text)
	inputKey := []byte(sign.Key)
	h := hmac.New(sha512.New, inputKey)
	h.Write(inputText)
	results := h.Sum(nil)
	var result string
	for _, r := range results {
		result += fmt.Sprintf("%02X", r)
	}
	log.Println(result)
	// hex string encoding
	encodedStr := hex.EncodeToString([]byte(result))
	return encodedStr
}
