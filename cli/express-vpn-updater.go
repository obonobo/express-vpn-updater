package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"os/user"

	"github.com/obonobo/express-vpn-updater/cli/argparse"
	"github.com/obonobo/express-vpn-updater/cli/client"
)

var (
	args argparse.Arguments = argparse.Parse()
	c    *client.Client     = client.New(args.Url)
)

func main() {
	panicIfNotRoot()
	resp, err := c.DownloadLatest()
	if err != nil {
		panic(err)
	}
	log.Println(aptUpdate())
	log.Println(aptInstall(resp))
}

func panicIfNotRoot() {
	if !amRoot() {
		panic(errors.New("this tool must be run as root"))
	}
}

func aptInstall(path string) string {
	cmd := exec.Command("apt", "install", "-y", path)
	var out bytes.Buffer
	cmd.Stderr = &out
	cmd.Stdout = &out
	if err := cmd.Run().(*exec.ExitError); err != nil {
		if err.ExitCode() != 100 {
			panic(err)
		}
	}
	return out.String()
}

func aptUpdate() string {
	cmd := exec.Command("apt", "update")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return out.String()
}

func amRoot() bool {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln("Error encountered when checking user - this tool must be run as root")
	}
	return usr.Uid == "0"
}
