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
	port = flag.String("port", "8080", "The address to listen on for HTTP requests.")
	url  = flag.String("url", "localhost", "The service URL ")
)

func main() {

	flag.Parse()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", getAll)
	r.GET("/ip", get_ip)
	r.GET("/country", get_country)
	r.GET("/city", get_city)
	r.GET("/as", get_as)
	r.GET("/asFull", get_asFull)
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

		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("\n\r"+outputString))

	} else {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
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

func get_as(c *gin.Context) {
	if strings.Contains(c.GetHeader("User-Agent"), "curl") {
		reqResponse := get_from_api(c.ClientIP())
		str := strings.Fields(reqResponse.As)
		c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(str[0]))
	}
}

func get_asFull(c *gin.Context) {
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
