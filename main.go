package main

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

  "golang.org/x/net/context"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/calendar/v3"
)

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
  cacheFile, err := tokenCacheFile()
  if err != nil {
    log.Fatalf("Unable to get path to cached credential file. %v", err)
  }

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
  if err != nil {
    log.Fatalf("Unable to retrieve token from web %v", err)
  }

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
  if err != nil {
    log.Fatalf("Unable to cache oauth token: %v", err)
  }
  defer f.Close()
  json.NewEncoder(f).Encode(token)
}

func main() {
    ctx := context.Background()
    b, err := ioutil.ReadFile("client_secret.json")
    if err != nil {
        log.Fatalf("Unable to read client secret file: %v", err)
    }
    config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
    if err != nil {
        log.Fatalf("Unable to parse client secret file to config: %v", err)
    }
    client := getClient(ctx, config)
    srv, err := calendar.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve calendar Client %v", err)
    }

    event := &calendar.Event{
        Summary:     "Sample event",
        Location:    "Sample location",
        Description: "This is a sample event.",
        Start: &calendar.EventDateTime{
            DateTime: "2017-04-22T00:00:00+09:00",
            TimeZone: "Asia/Tokyo",
        },
        End: &calendar.EventDateTime{
            DateTime: "2017-04-22T01:00:00+09:00",
            TimeZone: "Asia/Tokyo",
        },
    }
    calendarID := "osaka.hal.iw13a727@gmail.com"
    event, err = srv.Events.Insert(calendarID, event).Do()
    if err != nil {
        log.Fatalf("Unable to create event. %v\n", err)
    }
    fmt.Printf("Event created: %s\n", event.HtmlLink)
}
