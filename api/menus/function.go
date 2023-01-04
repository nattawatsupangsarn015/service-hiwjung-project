package menus

import (
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/responses"
)

func GetAllMenus() string {
	return controller.GetAllMenus()
}

func CreateMenu(menu responses.MenuRequestCreate) (interface{}, error) {
	err := controller.CreateMenu(menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}
