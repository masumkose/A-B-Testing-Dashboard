package handlers

import "gorm.io/gorm"

type Experiment struct {
	gorm.Model
	Name       		string 		`gorm:"unique;not null"`
	Variations		[]Variation	`gorm:"foreignKey:ExperimentID"`
}

type Variation struct {
	gorm.Model
	Name         	string 		`gorm:"not null"`
	Participants	uint		`gorm:"default:0"`
	Conversions		uint		`gorm:"default:0"`
	ExperimentID	uint		
}