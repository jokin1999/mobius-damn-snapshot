package bunny

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Init() {
	// set debug mode
	gin.SetMode(gin.DebugMode)

	// initialize routes
	router := InitRouter()

	// set endpoint

	// set timeout

	// run http server
	server := &http.Server{
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	fmt.Printf("[info] start http server listening %s\n", ":8082")
	server.ListenAndServe()
}
