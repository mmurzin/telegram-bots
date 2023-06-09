package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	tele "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	pref := tele.Settings{
		Token:  os.Getenv("MMURZIN_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello From Docker <3!")
	})

	b.Start()
}
