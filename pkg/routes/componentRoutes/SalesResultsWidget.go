package componentRoutes

import (
	"fmt"
	"net/http"
	"web-quickstart/pkg/components/xlist"

	"github.com/gin-gonic/gin"
)

func SalesResultsWidget(r *gin.Engine) {
	r.GET("/components/SalesResultsWidget", func(c *gin.Context) {
		totalSales := xlist.ToggleScoreLi("Total Sales", "$209,399", "Sales is simply money coming into the business through the guest transactions. Total sales includes both inside and outside sales.")
		insideSales := xlist.ToggleScoreLi("Inside Sales", "$134,555", "Inside sales is money which is derived from any transaction which is not catering delivery. If the order is not a catering delivery order, then the money for the order goes to inside sales.")
		outsideSales := xlist.ToggleScoreLi("Outside Sales", "$23,339", "Outside sales is money which is derived from catering delivery orders. If an order is a catering delivery order, then the money for the order goes to outside sales.")
		widget := fmt.Sprintf(`
			<div class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
				<div class='text-xs items-center justify-between gap-2 flex'>
					<h2 class='font-semibold self-start text-lg'>Sales Results</h2>
					<div class='flex flex-col'>
						<div class='flex flex-col gap-2'>
							<select class='border p-1'>
								<option>Southroads</option>
								<option>Utica</option>
							</select>
							<select class='border p-1'>
								<option>Current Month</option>
								<option>Previous Month</option>
								<option>Year To Date</option>
							</select>
						</div>
					</div>
				</div>
				<ul class='flex flex-col rounded gap-2 text-xs'>%s%s%s</ul>
			</div>
		`, totalSales, insideSales, outsideSales)
        c.Data(http.StatusOK, "text/html", []byte(widget))
    })
}