package xform

import "fmt"

func BaseForm(title string, action string, formErr string, inputs string) string {
	var loginFormErr string
	if formErr == "" {
		loginFormErr = ""
	} else {
		loginFormErr = fmt.Sprintf("<p class='text-xs text-red-500'>%s</p>", formErr)
	}
	return fmt.Sprintf(`
		<form action=%s method='POST' class='flex bg-white flex-col border rounded m-2 p-4 gap-4'>
			<div class='flex flex-col gap-2'>
				<h2>%s</h2>
				%s
			</div>
			<div class='flex flex-col gap-2'>
				%s
			</div>
			<div class='flex'>
				<input _='on click toggle .hidden on me then toggle .hidden on the next <div/>' type='submit' class='bg-black rounded text-xs text-white px-6 py-2' />
				<div class='h-8 w-8 hidden rounded-full border border-t-black animate-spin'></div>
			</div>
		</form>
	`, action, title, loginFormErr, inputs)
}