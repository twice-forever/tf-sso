package api

import (
	"net/http"
	"tf-sso/model"
	"tf-sso/service/common"
	"tf-sso/util/log"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		log.Logger.Error().Msg(err.Error())
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	user.Owner = "build-in"
	user.ID = uuid.NewString()
	user.PasswordSalt = common.GenerateSalt()
	user.Password = common.GeneratePassHash(user.Password, user.PasswordSalt)

	if err := user.Add(); err != nil {
		log.Logger.Error().Msg(err.Error())
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusInternalServerError,
		Msg:  "添加成功",
	})
}
