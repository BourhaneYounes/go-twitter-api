package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Credentials struct {
  ConsumerKey string
  ConsumerSecret string
  AccessToken string
  AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error){
  config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)

  token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

  httpClient := config.Client(oauth1.NoContext, token)
  client := twitter.NewClient(httpClient)

  verify := &twitter.AccountVerifyParams{
    SkipStatus: twitter.Bool(true),
    IncludeEmail: twitter.Bool(true),
  }

  user, _, err := client.Accounts.VerifyCredentials(verify)
  if err != nil {return nil, err}
  log.Printf("user : \n%+v\n",user)
  return client, nil
}

func main(){

  creds := Credentials{
    AccessToken: os.Getenv("ACCES_TOKEN"),
    AccessTokenSecret: os.Getenv("ACCES_TOKEN_SECRET"),
    ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
    ConsumerKey: os.Getenv("CONSUMER_KEY"),
  }
	
	fmt.Printf("%+v",creds)

  client, err := getClient(&creds)
  if err != nil {
    log.Printf("ERROR")
    log.Println(err)
  }

  fmt.Printf("Client: %+v\n",client)

	tweet, resp, err := client.Statuses.Update("test pour voir si ça tweet", nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n",resp)
	log.Printf("%+v\n",tweet)
}
