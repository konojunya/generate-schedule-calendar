package main

import (
	"github.com/konojunya/generate-schedule-calendar/Utils"
)

func main() {
	ed := make(map[string]string)
	datas := Utils.Read("schedule.csv")

	for _,schedule := range datas {
		ed["Title"] = schedule[0]
		ed["Location"] = schedule[1]
		ed["Year"] = schedule[2]
		ed["Month"] = schedule[3]
		ed["Day"] = schedule[4]
		ed["Start"] = schedule[5]
		ed["End"] = schedule[6]
		Utils.CreateEvent(ed)
	}
}