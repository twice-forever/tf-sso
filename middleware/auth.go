package middlerware

import (
	"net/http"
	"tf-sso/service/common"
	"tf-sso/service/session"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := session.Store.Get(c.Request(), "user-session")
		if _, ok := session.Values["id"]; !ok {
			return c.JSON(http.StatusInternalServerError, common.Response{
				Code: http.StatusInternalServerError,
				Msg:  "用户未登录",
			})
		}
		return next(c)
	}
}
