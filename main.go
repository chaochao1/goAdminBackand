package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lwnmengjing/goAdminBackand/routers"
	"github.com/lwnmengjing/goAdminBackand/utils"
	"log"
	"net/http"
)

func main() {
	port := utils.Config.HttpPort
	if port == "" {
		port = ":8080"
	}
	if err := http.ListenAndServe(port, routers.Router); err != nil {
		log.Fatalln(err)
	}
}
