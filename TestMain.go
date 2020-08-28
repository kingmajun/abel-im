package main


import (
	"log"
	"os/exec"
	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
				//nginx()
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("D:/avatar")
	if err != nil {
		log.Fatal(err)
	}
/*   err = watcher.Add("D:/avatar")//也可以监听文件夹
	if err != nil {
		log.Fatal(err)
	}*/
	<-done
}

func nginx() {
	cmd := exec.Command("/usr/local/bin/lunchy", "restart", "nginx") //重启命令根据自己的需要自行调整
	cmd.Run()
}
