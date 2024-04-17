package main

import (
	"errors"
	"os/exec"
	"runtime"

	hook "github.com/robotn/gohook"
)

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func main() {
	for {
		mleft := hook.AddEvent("mleft")
		if mleft {
			go Open("https://google.com")
		}
	}
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return errors.New("dont know how to open things platfrom system: " + runtime.GOOS)
	}
	cmd := exec.Command(run, uri)
	return cmd.Start()
}
