package main

import (
	"github.com/fsnotify"
	"log"
	"os/exec"
	"fmt"
	"strings"
	"strconv"
	"syscall"
)

func main() {
	processName := "main"
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				log.Println("notify runner to do the ln -s and restart server.")
				//查找进程
				pids := findProcessExist(processName)
				if len(pids) > 0 {
					for _,v := range pids {
						syscall.Kill(v, syscall.SIGTERM)
					}
				}
				execProcess(processName)

			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(".")
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}

// 查找进程
func findProcessExist(appName string) []int {
	var pids []int
	cmd := exec.Command("bash", "-c", "ps -ef | grep " + appName + " | grep -v 'grep ' | awk '{print $2}'")
	output, _ := cmd.Output()
	fields := strings.Fields(string(output))

	for _, v := range fields {
		pid,err := strconv.Atoi(v)
		if err == nil {
			pids = append(pids, pid)
		}
	}

	return pids
}

// 启动进程
func execProcess(appName string) {

	path := "./"    // app路径

	cmd := exec.Command(path + appName)
	cmd.Output()
	fmt.Println(appName, "进程启动")
}
