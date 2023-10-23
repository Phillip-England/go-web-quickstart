package xlayout

import "fmt"

func TeamNavMenu(currentPath string) string {
	talentLink := SimpleNavLink(currentPath, "Talent Engagement", "/team/talent")
	cemLink := SimpleNavLink(currentPath, "Customer Service", "/team/cem")
	salesLink := SimpleNavLink(currentPath, "Sales & Brand Growth", "/team/sales")
	financeLink := SimpleNavLink(currentPath, "Financial Stewardship", "/team/finance")
	logoutLink := SimpleNavLink(currentPath, "Logout", "/")

	return fmt.Sprintf(`
		<nav class='bg-white border-r w-3/5 fixed left-0 h-screen z-30 p-1'>
			<ul>
				%s%s%s%s%s
			</ul>
		</nav>
	`, talentLink, cemLink, salesLink, financeLink, logoutLink)
}