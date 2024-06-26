package task

import (
	"fmt"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/task/lib/apt_install"
	"github.com/sikalabs/gobble/pkg/task/lib/authorized_key"
	"github.com/sikalabs/gobble/pkg/task/lib/chmod"
	"github.com/sikalabs/gobble/pkg/task/lib/command"
	"github.com/sikalabs/gobble/pkg/task/lib/cp"
	"github.com/sikalabs/gobble/pkg/task/lib/echo"
	"github.com/sikalabs/gobble/pkg/task/lib/print"
	rsilf "github.com/sikalabs/gobble/pkg/task/lib/replace_string_in_local_file"
	"github.com/sikalabs/gobble/pkg/task/lib/template"
)

type Task struct {
	Name                         string                             `yaml:"name"`
	Echo                         echo.TaskEcho                      `yaml:"echo"`
	AptInstall                   apt_install.TaskAptInstall         `yaml:"apt_install"`
	Cp                           cp.TaskCp                          `yaml:"cp"`
	Template                     template.TaskTemplate              `yaml:"template"`
	Command                      command.TaskCommand                `yaml:"command"`
	Chmod                        chmod.TaskChmod                    `yaml:"chmod"`
	Print                        print.TaskPrint                    `yaml:"print"`
	AuthorizedKey                authorized_key.TaskAuthorizedKey   `yaml:"authorized_key"`
	TaskReplaceStringInLocalFile rsilf.TaskReplaceStringInLocalFile `yaml:"replace_string_in_local_file"`
}

func Run(
	taskInput libtask.TaskInput,
	task Task,
) libtask.TaskOutput {
	switch {
	case task.AptInstall.Name != "":
		return apt_install.Run(taskInput, task.AptInstall)
	case task.Echo.Message != "":
		return echo.Run(taskInput, task.Echo)
	case task.Cp.LocalSrc != "" || task.Cp.RemoteSrc != "":
		return cp.Run(taskInput, task.Cp)
	case task.Template.Path != "":
		return template.Run(taskInput, task.Template)
	case task.Command.Cmd != "":
		return command.Run(taskInput, task.Command)
	case task.Chmod.Path != "":
		return chmod.Run(taskInput, task.Chmod)
	case task.Print.Template != "":
		return print.Run(taskInput, task.Print)
	case task.AuthorizedKey.Key != "":
		return authorized_key.Run(taskInput, task.AuthorizedKey)
	case task.TaskReplaceStringInLocalFile.Path != "":
		return rsilf.Run(taskInput, task.TaskReplaceStringInLocalFile)
	}
	return libtask.TaskOutput{
		Error: fmt.Errorf("task \"%s\" not found", task.Name),
	}
}
