package cp

import (
	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/exec_utils"
)

type TaskCp struct {
	LocalSrc  string `yaml:"local_src"`
	RemoteDst string `yaml:"remote_dst"`
}

func Run(
	taskInput libtask.TaskInput,
	taskParams TaskCp,
) libtask.TaskOutput {
	err := exec_utils.SCP(taskInput, taskParams.LocalSrc, taskParams.RemoteDst)
	return libtask.TaskOutput{
		Error: err,
	}
}
