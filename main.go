package main

import (
	"fmt"

	"github.com/konojunya/generate-schedule-calendar/Utils"
)

func main() {
	datas := Utils.Read("schedule.csv")

	for _, schedule := range datas {
		s := Utils.SetSchedule(schedule)
		Utils.CreateEvent(s)
	}

	fmt.Println("finished!")
}
