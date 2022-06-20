package main

import (
	initMod "SE/src/init"
	"fmt"
)

//	"net/http"
//
//	"github.com/gin-gonic/gin"

func main() {

	// init system
	if err := initMod.Init(ConfPath); err != nil {
		fmt.Println(err.Error())
		return
	}

	//}
	//	router := gin.Default()
	//
	//	router.GET("/test", func(c *gin.Context) {
	//		c.IndentedJSON(http.StatusOK, "Test OK")
	//	})
	//
	//	err := router.Run(":8080")
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//db.TestDB()

}
