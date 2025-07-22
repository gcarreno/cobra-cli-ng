package main

import (
    "{{ .Parent.Module }}/{{ if .Path }}{{ .Path}}/{{ end }}{{ .CommandsPath }}"
)

func main() {
	cmd.Execute()
}
