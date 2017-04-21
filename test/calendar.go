package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	calendar "google.golang.org/api/calendar/v3"
)

func init() {
	registerDemo("calendar", calendar.CalendarScope, calendarMain)
}

func errorLog(err error,msg string){
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func calendarMain(client *http.Client, argv []string) {
	if len(argv) != 0 {
		fmt.Fprintln(os.Stderr, "Usage: calendar")
		return
	}

	svc, err := calendar.New(client)
	errorLog(err,"Unable to create calendar service: ")

	listRes, err := svc.CalendarList.List().Fields("items/id").Do()
	errorLog(err,"Unable to retrieve list of calendars: ")

	for _,v := range listRes.Items {
		log.Printf("Calendar ID: %v\n", v.Id)
	}

	if len(listRes.Items) > 0 {
		id := listRes.Items[0].Id
		res, err := svc.Events.List(id).Fields("items(summary,location,start,end)").Do()
		errorLog(err,"Unable to retrieve calendar events list: ")

		for _, v := range res.Items {
			log.Printf("\n内容: %q\n場所: %q\n開始時間(datetime): %v\n終了時間(datetime): %v\n\n",v.Summary, v.Location,v.Start.DateTime,v.End.DateTime)
		}
	}
}
