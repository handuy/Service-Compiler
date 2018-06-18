package crons

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os"
	"time"
)

type Cron struct {
	Interval time.Duration
	Pwd      string        //path of working folder
	Fdur     time.Duration //file duration
	UserName string
	MaxCPU   int
	MaxRAM   int
}

func (t *Cron) Remove(c time.Time) {
	time.Sleep(1 * time.Hour)
	info, err := ioutil.ReadDir(t.Pwd)
	if err != nil {
		fmt.Errorf("Error read dir %s \n Message: %s", t.Pwd, err.Error())
		return
	}
	for _, inf := range info {
		if time.Since(inf.ModTime()) > 1*time.Hour {
			fmt.Printf(" %s Tìm thấy file:  %s \n", c.Format("2006-01-02T15:04:05"), inf.Name())
			os.Remove(t.Pwd + "/" + inf.Name())
		}
	}
}

func (t *Cron) CheckCPU() {
	infos, err := process.Processes()
	if err != nil {
		fmt.Errorf("Error read dir %s \n Message: %s", t.Pwd, err.Error())
		return
	}
	for _, info := range infos {

		name, _ := info.Name()
		userName, _ := info.Username()
		percent, _ := info.CPUPercent()
		if userName == t.UserName {
			fmt.Printf("\r PID %v Name %v UserName %v  Percent %v \n", info.Pid, name, userName, percent)
		}

	}

}

func (t *Cron) Run() {
	//tick := time.Tick(t.Interval)
	//for c := range tick {
	//	t.Remove(c)
	//
	//}
	t.CheckCPU()
}
