package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const doc = `tchat is a terminal chat cli.

Usage:
	tchat
	tchat [options] <word>...
	tchat -h | --help
	tchat -v | --version

Options:
	-h --help               Show this screen.
	-v --version            Show version.
	-r --right              Say word on right.
	-i --icon=<textfile>    Icon AA file.
	-w --width=<width>      Display width.`

func main() {
	opts, _ := docopt.ParseDoc(doc)
	if ok := opts["--version"].(bool); ok {
		fmt.Println(Version)
		return
	}

	// var err error
	// switch cmd := opts["<command>"]; cmd {
	// case nil:
	// 	err = CmdSearch()
	// case "checkout":
	// 	err = CmdCheckout()
	// case "log":
	// 	err = CmdLog()
	// case "new-feature-branch":
	// 	err = CmdNewFeatureBranch()
	// case "new-hotfix-branch":
	// 	err = CmdNewHotfixBranch()
	// case "push-set-upstream":
	// 	err = CmdPushSetUpstream()
	// case "set-ssh-url":
	// 	err = CmdSetSshUrl()
	// case "table":
	// 	err = CmdTable()
	// default:
	// 	fmt.Fprintln(os.Stderr, doc)
	// 	msg := fmt.Sprintf("Illegal command. command = %v", cmd)
	// 	err = errors.New(msg)
	// }
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }

}
