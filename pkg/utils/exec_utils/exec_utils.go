package exec_utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sikalabs/gobble/pkg/libtask"
	"github.com/sikalabs/gobble/pkg/utils/random_utils"
)

func RawExec(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	err := c.Run()
	return err
}

func RawExecStdout(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	return err
}

func Exec(taskInput libtask.TaskInput, cmd string, args ...string) error {
	if taskInput.Dry {
		dryArgs := []string{}
		for _, arg := range args {
			if strings.Contains(arg, " ") {
				dryArgs = append(dryArgs, "'"+arg+"'")
			} else {
				dryArgs = append(dryArgs, arg)
			}
		}
		fmt.Println(strings.Join(append([]string{cmd}, dryArgs...), " "))
		return nil
	}
	err := RawExec(cmd, args...)
	return err
}

func SSH(taskInput libtask.TaskInput, cmdArray ...string) error {
	args := append([]string{taskInput.SSHTarget}, cmdArray...)
	if taskInput.Sudo {
		args = append([]string{taskInput.SSHTarget, "sudo"}, cmdArray...)
	}
	return Exec(taskInput, "ssh", args...)
}

func rawSCP(taskInput libtask.TaskInput, localPath string, remotePath string) error {
	return Exec(
		taskInput,
		"scp", localPath,
		taskInput.SSHTarget+":"+remotePath,
	)
}

func SCP(taskInput libtask.TaskInput, localSrc string, remoteDst string) error {
	var err error
	tmpPath := "/tmp/" + random_utils.RandomString(32)
	err = rawSCP(taskInput, localSrc, tmpPath)
	if err != nil {
		return err
	}
	err = SSH(
		taskInput,
		"mv", tmpPath, remoteDst,
	)
	return err
}
