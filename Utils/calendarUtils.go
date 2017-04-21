package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

const CALENDAR_ID string = "osaka.hal.iw13a727@gmail.com"

type Schedule struct {
	Title string
	Location string
	Year string
	Month string
	Day string
	Start string
	End string
}

func SetSchedule(schedule []string) *Schedule {
	return &Schedule{
		Title: schedule[0],
		Location: schedule[1],
		Year: schedule[2],
		Month: schedule[3],
		Day: schedule[4],
		Start: schedule[5],
		End: schedule[6],
	}
}

func CreateEvent(schedule *Schedule,ch chan string) {
	ctx := context.Background()
	b, err := ioutil.ReadFile("client_secret.json")
	errorLog("Unable to read client secret file: ", err)

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	errorLog("Unable to parse client secret file to config: ", err)

	client := getClient(ctx, config)
	srv, err := calendar.New(client)
	errorLog("Unable to retrieve calendar Client: ", err)

	_, err = srv.Events.Insert(CALENDAR_ID, createEventData(schedule)).Do()
	errorLog("Unable to create event. ", err)

	ch <- schedule.Title
	return
}

func TestRun(s *Schedule, ch chan string){
	time.Sleep(1 * time.Second)
	ch <- s.Title
	return
}

func errorLog(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	errorLog("Unable to get path to cached credential file.: ", err)

	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	errorLog("Unable to retrieve token from web: ", err)

	return tok
}

func tokenCacheFile() (string, error) {
	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir, url.QueryEscape("generate-schedule-calendar.json")), err
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)
	errorLog("Unable to cache oauth token: ", err)

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func createEventData(schedule *Schedule) *calendar.Event {

	start_datatime := schedule.Year + "-" + schedule.Month + "-" + schedule.Day + "T" + schedule.Start + ":00+09:00"
	end_datatime := schedule.Year + "-" + schedule.Month + "-" + schedule.Day + "T" + schedule.End + ":00+09:00"

	event := &calendar.Event{
		Summary:  schedule.Title,
		Location: schedule.Location,
		Start: &calendar.EventDateTime{
			DateTime: start_datatime,
			TimeZone: "Asia/Tokyo",
		},
		End: &calendar.EventDateTime{
			DateTime: end_datatime,
			TimeZone: "Asia/Tokyo",
		},
	}

	return event
}