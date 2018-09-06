package main

import (
	"net/http"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lwnmengjing/goAdminBackand/utils"
	"github.com/lwnmengjing/goAdminBackand/routers"
)

func main() {
	port := utils.Config.HttpPort
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, routers.Router); err != nil {
		log.Fatalln(err)
	}
}
