package models

import "encoding/json"

type TypeDocument int

const (
	OTHER TypeDocument = iota
	PREUVE
	CONTRAT
	FACTURE
	COURRIER
)

func (t *TypeDocument) String() string {
	switch *t {
	case PREUVE:
		return "PREUVE"
	case CONTRAT:
		return "CONTRAT"
	case FACTURE:
		return "FACTURE"
	case COURRIER:
		return "COURRIER"
	default:
		return "OTHER"
	}
}

func (t *TypeDocument) FromString(str string) {
	switch str {
	case "PREUVE":
		*t = PREUVE
	case "CONTRAT":
		*t = CONTRAT
	case "FACTURE":
		*t = FACTURE
	case "COURRIER":
		*t = COURRIER
	default:
		*t = OTHER
	}
}

func (t TypeDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *TypeDocument) UnmarshalJSON(data []byte) error{
	switch string(data) {
	case "PREUVE":
		*t = PREUVE
	case "CONTRAT":
		*t = CONTRAT
	case "FACTURE":
		*t = FACTURE
	case "COURRIER":
		*t = COURRIER
	default:
		*t = OTHER
	} 
	return nil
}
