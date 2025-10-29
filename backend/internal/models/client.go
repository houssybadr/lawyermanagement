package models

import (
	"fmt"
)

type Client struct {
	Personne
	Profession         string             `json:"profession"`
	SituationJuridique SituationJuridique `gorm:"type:int;default:1" json:"situation_juridique" `
	AvocatID           uint               `json:"avocat_id"`
	Avocat             *Avocat            `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"-"`
	Dossiers           []*Dossier         `gorm:"foreignKey:client_id" json:"dossiers"`
}

func (c *Client) ToString() string {
	return fmt.Sprintf("Client{%s, Profession: %s, SituationJuridique: %s, AvocatId:%d}", c.Personne.ToString(), c.Profession, c.SituationJuridique, c.AvocatID) 
}

func (c *Client) IsEmpty() bool {
	return c.Personne.IsEmpty() &&
		c.Profession == "" &&
		c.AvocatID == 0
}

func (c *Client) IsEqual(other Client) bool {
	return c.Personne.IsEqual(other.Personne) &&
		c.Profession == other.Profession &&
		c.SituationJuridique == other.SituationJuridique &&
		c.AvocatID == other.AvocatID
}
