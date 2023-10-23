package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xform"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageLogin(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		loginFormErr := c.Query("LoginFormErr")
        b := core.NewPageBuilder("CFA Suite | Login")
        b.Add(xlayout.SimpleBanner("CFA Suite"))
		b.Add(xform.LoginForm(loginFormErr))
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}