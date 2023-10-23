package xlayout

import "fmt"

func SimpleNavLink(currentPath string, text string, href string) string {
	var activeClasses string
	if currentPath == href {
		activeClasses = "underline"
	}
	return fmt.Sprintf(`
		<li class='%s flex rounded'>
			<a class='p-6 w-full' href=%s>%s</a>
		</li>
	`, activeClasses, href, text)
}