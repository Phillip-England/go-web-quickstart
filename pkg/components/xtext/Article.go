package xtext

import "fmt"

func Article(title string, details string) string {
	return fmt.Sprintf(`
		<div class='p-6 bg-white m-2 rounded'>
			<h2 class='text-lg mb-2 font-semibold'>%s</h2>
			<p class='text-sm'>%s</p>
		</div>
	`, title, details)
}