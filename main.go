package main

import (
	"fmt"
	"templateLoginRegister/configs"
	ac "templateLoginRegister/delivery/controllers/auth"
	uc "templateLoginRegister/delivery/controllers/user"
	"templateLoginRegister/delivery/routes"
	authRepo "templateLoginRegister/repository/auth"
	userRepo "templateLoginRegister/repository/user"
	"templateLoginRegister/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	config := configs.GetConfig()

	db, err := utils.InitDB(config)
	if err != nil {
		panic("error database")
	}
	defer db.Close()

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	routes.RegisterPath(e, authController, userController)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
