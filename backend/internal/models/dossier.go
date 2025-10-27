package models

import (
	"fmt"
	"time"
)

type Dossier struct {
	Id           uint        `gorm:"primaryKey" json:"id"`
	Titre        string      `gorm:"type:varchar(20)" json:"titre" binding:"required"`
	Description  string      `json:"Description" binding:"required"`
	DateCreation time.Time   `json:"date_creation"`
	ClientID     uint        `json:"client_id"`
	Client       *Client     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Documents    []*Document `gorm:"foreignKey:dossier_id" json:"documents"`
}

func (d *Dossier) ToString() string {
	return fmt.Sprintf("Dossier{id:%d, titre:%s, description:%s, DateCreation:%s}", d.Id, d.Titre, d.Description, d.DateCreation.String())
}

func (d *Dossier) IsEmpty() bool {
	return d.Titre == "" &&
		d.Description == "" &&
		d.DateCreation.IsZero()
}

func (d *Dossier) IsEqual(other Dossier) bool {
	const tolerance = time.Microsecond
	diff := d.DateCreation.Sub(other.DateCreation)
	if diff < 0 {
		diff = -diff
	}
	return d.Titre == other.Titre &&
		d.Description == other.Description &&
		diff < tolerance
}
