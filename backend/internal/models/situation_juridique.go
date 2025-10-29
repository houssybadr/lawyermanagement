package models

import "encoding/json"

type SituationJuridique int

const (
	// SituationAttente - Situation intermédiaire et provisoire
	SituationAttente SituationJuridique = iota+1
	
	// DroitSubjectif - Pouvoir reconnu à une personne par la loi
	DroitSubjectif 

	// Obligation - Contrainte légale qui pèse sur une personne
	Obligation

	// SimpleEsperance - Simple possibilité sans garantie juridique
	SimpleEsperance
)

func (s SituationJuridique) String() string {
	switch s {
	case DroitSubjectif:
		return "DroitSubjectif"
	case Obligation:
		return "Obligation"
	case SituationAttente:
		return "SituationAttente"
	case SimpleEsperance:
		return "SimpleEsperance"
	default:
		return "Inconnu"
	}
}

func (s SituationJuridique) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SituationJuridique) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	switch str {
	case "DroitSubjectif":
		*s = DroitSubjectif
	case "Obligation":
		*s = Obligation
	case "SituationAttente":
		*s = SituationAttente
	case "SimpleEsperance":
		*s = SimpleEsperance
	default:
		*s = -1
	}
	return nil
}
