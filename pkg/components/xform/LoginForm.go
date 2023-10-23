package xform

import "fmt"

func LoginForm(formErr string) string {
	inputs := fmt.Sprintf(`%s%s`,
		TextInput("username", "username"),
		TextInput("password", "password"),
	)
	return BaseForm("Login", "/action/user/login", formErr, inputs)
}