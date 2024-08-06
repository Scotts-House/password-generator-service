package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes     = "0123456789"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		charsParam := c.DefaultQuery("chars", "true")
		numsParam := c.DefaultQuery("numbers", "true")
		specialsParam := c.DefaultQuery("specials", "true")
		lengthParam := c.DefaultQuery("length", "64")

		chars, err := strconv.ParseBool(charsParam)
		if err != nil {
			c.JSON(422, gin.H{
				"code":    422,
				"error":   err.Error(),
				"message": "Unable to parse chars param",
			})
		}
		nums, err := strconv.ParseBool(numsParam)
		if err != nil {
			c.JSON(422, gin.H{
				"code":    422,
				"error":   err.Error(),
				"message": "Unable to parse numbers param",
			})
		}
		specials, err := strconv.ParseBool(specialsParam)
		if err != nil {
			c.JSON(422, gin.H{
				"code":    422,
				"error":   err.Error(),
				"message": "Unable to parse specials param",
			})
		}
		length, err := strconv.ParseInt(lengthParam, 10, 64)
		if err != nil {
			c.JSON(422, gin.H{
				"code":    422,
				"error":   err.Error(),
				"message": "Unable to parse length param",
			})
		}

		if chars == false && nums == false && specials == false {
			c.JSON(422, gin.H{
				"code":    422,
				"error":   "Invalid character restrictions",
				"message": "Characters, numbers, and specials cannot all be false.  You must allow at least one.",
			})
		}

		password := generatePassword(length, chars, specials, nums)

		c.JSON(http.StatusOK, gin.H{
			"password": password,
		})
	})
	err := r.Run()
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func generatePassword(length int64, useLetters bool, useSpecial bool, useNum bool) string {
	b := make([]string, length)

	opts := []string{}

	if useLetters == true {
		for _, l := range letterBytes {
			opts = append(opts, string(l))
		}
	}

	if useNum == true {
		for _, n := range numBytes {
			opts = append(opts, string(n))
		}
	}

	if useSpecial == true {
		for _, s := range specialBytes {
			opts = append(opts, string(s))
		}
	}

	for i := range b {
		rint, _ := rand.Int(rand.Reader, big.NewInt(int64(len(opts))))
		b[i] = opts[rint.Int64()]
	}
	return strings.Join(b, "")
}
