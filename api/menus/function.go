package menus

import (
	"example/service-hiwjung-project/controller"
)

func GetAllMenus() string {
	return controller.GetAllMenus()
}

func CreateMenu(menu interface{}) (interface{}, error) {
	return menu, nil
}
