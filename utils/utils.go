package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HandleBadRequest(c *gin.Context, structure interface{}) {
	if err := c.BindJSON(&structure); err != nil {
		c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
}

func HandleResponse(c *gin.Context, response interface{}, err error, status int, statusError int) {
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(statusError, err.Error())
	} else {
		c.JSON(status, response)
	}
}

func ConvertToStruct(input interface{}, output interface{}) error {
	jsonBytes, ok := input.([]byte)
	if !ok {
		return errors.New("Cannot convert data to byte.")
	}

	err := json.Unmarshal(jsonBytes, &output)
	if err != nil {
		return err
	}

	return nil
}

func SignJwt(secret []byte, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SECRET_KEY := os.Getenv("SECRET_KEY")
		tokenString := c.GetHeader("Authorization")
		splitToken := strings.Split(tokenString, " ")
		if len(splitToken) != 2 {
			HandleResponse(c, nil, errors.New("Unauthorize"), 0, 401)
			return
		}

		token, err := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRET_KEY), nil
		})
		if err != nil {
			HandleResponse(c, nil, errors.New("Unauthorize"), 0, 401)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			HandleResponse(c, nil, errors.New("Unauthorize"), 0, 401)
			return
		}

		c.Next()
	}
}

func LogWithTypeStruct(data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(jsonData))
}

func ReplyMessageLine(Message model.ReplyMessage) error {
	value, _ := json.Marshal(Message)

	CHANNEL_TOKEN := os.Getenv("CHANNEL_TOKEN")
	url := os.Getenv("API_LINE_REPLY")

	var jsonStr = []byte(value)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+CHANNEL_TOKEN)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	// log.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// log.Println("response Body:", string(body))

	return err
}
