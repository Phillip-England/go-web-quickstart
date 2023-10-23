package xlayout

import "fmt"

func SimpleBanner(title string) string {
	return fmt.Sprintf(`
		<div class='flex flex-row font-semibold justify-between bg-white border-b'>
			<h1 class='text-xl p-6'>%s</h1>
		</div>
	`, title)
}