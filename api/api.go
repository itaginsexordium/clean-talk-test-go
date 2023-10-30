package api

import (
	"net"
	"regexp"
	"net/http"
	"encoding/json"
	"strings"
	
	log "github.com/sirupsen/logrus"
	"github.com/labstack/echo/v4"
	"github.com/itaginsexordium/clean-talk-test-go/config"
	"github.com/itaginsexordium/clean-talk-test-go/storage"
	"github.com/oschwald/geoip2-golang"
)

var echoRouteRegex = regexp.MustCompile(`(?P<start>.*):(?P<param>[^\/]*)(?P<end>.*)`)

type GeoIpAPI struct {
	config *config.Config
	mc     *storage.MemcacheClient
	db     *geoip2.Reader
	echo   *echo.Echo
}

func New(config *config.Config, mc *storage.MemcacheClient , db *geoip2.Reader ) *GeoIpAPI {
	echo := echo.New()
	api := &GeoIpAPI{
		config: config,
		mc:     mc,
		db: 	db,
		echo:   echo,
	}

	echo.GET("/", api.getRoot)
	return api
}

func (api *GeoIpAPI) Start() error {
	return api.echo.Start(":"+api.config.HTTPBindAddr)
}

func (api *GeoIpAPI) getRoot(c echo.Context) error {
	rIp := c.QueryParam("ip")
	if rIp == "" || rIp == "null" {
		return c.String(http.StatusInternalServerError, "ip param is required and can't be empty or null")
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	jsonData, err := api.mc.Get(rIp)
	if err == nil {
		log.Info("from mem:")
		return c.JSONBlob(http.StatusOK, jsonData)
	} 

	parts := strings.Split(rIp, "/")
	if len(parts) != 2 {
		log.Info("Invalid IP/CIDR format:")
		return c.String(http.StatusInternalServerError, "Invalid IP/CIDR format")
	}

	ipStr := parts[0]
	// Parse the IP address
	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Info("Invalid IP address:")
		return c.String(http.StatusInternalServerError, "Invalid IP address")
	}

	// Parse the CIDR notation
	_, ipNet, err := net.ParseCIDR(rIp)
	if err != nil {
		log.Info("Invalid CIDR notation:")
		return c.String(http.StatusInternalServerError, "Invalid CIDR notation")
	}

	enty, err := api.db.City(ipNet.IP)
	if err != nil {
		log.Info(err)
	}

	jsonData, err = json.Marshal(enty)
	if err != nil {
		log.Info("Error marshaling JSON:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	log.Info("set mem")
	api.mc.Set(rIp, jsonData)
	return c.JSONBlob(http.StatusOK, jsonData)
}
 
