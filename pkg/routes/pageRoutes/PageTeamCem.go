package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamCem(r *gin.Engine) {
	r.GET("/team/cem", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Customer Service")
        b.Add(xlayout.Banner("CFA Suite", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Customer Service", "Our customer service is what sets apart from our competitors. Chick-fil-A is known for being a company that cares about people, not just their money. The results in this section help give us an idea about how our guest feel about our business."))
		b.Add(`
			<div hx-get='/components/CustomerServiceScoreWidget' hx-swap='outerHTML' hx-trigger='load' class='bg-white m-2 rounded p-6 flex flex-col gap-4'>
				<div class='bg-gray-200 w-4/5 h-6 mb-2 animate-pulse rounded'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
			</div>
		`)
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}