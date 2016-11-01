package cli

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v2"
	"time"
)

// argsParam parse one by one the given argumentes
func argsParam(params *cli.Context) []string {
	argsN := params.NArg()
	if argsN > 0 {
		var args []string
		for i := 0; i <= argsN-1; i++ {
			args = append(args, params.Args().Get(i))
		}
		return args
	}
	return nil
}

// BoolParam is used to check the presence of a bool flag
func boolFlag(b bool) bool {
	if b {
		return false
	}
	return true
}

// Duplicates check projects with same name or same combinations of main/path
func duplicates(value Project, arr []Project) (Project, error) {
	for _, val := range arr {
		if value.Path == val.Path || value.Name == val.Name {
			return val, errors.New("There is a duplicate of '" + val.Name + "'. Check your config file!")
		}
	}
	return Project{}, nil
}

// check if a string is inArray
func inArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// Cewrites the log timestamp
func (w logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(w.Yellow.Regular("[") + time.Now().Format("15:04:05") + w.Yellow.Regular("]") + string(bytes))
}
