package main

import (
	"fmt"

	"github.com/konojunya/generate-schedule-calendar/Utils"
)

func main() {
	datas := Utils.Read("schedule.csv")
	ch := make(chan string,len(datas))

	for _, schedule := range datas {
		s := Utils.SetSchedule(schedule)
		go Utils.TestRun(s,ch)
	}

	for out := range ch {
		fmt.Printf("Set Schedule\n科目: %s\n\n",out)
	}
}
