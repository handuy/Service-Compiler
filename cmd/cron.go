package main

import "git.hocngay.com/techmaster/service-complier/crons"

func main() {
	cron := crons.NewCron()
	cron.Run()
}
