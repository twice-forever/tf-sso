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
		Code: http.StatusOK,
		Msg:  "添加成功",
	})
}

func DeleteUser(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	err = user.Delete()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Msg:  "删除成功",
	})
}

func GetUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
	}
	if err := user.Get(); err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}
	if user.Password != "" {
		user.Password = "***"
	}
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Data: user,
	})
}

func GetUsers(c echo.Context) error {
	users, err := model.GetUsers(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}
	for i, v := range *users {
		if v.Password != "" {
			(*users)[i].Password = "***"
		}
	}
	count, err := new(model.User).Count()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	}
	return c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Data: common.PageResponse{
			Total: int(count),
			Data:  users,
		},
	})
}
