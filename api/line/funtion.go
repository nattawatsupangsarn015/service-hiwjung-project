package line

import (
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"
)

func ReplyUser(line responses.LineMessage) (string, error) {
	utils.LogWithTypeStruct(line)
	text := model.Text{
		Type: "text",
		Text: "ข้อความเข้ามา : " + line.Events[0].Message.Text + " ยินดีต้อนรับ : ",
	}

	message := model.ReplyMessage{
		ReplyToken: line.Events[0].ReplyToken,
		Messages: []model.Text{
			text,
		},
	}

	err := utils.ReplyMessageLine(message)
	if err != nil {
		return "", err
	}

	return "OK", nil
}
