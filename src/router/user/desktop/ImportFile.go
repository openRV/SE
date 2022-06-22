package desktop

import (
	comInterface "SE/src/interface"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImportFile(c *gin.Context) {

	file, err := c.FormFile("FormData")
	if err != nil {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: "upload error"})
	}

	fileName := file.Filename
	fmt.Println(fileName)
	c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: true, Msg: fileName})
}
