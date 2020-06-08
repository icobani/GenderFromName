/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 08.06.2020  17:58
   Notes   :
*/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/icobani/GenderFromName/AppConfig"
	"github.com/icobani/GenderFromName/apiroots"
)

func main() {
	app := gin.Default()
	api := app.Group("api")
	apiroots.GenderFromNameApi(api)

	// Database bağlantısını ayağa kaldırma
	AppConfig.InitDB()
	app.Run(":8014")
}
