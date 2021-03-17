package i3

import (
	"errors"
	"fmt"

	"github.com/lpicanco/i3-autodisplay/config"
	"go.i3wm.org/i3/v4"
)

func GetCurrentWorkspaceNumber() (int64, error) {
	ws, err := i3.GetWorkspaces()
	if err != nil {
		return -1, err
	}

	for _, w := range ws {
		if w.Focused {
			return w.Num, nil
		}
	}

	return -1, errors.New("Cant find current workspace")
}

func SetCurrentWorkspace(workspaceNum int64) error {
	command := fmt.Sprintf("workspace %d", workspaceNum)
	_, err := i3.RunCommand(command)
	return err
}

func UpdateWorkspace(display config.Display, workspace int) error {
	command := fmt.Sprintf("workspace %d; move workspace to %s", workspace, display.Name)
	fmt.Println(command)
	_, err := i3.RunCommand(command)

	if err != nil {
		return err
	}

	return nil
}
