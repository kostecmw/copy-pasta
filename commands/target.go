package commands

import (
	"fmt"

	"github.com/jutkko/copy-pasta/runcommands"
	"github.com/mitchellh/cli"
)

type TargetCommand struct {
	Ui cli.Ui
}

func (t *TargetCommand) Help() string {
	return "Changes the current target to the provided target"
}

func (t *TargetCommand) Run(args []string) int {
	if len(args) > 0 {
		config, err := loadRunCommands()
		if err != nil {
			return 1
		}

		if target, ok := config.Targets[args[0]]; ok {
			if err := runcommands.Update(target.Name, target.AccessKey, target.SecretAccessKey, target.BucketName); err != nil {
				t.Ui.Error(fmt.Sprintf("Failed to update the current target: %s", err.Error()))
				return 2
			} else {
				return 0
			}
		} else {
			t.Ui.Error("Target is invalid")
			return 3
		}
	} else {
		t.Ui.Error("No target provided")
		return 4
	}
}

func (t *TargetCommand) Synopsis() string {
	return "Changes the current target to the provided target"
}
