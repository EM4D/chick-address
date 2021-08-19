package main

import (
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	port = flag.String("port", "9999", "The address to listen on for HTTP requests.")
	url  = flag.String("url", "localhost", "The service URL ")
)

func main() {

	flag.Parse()

	r := gin.Default()
	r.LoadHTMLGlob("web-interface/*")

	r.GET("/", getAll)
	r.GET("/ip", get_ip)
	r.GET("/country", get_country)
	r.GET("/city", get_city)
	r.GET("/asn", get_asn)
	r.GET("/asnfull", get_asnfull)
	r.GET("/timezone", get_timezone)
	r.Static("/assets", "./assets")
	r.Run(":" + *port)

}

func getAll(c *gin.Context) {
	reqResponse := get_from_api(c.ClientIP())
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {

		var outputString string
		v := reflect.ValueOf(reqResponse)

		for i := 0; i < v.NumField(); i++ {
			if len(v.Field(i).String()) > 0 {
				outputString += "\033[96m" + v.Type().Field(i).Name + ": \033[92m" + fmt.Sprintf("%v", v.Field(i)) + "\033[37m \n\r"
			}

		}
		outputString += "\n\033[93mSponsored with \033[0;31mLOVE\033[93m by Pardis Co.\n\rwww.Pardisco.co\n\rSpecial Thanks to Twitter@EmadMahmoudpour \033[37m" + "\n\r" + "For IP Only use $curl myadd.ir/ip"
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("\n\r"+outputString))

	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"data": reqResponse,
			"url":  *url,
		})

	}

}

func get_ip(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(c.ClientIP()))
	}
}

func get_country(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(reqResponse.Country))
	}
}

func get_city(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(reqResponse.City))
	}
}

func get_asn(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		str := strings.Fields(reqResponse.As)
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(str[0]))
	}
}

func get_asnfull(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(reqResponse.As))
	}
}

func get_timezone(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(reqResponse.Timezone))
	}
}
