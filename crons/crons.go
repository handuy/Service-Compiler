package crons

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os"
	"time"
	"github.com/sirupsen/logrus"
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
	user := flag.String("user", "dev", "user name of container ")
	pDur := flag.String("pdur", "4s", "time duration of process ")
	interval := flag.String("interval", "3s", "time duration to reload delete old file")
	fDur := flag.String("fdur", "1h", "old file is exist duration")
	maxcpu := flag.Float64("maxcpu", 60, "max percent of cpu ")
	maxram := flag.Float64("maxram", 60, "max percent of ram ")
	flag.Parse()
	ps, _ := time.ParseDuration(*pDur)
	itv, _ := time.ParseDuration(*interval)
	fd, _ := time.ParseDuration(*fDur)
	cron:= &Cron{
		Pwd:        *path,
		UserName:   *user,
		MaxCPU:     *maxcpu,
		MaxRAM:     float32(*maxram),
		ProcessDur: ps,
		Interval:   itv,
		Fdur:       fd,
	}
	fmt.Println(cron)
	return cron
}

func (t *Cron) Remove(c time.Time) {
	logrus.Infoln("Tìm kiếm file đã hết hạn: ")
	info, err := ioutil.ReadDir(t.Pwd)
	if err != nil {
		fmt.Errorf("Không đọc được thư mục %s \n Message: %s", t.Pwd, err.Error())
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
	logrus.Infoln("Kiểm tra các Process: ")
	infos, err := process.Processes()
	if err != nil {
		fmt.Errorf("Không lấy được danh sách các Process %s \n Message: %s", t.Pwd, err.Error())
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

func (t *Cron) String() string{
	return fmt.Sprintf(`
		Cron job infomation:

		Path: 				%s
		UserName :			%s
		MaxCPU: 			%v
		MaxRam: 			%v
		Process Duration: 	%v
		Interval: 			%v
		Filde Duration: 	%v

`,t.Pwd,t.UserName,t.MaxCPU,t.MaxRAM,t.ProcessDur,t.Interval,t.Fdur)
}