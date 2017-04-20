package main

import (
	"./Utils"
	"fmt"
)

func main() {
	schedules := Utils.Read("schedule.csv")

	for _, schedule := range schedules {
		fmt.Println(schedule)
	}
}
