package crons

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os"
	"time"
)

type Cron struct {
	Interval   time.Duration
	Pwd        string        //path of working folder
	Fdur       time.Duration //file duration
	UserName   string
	MaxCPU     float64
	MaxRAM     float32
	ProcessDur time.Duration
}

func NewCron() *Cron {
	path := flag.String("path", "/root/temp", "path of folder contains files are complied ")
	user := flag.String("path", "root", "path of folder contains files are complied ")
	pDur := flag.String("path", "4s", "path of folder contains files are complied ")
	interval := flag.String("interval", "3s", "time duration to reload delete old file")
	fDur := flag.String("fdur", "1h", "old file is exist duration")
	maxcpu := flag.Float64("maxcpu", 60, "max percent of cpu ")
	maxram := flag.Float64("maxram", 60, "max percent of ram ")
	flag.Parse()
	ps, _ := time.ParseDuration(*pDur)
	itv, _ := time.ParseDuration(*interval)
	fd, _ := time.ParseDuration(*fDur)
	return &Cron{
		Pwd:        *path,
		UserName:   *user,
		MaxCPU:     *maxcpu,
		MaxRAM:     float32(*maxram),
		ProcessDur: ps,
		Interval:   itv,
		Fdur:       fd,
	}
}

func (t *Cron) Remove(c time.Time) {
	time.Sleep(1 * time.Hour)
	info, err := ioutil.ReadDir(t.Pwd)
	if err != nil {
		fmt.Errorf("Error read dir %s \n Message: %s", t.Pwd, err.Error())
		return
	}
	for _, inf := range info {
		if time.Since(inf.ModTime()) > t.Fdur {
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
		percentRam, _ := info.MemoryPercent()
		createdTime, _ := info.CreateTime()
		duration := time.Since(time.Unix(createdTime, 0))
		if userName == t.UserName && percent >= t.MaxCPU || percentRam >= t.MaxRAM && duration >= t.ProcessDur {
			fmt.Printf("\r PID %v Name %v UserName %v  Percent %v \n", info.Pid, name, userName, percent)
			info.Kill()
		}
	}

}

func (t *Cron) Run() {
	tick := time.Tick(t.Interval)
	for c := range tick {
		t.Remove(c)
		t.CheckCPU()
	}
}
