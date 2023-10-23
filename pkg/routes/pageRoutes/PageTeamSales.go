package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamSales(r *gin.Engine) {
	r.GET("/team/sales", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Sales & Brand Growth")
        b.Add(xlayout.Banner("CFA Suite", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Sales & Brand Growth", "Our people are great, but we wouldn't have this opportunity without our guests. If they did not come and make purchases, we would be employeed elsewhere. The results in this section help us keep an eye on how much money our business is generating."))
		b.Add(`
			<div hx-get='/components/SalesResultsWidget' hx-swap='outerHTML' hx-trigger='load' class='bg-white m-2 rounded p-6 flex flex-col gap-4'>
				<div class='bg-gray-200 w-4/5 h-6 mb-2 animate-pulse rounded'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			</div>
		`)
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}