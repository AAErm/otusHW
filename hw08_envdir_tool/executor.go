package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 0
	}

	for k, v := range env {
		if v.NeedRemove {
			if err := os.Unsetenv(k); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if err := os.Setenv(k, v.Value); err != nil {
			log.Fatal(err)
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = os.Stdout // #nosec
	command.Stdin = os.Stdin
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return command.ProcessState.ExitCode()
	}

	return command.ProcessState.ExitCode()
}
