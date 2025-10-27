package models

import (
	"strconv"
	"time"
)

type Admin struct {
	Personne
	DateCreationCompte time.Time `json:"date_creation_compte"`
	Avocats            []*Avocat `gorm:"foreignKey:admin_id" json:"avocats"`
	UserID             uint      `json:"user_id"`
	User               *User     `gorm:"constraint:onUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}

func (a *Admin) GetRoleName() Role { return AdminRole }
func (a *Admin) SetUserID(id uint) { a.UserID = id }
func (a *Admin) IsEqual(other Admin) bool {
	const tolerance = time.Microsecond 
	diff := a.DateCreationCompte.Sub(other.DateCreationCompte)
	if diff < 0 {
		diff = -diff
	}

	return a.Personne.IsEqual(other.Personne) &&
		a.UserID == other.UserID &&
		diff < tolerance
}

func (a *Admin) IsEmpty() bool {
	return a.Personne.IsEmpty() && a.UserID == 0 && a.DateCreationCompte.IsZero()
}
func (a *Admin) ToString() string {
	return "Admin{" + a.Personne.ToString() + ",DateCreationCompte:" + a.DateCreationCompte.String() + ",UserID:" + strconv.FormatUint(uint64(a.UserID), 10) + "}"
}
