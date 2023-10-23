package actionRoutes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func ActionLoginUser(r *gin.Engine) {
	r.POST("/action/user/login", func(c *gin.Context) {
		fmt.Println("hit")
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == os.Getenv("SOUTHROADS_USERNAME") && password == os.Getenv("SOUTHROADS_PASSWORD") {
			c.Redirect(303, "/team/talent")
			return
		}
		c.Redirect(303, "/?LoginFormErr=invalid credentails")
	})
}