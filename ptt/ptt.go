package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// "io/ioutil"
	"os/exec"

	"github.com/urfave/cli"
)

type Diff struct {
	Diff string `json:"diff"`
}

// type Commits struct {
// 	Commit1 string `json:"commit1"`
// 	Commit2 string `json:"commit2"`
// }

func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func diff_commits(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	// fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	var c1 string
	var c2 string
	for k, v := range r.Form {
		fmt.Println(k)
		if k == "commit1" {
			c1 = strings.Join(v, "")
			fmt.Println(c1)
		}
		if k == "commit2" {
			c2 = strings.Join(v, "")
			fmt.Println(c2)
		}
	}

	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash", "-c", "git diff -U1 "+c1+" "+c2)

	out, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
		return
	}

	d := string(out)

	res := Diff{d}

	js, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

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
				" directory in current git project, and ignore it in .gitignore file. It will " +
				" hold some cached files for diff useage.",
			Category: "git",
			Action: func(c *cli.Context) error {
				var cmd *exec.Cmd
				cmd = exec.Command("/bin/bash", "-c", "echo '/.ptt'>>.gitignore"+
					"mkdir .ptt &&"+
					"cd .ptt &&"+
					"mkdir c &&"+
					"mkdir ca &&"+
					"mkdir cb &&"+
					"cd ..")
				out, err := cmd.Output()
				if err != nil {
					fmt.Print(out)
					fmt.Print(err)
					return err
				}
				fmt.Print(out)
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
					"git push origin "+cb+":gh-review")
				out, err = git.Output()
				fmt.Print("pushing code to gh-review ...\n")
				if err != nil {
					fmt.Print(string(out))
					fmt.Print(err)
					return err
				}
				fmt.Print(string(out) + "Completed\n")
				return nil
			},
		},
		{
			Name:     "server",
			Aliases:  []string{"s"},
			Usage:    "launch a server for git information",
			Category: "git",
			Action: func(c *cli.Context) {
				http.HandleFunc("/", hello_world)
				http.HandleFunc("/apis/diff_text", diff_commits) //设置访问的路由
				err := http.ListenAndServe(":9090", nil)         //设置监听的端口
				if err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
