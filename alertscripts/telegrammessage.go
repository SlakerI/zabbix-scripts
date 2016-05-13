package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bot-api/telegram"
	"github.com/kardianos/osext"
	"github.com/naoina/toml"
	"golang.org/x/net/context"
)

// Config Info
type Config struct {
	Token string
	Debug bool
}

func main() {

	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(folderPath + "/telegrammessage.conf")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}

	api := telegram.New(config.Token)
	api.Debug(config.Debug)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := telegram.MessageCfg{
		Text: os.Args[2] + "\n" + os.Args[3],
	}
	cfg.ChannelUsername = os.Args[1]

	_, err = api.SendMessage(ctx, cfg)

	if err != nil {
		log.Println(err.Error())
	}

}
