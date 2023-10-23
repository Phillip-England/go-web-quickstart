package xlayout

import "fmt"

func Banner(title string, menu string) string {
	overlay := Overlay()
	return fmt.Sprintf(`
		<div class='flex flex-row border-b justify-between h-20 fixed top-0 z-40 w-full bg-white'>
			<h1 class='text-xl p-6 font-semibold'>%s</h1>
			<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group p-6'>
				<i class='fa-solid fa-bars fa-lg'></i>
			</div>
			<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group p-6 hidden'>
				<i class='fa-solid fa-x fa-lg'></i>
			</div>
		</div>
		<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group hidden'>%s</div>
		<div class='h-20'></div>
		<div class='top-nav-group hidden flex text-sm'>%s</div>
	`, title, overlay, menu)
}