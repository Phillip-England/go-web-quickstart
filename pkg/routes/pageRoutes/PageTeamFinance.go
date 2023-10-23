package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamFinance(r *gin.Engine) {
	r.GET("/team/finance", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Financial Stewardship")
        b.Add(xlayout.Banner("CFA Suite", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Financial Stewardship", "We can bring in sales, but that does not make us profitable. We are only profitable is our sales are greater than our expenses. The results in this section indicate how profitable we are and give a good idea of where we are spending money."))
		b.Add(`
			<div hx-get='/components/FinancialResultsWidget' hx-swap='outerHTML' hx-trigger='load' class='bg-white m-2 rounded p-6 flex flex-col gap-4'>
				<div class='bg-gray-200 w-4/5 h-6 mb-2 animate-pulse rounded'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			</div>
		`)
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}