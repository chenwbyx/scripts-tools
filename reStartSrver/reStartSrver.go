/**
 * reStartSrver.go：监控某个文件夹下的文件是否有变化，当配置文件发生变化的时候，发送SIGHUP信号给指定进（这里测试使用的是server.go）
 * 设置CD每60秒内只能发送一次SIGHUP
 * Create By ChenWenBo
 * 2019-10-15
 */
package main

import (
	"github.com/howeyc/fsnotify"
	"log"
	"os/exec"
	"strings"
	"strconv"
	"syscall"
	"time"
)

func main() {
	processName := "server"  //需要重启的进程
	filePath := "."          //监控的配置文件的位置
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	lastTime := 0
	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				if int(time.Now().Unix()) - lastTime > 60 {
					lastTime = int(time.Now().Unix())
					//查找进程
					pids := findProcessExist(processName)
					if len(pids) > 0 {
						for _, v := range pids {
							syscall.Kill(v, syscall.SIGHUP)
						}
					}
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(filePath)
	if err != nil {
		log.Fatal(err)
	}

	<-done

	watcher.Close()
}

// 查找进程
func findProcessExist(appName string) []int {
	var pids []int
	// ps -ef | grep main | grep -v 'grep ' | awk '{print $2}'
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
