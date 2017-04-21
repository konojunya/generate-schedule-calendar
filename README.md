# generate-schedule-calendar

**HALの授業スケジュールをGoogle Calendar APIを叩いて、自動的に予定を追加するスクリプト**

## Library

[google-api-go-client](https://github.com/google/google-api-go-client)

## Install

```bash
go get golang.org/x/net/context
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
go get google.golang.org/api/calendar/v3
go get github.com/konojunya/generate-schedule-calendar/Utils
```

## Usage

```bash
$ go run main.go
```

or

```bash
$ go build main.go
$ ./main
```