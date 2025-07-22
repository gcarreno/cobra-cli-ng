package projects

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type GoList struct {
	Path string `json:"Path"`
	Dir  string `json:"Dir"`
}

func NewGoList() (*GoList, error) {
	out, err := exec.Command("go", "list", "-json", "-m").Output()
	if err != nil {
		return nil, err
	}

	goList := &GoList{}

	if err := json.NewDecoder(strings.NewReader(string(out))).Decode(goList); err != nil {
		return nil, err
	}

	if goList.Dir == "" {
		return nil, fmt.Errorf(`please run "go mod init <MODNAME>" first`)
	}

	return goList, nil
}
