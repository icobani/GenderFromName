/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 08.06.2020  17:56
   Notes   :
*/
package apiroots

import (
	"github.com/gin-gonic/gin"
	"github.com/icobani/GenderFromName/apicontroller"
)

func GenderFromNameApi(api *gin.RouterGroup) {
	api.GET("/genderfromname/:name", apicontroller.GET_GenderFromName)
}
