package pageRoutes

import (
	"net/http"
	"web-quickstart/pkg/components/xlayout"
	"web-quickstart/pkg/components/xtext"
	"web-quickstart/pkg/core"

	"github.com/gin-gonic/gin"
)


func PageTeamTalent(r *gin.Engine) {
	r.GET("/team/talent", func(c *gin.Context) {
        b := core.NewPageBuilder("CFA Suite | Talent Engagement")
        b.Add(xlayout.Banner("CFA Suite", xlayout.TeamNavMenu(c.Request.URL.Path)))
		b.Add(xtext.Article("Talent Engagement", "All of the metrics in this section revolve around the people in our business. These results are important because we believe people are a critical success factor."))
		b.Add(`
			<div hx-get='/components/TalentScoresWidget' hx-swap='outerHTML' hx-trigger='load' class='bg-white m-2 rounded p-6 flex flex-col gap-4'>
				<div class='bg-gray-200 w-4/5 h-6 mb-2 animate-pulse rounded'></div>
				<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
				<div class='bg-gray-200 w-2/3 pr-2 h-4 animate-pulse rounded self-end'></div>
			</div>
		`)
        c.Data(http.StatusOK, "text/html", b.GetHTMLBytes())
    })
}