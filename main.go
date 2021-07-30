package main

import (
	"flag"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "The address to listen on for HTTP requests.")
	url := flag.String("url", "localhost", "The service URL ")
	flag.Parse()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
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

	})
	r.Run(":" + *port)

}
