package models

import (
	"fmt"
)

type Personne struct {
	Id              uint   `gorm:"primaryKey" json:"id"`
	Nom             string `gorm:"type:varchar(30)" json:"nom"`
	Prenom          string `gorm:"type:varchar(30)" json:"prenom"`
	Age             uint8  `json:"age"`
	NumeroTelephone string `gorm:"type:varchar(10)" json:"numero_telephone"`
}

func (p *Personne) IsEqual(other Personne) bool {
	return p.Nom == other.Nom &&
		p.Prenom == other.Prenom &&
		p.Age == other.Age &&
		p.NumeroTelephone == other.NumeroTelephone
}

func (p *Personne) IsEmpty() bool {
	return p.Nom == "" && p.Prenom == "" && p.Age == 0 && p.NumeroTelephone == ""
}
func (p *Personne) ToString() string {
	return fmt.Sprintf("Id:%d, Nom:%s, Prenom:%s, Age:%d, NumeroTelephone:%s", p.Id, p.Nom, p.Prenom, p.Age, p.NumeroTelephone)
}
