package componentRoutes

import (
	"fmt"
	"net/http"
	"web-quickstart/pkg/components/xlist"

	"github.com/gin-gonic/gin"
)

func CutomerServiceScoreWidget(r *gin.Engine) {
	r.GET("/components/CustomerServiceScoreWidget", func(c *gin.Context) {
		osatCount := xlist.ToggleScoreLi("OSAT", "70", "OSAT stands for 'overall satisfaction' and measures how satisfied a guest was with their experience overall. OSAT is the score we use to get a general feel for how our guests percieve our business.")
		tasteCount := xlist.ToggleScoreLi("Taste", "78", "Taste measures how much guests enjoyed their food. When a guest completes a survey, they will be asked questions about the specific items they ordered with their meal. The responses to those questions generate this score.")
		speedCount := xlist.ToggleScoreLi("Speed", "72", "Speed measures how fast guests percieve our business to be. Because speed is based on guest perception, speed can be influenced by a variety of factors. For example, guest engagement has the potential to change the way guests percieve time and make them feel like they waited less.")
		aceCount := xlist.ToggleScoreLi("ACE", "68", "ACE stands for, 'attentive and courteous employees' and measure show friendly guests percieve our team to be. When our guests visit Chick-fil-A, they have a certain standard they expect our team to meet. Chick-fil-A team members are held to a higher standard than other chains in our industry.")
		cleanlinessCount := xlist.ToggleScoreLi("Cleanliness", "77", "Cleanliness is an overall measure of how clean guests percieve our business to be. This score takes into account all areas of cleanliness including parking lot, dining room, drive-thru, and restroom cleanliness.")
		accuracyCount := xlist.ToggleScoreLi("Accuracy", "96", "Accuracy is the only score on this list which is derived from a yes or no question. Guests are simply asked if they recieved their order correctly, and then given the option to mark what was missed if they did not receieve what they expected.")
		widget := fmt.Sprintf(`
			<div class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
				<div class='text-xs items-center justify-between gap-2 flex'>
					<h2 class='font-semibold self-start text-lg'>CEM Results</h2>
					<div class='flex flex-col'>
						<div class='flex flex-col gap-2'>
							<select class='border p-1'>
								<option>Southroads</option>
								<option>Utica</option>
							</select>
							<select class='border p-1'>
								<option>Current Month</option>
								<option>Three Month Rolling</option>
								<option>Year To Date</option>
							</select>
						</div>
					</div>
				</div>
				<ul class='flex flex-col rounded gap-2 text-xs'>%s%s%s%s%s%s</ul>
			</div>
		`, osatCount, tasteCount, speedCount, aceCount, cleanlinessCount, accuracyCount)
        c.Data(http.StatusOK, "text/html", []byte(widget))
    })
}