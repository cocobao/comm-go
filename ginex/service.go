package ginex

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RunService(port int, router *gin.Engine) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		fmt.Println(err, "setup secure api service fail")
		os.Exit(0)
	}
}
