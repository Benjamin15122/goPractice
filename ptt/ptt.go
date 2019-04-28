package main

import (
	"fmt"
	"log"
	"os"

	// "io/ioutil"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pet"
	app.Usage = "performance evaluation track tool"
	app.Action = func(c *cli.Context) error {
		fmt.Println("This is a tool for performance tracking, try pet -h to get some help with commands")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage: "This command create a .ptt" +
				" directory in current git project, and ignore it in .gitignore file. It will download a built web project " +
				" to .ptt/ and deploy it to gh-page branch.",
			Category: "git",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "repo,r",
					Value: "",
					Usage: "give a git repo url to track files in it",
				},
			},
			Action: func(c *cli.Context) error {
				var cmd *exec.Cmd
				repo := c.String("repo")
				if repo == "" {
					log.Println("a git repo url must be attached to initialize")
					return nil
				}
				cmd = exec.Command("/bin/bash", "-c", "git submodule add "+repo)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:     "commit",
			Aliases:  []string{"t"},
			Usage:    "commit all files to a gh-ptt branch",
			Category: "git",
			Action: func(c *cli.Context) error {
				var cmd *exec.Cmd
				cmd = exec.Command("/bin/bash", "-c", "git symbolic-ref --short -q HEAD")
				// cmd.Stderr = os.Stderr
				// cmd.Stdout = os.Stdout
				// err := cmd.Run()
				out, err := cmd.Output()
				if err != nil {
					fmt.Print(err)
					return err
				}
				cb := string(out)
				fmt.Print("current branch: " + cb)
				git := exec.Command("/bin/bash", "-c", "git add .&&"+
					"git commit -m \"gh-review update\"&&"+
					"git push "+cb+":gh-review")
				out, err = git.Output()
				if err != nil {
					fmt.Print(err)
					return err
				}
				fmt.Print("pushing code to gh-review ...\n" + string(out) + "Completed\n")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
