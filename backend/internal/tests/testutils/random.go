package testutils

import (
	"math/rand"
	"time"
	"strconv"
	"test/internal/models"
)

func generateRandomNom() string{
	rand.Seed(time.Now().UnixNano())
	noms:=[]string{"badr","omar","yassine","sara","lina","adam","nora","khalid","salma","hassan","oussama","amina"}
	return noms[rand.Intn(len(noms))]
}

func generateRandomPrenom() string{
	rand.Seed(time.Now().UnixNano())
	prenoms:=[]string{"houssy","bennani","elkhattabi","elidrissi","mohammedi","elghazali","fassi","elouardi","chater","zaidi","eljamai","elbouzidi"}
	return prenoms[rand.Intn(len(prenoms))]
}

func generateRandomAge() uint8{
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(70-18)+18)
}

func generateRandomEnmail() string{
	rand.Seed(time.Now().UnixNano())
	domains:=[]string{"example.com","test.com","mail.com","demo.com","sample.com"}
	return generateRandomNom()+"."+generateRandomPrenom()+"@"+domains[rand.Intn(len(domains))]
}
func generateRandomNumeroTelephone() string{
	rand.Seed(time.Now().UnixNano())
	prefixes:=[]string{"05","07","06"}
	return prefixes[rand.Intn(len(prefixes))]+strconv.Itoa(rand.Intn(99999999-1000000)+1000000)
}


func GenerateRandomPassword() string{
	upperCaseCharsSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCaseCharsSet := "abcdefghijklmnopqrstuvwxyz"
	numberSet := "0123456789"
	specialCharsSet := "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	mdp_length := rand.Intn(12)+8
	allcharset := upperCaseCharsSet + lowerCaseCharsSet + numberSet + specialCharsSet
	var password string
	for i := 0; i < mdp_length; i++ {
		password+= string(allcharset[rand.Intn(len(allcharset))])
	}
	return password
}

func generateRandomPersonne() models.Personne{
	return models.Personne{
		Nom: generateRandomNom(),
		Prenom: generateRandomPrenom(),
		Age: generateRandomAge(),
		NumeroTelephone: generateRandomNumeroTelephone(),
	}
}

func generateRandomProfession()string{
	professions := []string{"Médecin","Avocat","Ingénieur logiciel","Architecte","Professeur","Infirmier","Comptable","Chef de projet","Designer graphique","Journaliste","Pharmacien","Consultant en management","Électricien","Mécanicien","Dentiste","Entrepreneur","Scientifique","Analyste financier","Développeur web","Technicien réseau","Chef cuisinier","Agriculteur","Pilote de ligne","Psychologue","Traducteur","Data Scientist","Agent immobilier","Chercheur","Responsable marketing","Bibliothécaire",}
	return professions[rand.Intn(len(professions))]
}

func generateRandomAdresse() string{
	rand.Seed(time.Now().UnixNano())
	villes:=[]string{"Casablanca","Rabat","Marrakech","Fes","Tanger","Agadir","Essaouira","Ouarzazate","Chefchaouen","Meknes"}
	rues:=[]string{"rue de la paix","avenue des champs","boulevard saint-michel","place de la concorde","rue victor hugo","avenue jean jaures","boulevard haussmann","place vendome","rue de rivoli","avenue montaigne"}
	return villes[rand.Intn(len(villes))]+rues[rand.Intn(len(rues))]+"n"+strconv.Itoa(rand.Intn(100)+1)
}

func generateRandomBytes() []byte {
	size:=rand.Intn(60-10)+10
    b := make([]byte, size)
    _, _ = rand.Read(b)
    return b
}

func GenerateRandomUser() models.User{
	return models.User{
		Email: generateRandomEnmail(),
		Password: GenerateRandomPassword(),
	}
}

func GenerateRandomAdmin() models.Admin{
	return models.Admin{
		Personne: generateRandomPersonne(),
		DateCreationCompte: time.Now().Add(time.Duration(rand.Intn(10000)) * time.Hour),
	}
}
func GenerateRandomAvocat() models.Avocat{
	return models.Avocat{
		Personne: generateRandomPersonne(),
		Cabinet: "Cabinet "+generateRandomNom(),
		NumeroBarreau: "NB"+strconv.Itoa(rand.Intn(9999)),
		Specialite: models.SpecialiteAvocat(rand.Intn(6)),

	}
}

func GenerateRandomClient() models.Client{
	return models.Client{
		Personne: generateRandomPersonne(),
		SituationJuridique: models.SituationJuridique(rand.Intn(4)),
		Profession: generateRandomProfession(),
	}
}

func GenerateRandomDossier() models.Dossier{
	return models.Dossier{
		Titre:"Titre De Dossier"+strconv.Itoa(rand.Intn(300)),
		Description:"Description De Dossier"+strconv.Itoa(rand.Intn(300)),
		DateCreation:time.Now().Add(time.Duration(rand.Intn(10000)) * time.Hour),
	}
}

func GenerateRandomDocument() models.Document{
	return models.Document{
		Nom:"Nom document"+strconv.Itoa(rand.Intn(300)),
		DateCreation: time.Now().Add(time.Duration(rand.Intn(10000)) * time.Hour),
		TypeFichier: models.TypeFichier(rand.Intn(6)),
		TypeDocument: models.TypeDocument(rand.Intn(5)),
		Contenu: generateRandomBytes(),
	}
}