package models

import "time"

type Document struct {
	Id           uint         `gorm:"primaryKey" json:"id"`
	Nom          string       `gorm:"type:varchar(20)" json:"nom" binding:"required"`
	Contenu      []byte       `gorm:"type:bytea" json:"-"`
	DateCreation time.Time    `json:"date_creation"`
	TypeFichier  TypeFichier  `gorm:"type:int;not null" json:"type_fichier"`
	TypeDocument TypeDocument `gorm:"type:int;default:0" json:"type_document"`
	DossierID    uint         `json:"dossier_id"`
	Dossier      *Dossier     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (d *Document)IsEmpty()bool{
	return d.Nom==""&&
		d.Contenu==nil &&
		d.DateCreation.IsZero()
}

func (d *Document)IsEqual(other Document)bool{
	tolerance:=time.Microsecond
	diff:=d.DateCreation.Sub(other.DateCreation)
	if diff<0{
		diff=-diff
	}
	return d.Nom==other.Nom&&
			d.TypeDocument==other.TypeDocument&&
			d.TypeFichier==other.TypeFichier&&
			diff<tolerance
}