package main;

import (
	log "github.com/sirupsen/logrus"
	"github.com/itaginsexordium/clean-talk-test-go/config"
	"github.com/itaginsexordium/clean-talk-test-go/storage"
	"github.com/itaginsexordium/clean-talk-test-go/api"
	"github.com/oschwald/geoip2-golang"
)

func main (){
	cnf, err := config.Get()
	if err != nil {
		log.Fatal("Failed to get config : ", err)
	}

	mcConfig := []string{cnf.MemcacheURL}
	mc := storage.NewMemcacheClient(mcConfig);	
	db, err := geoip2.Open(cnf.GeoIpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	api  := api.New(cnf, mc , db)
	if err = api.Start(); err != nil {
		log.Fatal("Failed to start  API: ", err)
	}
}