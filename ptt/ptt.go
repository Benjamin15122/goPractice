package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	// "io/ioutil"
	"os/exec"

	"github.com/urfave/cli"
)

type Err struct {
	Error string `json: error`
}

type Out struct {
	Files []string `json: files`
}

type Diff struct {
	Diff string `json:"diff"`
}

// 测试文件或目录是否存在，有bug
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && err.Error() == os.ErrNotExist.Error() {
		return false
	}
	return true
}

// 对http请求返回helloworld
func hello_world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// 请求某个commit的输出内容，返回符合要求的文件名列表
func commit_output(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var sha string
	for k, v := range r.Form {
		fmt.Println(k)
		if k == "sha" {
			sha = strings.Join(v, "")
			break
		}
	}

	//获取当前程序运行目录
	// dir, d_err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	// if d_err != nil {
	// 	log.Println(d_err)
	// }

	//获取当前git 分支
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/bash", "-c", "git symbolic-ref --short -q HEAD")

	out, err := cmd.Output()
	if err != nil {
		js, _ := json.Marshal(Err{"fetch git HEAD ref failed, check console output for more information.\n" + err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	cb := string(out)

	//	尝试checkout到指定历史版本，失败则返回错误信息
	cmd = exec.Command("/bin/bash", "-c", "git checkout "+sha)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		js, _ := json.Marshal(Err{"git checkout failed, maybe have unstaged files, try commit first, check backend console output for more information. \n" + err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}

	// 将指定版本文件存储到指定目录
	cmd = exec.Command("/bin/bash", "-c", "rm -rf .ptt/c/__out &&"+
		"cp -r __out .ptt/c/ ")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		js, _ := json.Marshal(Err{"save file failed," +
			" check console output for more information.\n" +
			err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	fmt.Println("save file succeeded, ready to serve")

	// 返回原分支HEAD
	cmd = exec.Command("/bin/bash", "-c", "git checkout "+cb)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		js, _ := json.Marshal(Err{"git checkout to HEAD failed," +
			" check console output for more information.\n" +
			err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}

	//在保存的文件夹中查找json文件
	j_array, j_err := filepath.Glob(".ptt/c/__out/*.json")
	// l_array, l_err := filepath.Glob(".ptt/c/__out/log.txt")

	//若查找失败返回则response写入空数组
	if j_err != nil {
		log.Println(j_err)
		js, _ := json.Marshal(Err{"json files not found," +
			" check backend console output for more information. \n" +
			j_err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}
	// if l_err != nil {
	// 	log.Println(l_err)
	// 	l_array = []string{}
	// 	fmt.Println("log status: not found")
	// }

	//将查找文件名组织成json写入response
	// for i := 0; i < len(f_array); i++ {
	// 	f_url := dir + "/" + f_array[i]
	// 	f_array[i] = f_url
	// }

	// for i := 0; i < len(l_array); i++ {
	// 	l_url := dir + "/" + l_array[i]
	// 	l_array[i] = l_url
	// }

	fmt.Println("found json files: ", j_array)
	// fmt.Println("found txt files: ", l_array)

	res := Out{j_array}
	js, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
				cmd = exec.Command("/bin/bash", "-c", "echo '/.ptt'>>.gitignore &&"+
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
				http.HandleFunc("/apis/commit_out", commit_output)
				http.HandleFunc("/apis/diff_text", diff_commits) //设置访问的路由
				err := http.ListenAndServe(":9090", nil)         //设置监听的端口
				if err != nil {
					log.Fatal("ListenAndServe: ", err)
				}
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test some env",
			Action: func(c *cli.Context) {
				dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
				if err != nil {
					log.Println(err)
				}
				fmt.Println(dir)
			},
		},
		{
			Name:  "stage",
			Usage: "stage git change",
			Action: func(c *cli.Context) {
				var cmd *exec.Cmd
				cmd = exec.Command("/bin/bash", "-c", "git add .&&git commit -m \"update\"&&git push")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
