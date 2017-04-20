package fileUtils

import (
  "encoding/csv"
  "log"
  "os"
)

func failOnError(err error) {
  if err != nil {
    log.Fatal("Error: ",err)
  }
}

func Read(filename string) [][]string {
  fr, err := os.Open(filename)
  failOnError(err)
  defer fr.Close()

  r := csv.NewReader(fr)
  rows, err := r.ReadAll()
  failOnError(err)

  var datas [][]string

  for _, row := range rows {
    datas = append(datas,row)
  }

  return datas
}