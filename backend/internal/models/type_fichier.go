package models

import (
	"encoding/json"
	"errors"
)

type TypeFichier int

const(
	PNG TypeFichier =iota+1
	JPG
	PDF
	DOCX
	XLSX
	CSV
)

func (t *TypeFichier) String() (string,error){
	switch *t{
		case PNG:
			return "PNG",nil
		case JPG:
			return "JPG",nil
		case PDF:
			return "PDF",nil
		case DOCX:
			return "DOCX",nil
		case XLSX:
			return "XLSX",nil
		case CSV:
			return "CSV",nil
		default:
			return "",errors.New("unknown file type")
	}
}

func (t *TypeFichier) ToContentType() (string,error){
	switch *t{
		case PNG:
			return "image/png",nil
		case JPG:
			return "image/jpeg",nil
		case PDF:
			return "application/pdf",nil
		case CSV:
			return "text/csv",nil
		case DOCX:
			return "application/vnd.openxmlformats-officedocument.wordprocessingml.document",nil
		case XLSX:
			return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",nil
		default:
			return "",errors.New("unknown file type")
	}
}

func (t*TypeFichier) FromString(str string) error{
	switch str{
		case "PNG":
			*t=PNG
		case "JPG":
			*t=JPG
		case "PDF":
			*t=PDF
		case "DOCX":
			*t=DOCX
		case "XLSX":
			*t=XLSX
		case "CSV":
			*t=CSV
		default:
			return errors.New("invalid file type")
	}
	return nil
}

func (t *TypeFichier) ShouldBeDownloaded() bool{
	switch *t{
		case DOCX,XLSX:
			return true
		case PNG,JPG,PDF,CSV:
			return false
		default:
			return true
	}
}

func (t TypeFichier) MarshalJSON() ([]byte,error){
	str,err:=t.String()
	if err!=nil{
		return nil,err
	}
	return json.Marshal(str)
}

func (t *TypeFichier) UnmarshalJSON(data []byte) error{
	switch string(data){
		case "PNG":
			*t=PNG
		case "JPG":
			*t=JPG
		case "PDF":
			*t=PDF
		case "DOCX":
			*t=DOCX
		case "XLSX":
			*t=XLSX
		case "CSV":
			*t=CSV
		default:
			return errors.New("unknown file type")
	}
	return nil
}

