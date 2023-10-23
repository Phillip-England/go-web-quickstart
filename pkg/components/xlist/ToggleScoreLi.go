package xlist

import "fmt"

func ToggleScoreLi(name string, score string, detail string) string {
	return fmt.Sprintf(`
	<li _='on click toggle .hidden on <p.toggle-score-detail/> in me then toggle .hidden on <div.toggle-score-caret/> in me' class='flex-col flex p-2 pb-4 border-b gap-4'>
		<div class='flex-col flex'>
			<div class='flex justify-between items-center'>
				<div class='flex flex-row gap-4 items-center'>
					<div class='toggle-score-caret'>
						<i class='fa-solid fa-caret-right'></i>
					</div>
					<div class='hidden toggle-score-caret'>
						<i class='fa-solid fa-caret-down'></i>
					</div>
					<h2 class=''>%s</h2>
				</div>
				<p class=''>%s</p>
			</div>
			<p class='toggle-score-detail pt-4 hidden'>%s</p>
		</div>
	</li>

	`, name, score, detail)
}