package main

import (
  "fmt"
  "github.com/kkdai/twitter"
  "os"
)

type Credentials struct {
  ConsumerKey string
  ConsumerSecret string
}

func main(){
  creds := Credentials{
    ConsumerKey: os.Getenv("CONSUMER_KEY"),
    ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
  }

  Client = NewDesktopClient(creds.ConsumerKey, creds.ConsumerSecret)

  Client.DoAuth()
}
