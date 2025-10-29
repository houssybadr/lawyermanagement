package models

import (
	"encoding/json"
	"errors"
)

type SpecialiteAvocat int

const (
	Generaliste SpecialiteAvocat = iota+1
	DroitAffaires
	DroitCivil
	DroitTravail
	DroitProprieteIntellectuelle
	Criminalite
)

func (s SpecialiteAvocat) String() string {
	switch s {
	case Criminalite:
		return "Criminalite"
	case DroitAffaires:
		return "DroitAffaires"
	case DroitCivil:
		return "DroitCivil"
	case DroitTravail:
		return "DroitTravail"
	case DroitProprieteIntellectuelle:
		return "DroitProprieteIntellectuelle"
	case Generaliste:
		return "Generaliste"
	default:
		return "Unknown"
	}
}

func (s SpecialiteAvocat) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SpecialiteAvocat) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	switch str {
	case "Criminalite":
		*s = Criminalite
	case "DroitAffaires":
		*s = DroitAffaires
	case "DroitCivil":
		*s = DroitCivil
	case "DroitTravail":
		*s = DroitTravail
	case "DroitProprieteIntellectuelle":
		*s = DroitProprieteIntellectuelle
	case "Generaliste":
		*s = Generaliste
	default:
		return errors.New("unknown specialite")
	}
	return nil
}
