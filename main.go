package main

import (
	"fmt"

	"github.com/konojunya/generate-schedule-calendar/Utils"
)

func main() {
	datas := Utils.Read("schedule.csv")
	calendarId := Utils.GetCalendarId()
	for _, schedule := range datas {
		s := Utils.SetSchedule(schedule)
		Utils.CreateEvent(s, calendarId)
	}

	fmt.Println("finished!")
}
