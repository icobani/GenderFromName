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
	"github.com/icobani/GenderFromName/AppConfig"
	"github.com/icobani/GenderFromName/models"
	"net/http"
)

func GET_GenderFromName(c *gin.Context) {
	var name = c.Params.ByName("name")
	//AppConfig.DB.
	gender := models.Gender{}
	profile := models.Profile{}

	AppConfig.DB.Where("first_name =? AND (gender = ? OR gender = ?)",
		name, "male", "erkek").
		Find(&profile).Count(&gender.MaleCount)

	AppConfig.DB.Where("first_name =? AND (gender = ? OR gender = ?)",
		name, "female", "kadın").
		Find(&profile).Count(&gender.FemaleCount)

	gender.TotalCount = gender.FemaleCount + gender.MaleCount
	if gender.TotalCount > 0 {

		if gender.FemaleCount > gender.MaleCount {
			gender.Gender = "Female"
			gender.Probability = float64(gender.FemaleCount / gender.TotalCount)
		} else {
			gender.Gender = "Male"
			gender.Probability = float64(gender.MaleCount / gender.TotalCount)
		}
	} else {
		gender.Gender = ""
		gender.Probability = 0
	}
	c.JSON(http.StatusOK, gender)
}
