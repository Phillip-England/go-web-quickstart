package componentRoutes

import (
	"fmt"
	"net/http"
	"web-quickstart/pkg/components/xlist"

	"github.com/gin-gonic/gin"
)

func FinancialResultsWidget(r *gin.Engine) {
	r.GET("/components/FinancialResultsWidget", func(c *gin.Context) {
		laborCost := xlist.ToggleScoreLi("Labor Cost", "21.3%", "Labor represents all the money we spent to pay for team members to work in our business. Labor is our 2nd biggest expense and is an area of the business which has a lot of potetial to leak revenue.")
		foodCost := xlist.ToggleScoreLi("Food Cost", "30.1%", "Food cost represents all the money we spent on food to run the business. Food is our biggest expense and food waste is a big potential source of lost revenue.")
		cashShort := xlist.ToggleScoreLi("Cash Short", "$21.33", "Cash short represents all money lost in the business through the mishandling of cash by team members or business leadership. Cash short can be an indicator of theft or human error in regarding cash accountability.")
		profit := xlist.ToggleScoreLi("Net Profit", "12.3%", "Profit is simply what money is left on the table after all expenses are paid for. If you are given $100 and spend $80 on business expenses in a month, then the profit for that month is $20 or 20%.")
		widget := fmt.Sprintf(`
			<div class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
				<div class='text-xs items-center justify-between gap-2 flex'>
					<h2 class='font-semibold self-start text-lg'>Financial Results</h2>
					<div class='flex flex-col'>
						<div class='flex flex-col gap-2'>
							<select class='border p-1'>
								<option>Southroads</option>
								<option>Utica</option>
							</select>
							<select class='border p-1'>
								<option>Previous Month</option>
								<option>Year To Date</option>
							</select>
						</div>
					</div>
				</div>
				<ul class='flex flex-col rounded gap-2 text-xs'>%s%s%s%s</ul>
			</div>
		`, laborCost, foodCost, cashShort, profit)
        c.Data(http.StatusOK, "text/html", []byte(widget))
    })
}