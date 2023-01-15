package line

import (
	"errors"
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReplyUser(line responses.LineMessage) (string, error) {
	utils.LogWithTypeStruct(line)
	lineId := line.Events[0].Source.UserID
	findUser, err := controller.FindUserByLineId(lineId)
	if err != nil {
		return "", err
	}

	if findUser == nil {
		newUser := &model.Users{
			ID:        primitive.NewObjectID(),
			LineId:    lineId,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = controller.CreateUser(newUser)
		if err != nil {
			return "", err
		}
	}

	var replyText string
	message := line.Events[0].Message

	if message.Type != "text" {
		replyText = "ตอนนี้หิวจังยังรองรับแค่การพิมปกติน้า 🥺"
	} else {
		findReply, err := controller.FindChatbotMessages(message.Text)
		if err != nil {
			return "", err
		}

		if findReply != nil {
			replyStruct, ok := findReply.(model.ChatbotReply)
			if !ok {
				return "", errors.New("Cannot convert structure")
			}

			replyText = replyStruct.Reply
		} else {
			replyText = "น้องหิวจังยังไม่เรียนรู้ ขอโทษด้วยน้า 😥"
		}
	}

	text := model.Text{
		Type: "text",
		Text: replyText,
	}

	messageReply := model.ReplyMessage{
		ReplyToken: line.Events[0].ReplyToken,
		Messages: []model.Text{
			text,
		},
	}

	err = utils.ReplyMessageLine(messageReply)
	if err != nil {
		return "", err
	}

	return "OK", nil
}
