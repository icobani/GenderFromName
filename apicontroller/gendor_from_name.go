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
	"log"
	"net/http"
)

func GET_GenderFromName(c *gin.Context) {
	var name = c.Params.ByName("name")
	//AppConfig.DB.
	gender := models.Gender{}
	profile := []models.Profile{}

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
			gender.Probability = float64(gender.FemaleCount) / float64(gender.TotalCount)
			log.Println(gender.FemaleCount, gender.MaleCount, gender.TotalCount, gender.Probability)
		} else {
			gender.Gender = "Male"
			gender.Probability = float64(gender.MaleCount) / float64(gender.TotalCount)
		}
	} else {
		gender.Gender = ""
		gender.Probability = 0
	}
	c.JSON(http.StatusOK, gender)
}
