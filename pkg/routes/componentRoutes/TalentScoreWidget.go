package componentRoutes

import (
	"fmt"
	"net/http"
	"web-quickstart/pkg/components/xlist"

	"github.com/gin-gonic/gin"
)

func TalentScoresWidget(r *gin.Engine) {
	r.GET("/components/TalentScoresWidget", func(c *gin.Context) {
		teamCount := xlist.ToggleScoreLi("Team Count", "199", "We aim to keep our team staffed to a healthy number throughout the entire year. We are always searching for top-talet to partner with us and have a positve influence on our guests and team.")
		turnoverCount := xlist.ToggleScoreLi("Retention (12 month)", "88", "Retention is a measure of how well we are keeping people over a given peiod of time. A higher retention could be an indication that we are creating a healthy culture where people want to stay long-term.")
		retentionCount := xlist.ToggleScoreLi("Turnover (12 month)", "65", "Turnover is a meausre of how many people are leaving the business over a given period over time. A higher turnover could indicate an issue with our culture, onboarding, or training process.")
		widget := fmt.Sprintf(`
			<div class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
				<div class='text-xs items-center justify-between gap-2 flex'>
					<h2 class='font-semibold text-lg'>Talent Results</h2>
					<div class='flex flex-col'>
						<select class='border p-1'>
							<option>Southroads</option>
							<option>Utica</option>
						</select>
					</div>
				</div>
				<ul class='flex flex-col rounded gap-2 text-xs'>
					%s
					%s
					%s
				</ul>
			</div>
		`, teamCount, turnoverCount, retentionCount)
        c.Data(http.StatusOK, "text/html", []byte(widget))
    })
}