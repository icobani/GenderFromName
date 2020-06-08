/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 08.06.2020  18:15
   Notes   :
*/
package models

type Gender struct {
	Gender      string  `json:"gender"`
	TotalCount  int     `json:"total_count"`
	MaleCount   int     `json:"male_count"`
	FemaleCount int     `json:"female_count"`
	Probability float64 `json:"probability"`
}
