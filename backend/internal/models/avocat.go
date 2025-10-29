package models

type Avocat struct {
	Personne
	Cabinet       string           `gorm:"type:varchar(100)" json:"cabinet"`
	NumeroBarreau string           `gorm:"type:varchar(100)"  json:"numero_barreau"`
	Specialite    SpecialiteAvocat `gorm:"type:int;default:1" json:"specialite" `
	AdminID       uint             `json:"admin_id"`
	Admin         *Admin           `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;" json:"-"`
	Clients       []*Client        `gorm:"foreignKey:avocat_id" json:"clients"`
	UserID        uint             `json:"user_id"`
	User          *User            `gorm:"constraint:onUpdate:CASCADE, OnDelete:CASCADE;" json:"-"`
}

func (a *Avocat) GetRoleName() Role { return AvocatRole }
func (a *Avocat) SetUserID(id uint) { a.UserID = id }
func (a *Avocat) IsEmpty() bool {
	return a.Personne.IsEmpty() &&
		a.Cabinet == "" &&
		a.NumeroBarreau == "" &&
		a.AdminID == 0 &&
		a.UserID == 0
}

func (a *Avocat) IsEqual(other Avocat) bool{
	return a.Personne.IsEqual(other.Personne) &&
		a.Cabinet==other.Cabinet&&
		a.NumeroBarreau==other.NumeroBarreau&&
		a.Specialite==other.Specialite&&
		a.AdminID==other.AdminID&&
		a.UserID==other.UserID
}