package xloading

import "fmt"

func HxGetIndicator(hxGet string, trigger string, containerClasses string) string {
	return fmt.Sprintf(`
		<div hx-get='%s' hx-swap='outerHTML' hx-trigger='%s' class='%s'>
			%s
		</div>
	`, hxGet, trigger, containerClasses, "")
}