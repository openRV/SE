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
		return
	}

	fileName := file.Filename
	fmt.Println(fileName)

	f, err := file.Open()
	if err != nil {
		c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: "upload error"})
		return
	}
	defer f.Close()
	buf := make([]byte, file.Size)
	var chunk []byte
	for {
		n, err := f.Read(buf)
		if err != nil && err.Error() != "EOF" {
			fmt.Println(err)
			c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: false, Msg: "upload error"})
			return
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}

	fmt.Println(string(chunk))

	c.IndentedJSON(http.StatusOK, comInterface.ErrorRes{Success: true, Msg: fileName})
}
