package line

import (
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/responses"
	"fmt"
)

func ReplyUser(jsonData responses.LineMessage) (string, error) {
	fmt.Println(jsonData.Events)
	return controller.ReplyUser(), nil
}
