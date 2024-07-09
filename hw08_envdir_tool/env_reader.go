package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	st, err := os.Stat(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to get stat %s with error %s", dir, err.Error())
	}

	if !st.IsDir() {
		return nil, errors.New("expected dir")
	}

	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir %s with error %s", dir, err.Error())
	}

	env := make(Environment)
	for _, dirEntry := range dirEntries {
		filepath := dir + "/" + dirEntry.Name()
		envVar := strings.ReplaceAll(dirEntry.Name(), "=", "")

		file, err := os.Open(filepath)
		if err != nil {
			return nil, fmt.Errorf("failed to open file %s with err %s", filepath, err.Error())
		}
		defer file.Close()

		env[envVar] = EnvValue{
			NeedRemove: true,
		}

		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			content := strings.ReplaceAll(scanner.Text(), "\x00", "\n")
			content = strings.TrimRight(content, " \t")
			env[envVar] = EnvValue{
				Value:      content,
				NeedRemove: false,
			}
		}
	}

	return env, nil
}
