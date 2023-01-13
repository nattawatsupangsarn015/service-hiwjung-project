package admins

import (
	"errors"
	"example/service-hiwjung-project/controller"
	"example/service-hiwjung-project/model"
	"example/service-hiwjung-project/responses"
	"example/service-hiwjung-project/utils"
	"fmt"
	"os"

	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(admin responses.AdminRequestCreate) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newAdmin := &model.Admin{
		ID:        primitive.NewObjectID(),
		Username:  admin.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = controller.CreateAdmin(newAdmin)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func LoginAdmin(request responses.AdminRequestCreate) (interface{}, error, int) {
	admin, err := controller.GetAdminByUsername(request.Username)
	if err != nil {
		return nil, err, 500
	}

	if admin == nil {
		return nil, errors.New("Forbidden"), 403
	}

	structureAdmin, ok := admin.(model.Admin)
	if !ok {
		return nil, errors.New("Cannot convert structure"), 500
	}

	plaintextPassword := []byte(request.Password)
	hashedPassword := []byte(structureAdmin.Password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, plaintextPassword)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Forbidden"), 403
	}

	SECRET_KEY := os.Getenv("SECRET_KEY")
	jwt, err := utils.SignJwt([]byte(SECRET_KEY), structureAdmin.Username)
	return jwt, nil, 0
}
