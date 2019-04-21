package main

import (
  "fmt"
  "log"
  "os"
  "io/ioutil"
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
			Name: "commit",
			Aliases: []string{"c"},
			Usage: "create a git commit with performance",
			Category: "git actions",
			Action: func(c *cli.Context) error {
				cmd := exec.Command("ls", "-a", "-l")
				stdout, err := cmd.StdoutPipe()
				if err != nil {
					log.Fatal(err)
				}
				defer stdout.Close()
				if err := cmd.Start(); err != nil {
					log.Fatal(err)
				}
				opBytes, err := ioutil.ReadAll(stdout)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(string(opBytes))
				return nil
			},
		},
	}

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}