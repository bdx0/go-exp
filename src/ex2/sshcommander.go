
// http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

type SSHCommander struct {
	User string
	IP   string
}

func (s *SSHCommander) Command(cmd ...string) *exec.Cmd {
	arg := append(
		[]string{
			fmt.Sprintf("%s@%s", s.User, s.IP),
		},
		cmd...,
	)
	return exec.Command("ssh", arg...)
}

func main() {
	commander := SSHCommander{"root", "50.112.213.24"}

	cmd := []string{
		"apt-get",
		"install",
		"-y",
		"jq",
		"golang-go",
		"nginx",
	}

	// am I doing this automation thing right?
	if err := commander.Command(cmd...); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running SSH command: ", err)
		os.Exit(1)
	}
}