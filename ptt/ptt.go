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

	app.Commands = []cli.Command {
		{
			Name: "init",
			Aliases: []string{"i"},
			Usage: "create/add a git submodule for tracking file",
			Category: "git",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "repo,r",
					Value: "",
					Usage: "give a git repo url to track files in it",
				},
			},
			Action: func(c *cli.Context) error {
				var cmd *exec.Cmd
				repo := c.String("repo")
				if repo == ""{
					log.Println("a git repo url must be attached to initialize")
					return nil
				}
				cmd = exec.Command("/bin/bash","-c","git submodule add "+repo)
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
			Name: "track",
			Aliases: []string{"t"},
			Usage: "commit submodule changes of files to track",
			Category: "git",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "dir,d",
					Value: "",
					Usage: "give the directory of submodule to commit files to track",
				},
			},
			Action: func(c *cli.Context) error {
				var cmd *exec.Cmd
				dir := c.String("dir")
				fmt.Println(dir)
				if dir == ""{
					log.Println("a submodule path must be attached to commit")
					return nil
				}
				cmd = exec.Command("/bin/bash","-c","pwd&&cd "+dir+"&&pwd&&git add out.json&& git commit -m \"track performance\"&&git push")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
        	return err
    		}
				return nil
			},
		},
	}

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}