/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 08.06.2020  17:53
   Notes   :
*/
package apicontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GET_GenderFromName(c *gin.Context) {
	c.JSON(http.StatusOK, "{'data':'hey'}")
}
