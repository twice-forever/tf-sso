package api

import (
	"net/http"
	"tf-sso/model"
	"tf-sso/service/common"
	"tf-sso/service/session"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	login := new(common.Login)
	if err := c.Bind(login); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}

	if err := c.Validate(login); err != nil {
		return err
	}

	user := model.User{Username: login.Username}
	if err := user.Get(); err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  "登录失败,请重试!",
		})
	}

	hashPassword := common.GeneratePassHash(login.Password, user.PasswordSalt)
	if hashPassword != user.Password {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  "登录失败,请重试!",
		})
	}

	session, _ := session.Store.Get(c.Request(), "user-session")
	session.Options.HttpOnly = true
	session.Values["id"] = user.ID
	session.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Msg:  "登录成功!",
	})
}

func Logout(c echo.Context) error {
	session, _ := session.Store.Get(c.Request(), "user-session")
	session.Options.MaxAge = -1
	session.Save(c.Request(), c.Response())
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Msg:  "登出成功!",
	})
}
