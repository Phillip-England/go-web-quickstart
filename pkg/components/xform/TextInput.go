package xform

import "fmt"

func TextInput(name string, serverName string) string {
	return fmt.Sprintf(`
		<label class='text-xs' for=%s>%s</label>
		<input name=%s class='text-xs rounded p-1 border' type='text' />
	`, serverName, name, serverName)
}