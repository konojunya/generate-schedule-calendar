package main

import (
  "./Utils"
  "fmt"
)

func main() {
  datas := fileUtils.Read("schedule.csv")
  fmt.Println(datas)
}