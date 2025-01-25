package main

import (
	ws "SnappChatroom/api/handlers/websocket"
	"SnappChatroom/app"
	"SnappChatroom/config"
	"flag"
	"log"
	"os"
)

var configPath = flag.String("config", "config.json", "service configuration file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.NewMustApp(c)

	log.Fatal(ws.Run(appContainer, c.Server))
}
