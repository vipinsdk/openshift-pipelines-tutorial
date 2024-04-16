package main

import (
	"example_server/route"
	"os"
	"log"
)

var (
	Infolog  *log.Logger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Debuglog *log.Logger = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warnlog  *log.Logger = log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile)
	Errorlog *log.Logger = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	port := ":8888"
	router := route.SetupRouter()
	Infolog.Println("Http server listening on 8888")
	router.Run(port)
}