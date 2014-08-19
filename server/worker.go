package server

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type EstadoCivil uint8

const (
	Solteiro EstadoCivil = iota
	Casado
	Divorciado
	Viuvo
)

type CartaProfissional uint8

func WorkerCol(s *mgo.Session) *mgo.Collection {
	return s.DB("365").C("Workers")
}

const (
	Coordenador CartaProfissional = iota
	//assistente recinto desportivo
	ARD
	//assistente recinto espectaculos
	ARE
	VigilantePorteiro
	Vigilante
)

type Cartao struct {
	Anexo    bson.ObjectId `bson:",omitempty"`
	Validade time.Time
	Numero   int64
}

type Hierarquia uint8

const (
	DirectorSeguranca Hierarquia = iota
	CoordenadorSeguranca
	Supervisor
	ChefeDeGrupo
	Vigilantes
)

type Address struct {
	Rua          string
	CodigoPostal string
	Localidade   string
	Morada       string
}

type Worker struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	PrimeiroNome string
	UltimoNome   string
	Altura       string
	Peso         uint8
	Email        string
	Telefone     string
	Telemovel    string

	BI_CC          string //bi ou cartão de cidadão
	BI_CCExpiracao time.Time
	Nif            string
	NISS           string

	Nib         string
	EstadoCivil string

	Address Address

	Posto Hierarquia

	CartaProfissional map[CartaProfissional]Cartao

	AnexoComprovativoMorada bson.ObjectId `bson:",omitempty"`
	AnexoRegistoCriminal    bson.ObjectId `bson:",omitempty"`
	AnexoCV                 bson.ObjectId `bson:",omitempty"`
	AnexoNib                bson.ObjectId `bson:",omitempty"`
	AnexoNIF                bson.ObjectId `bson:",omitempty"`
	AnexoNISS               bson.ObjectId `bson:",omitempty"`
	AnexoDocumetoID         bson.ObjectId `bson:",omitempty"`
	AnexoDeclaracaoConduta  bson.ObjectId `bson:",omitempty"`

	Foto bson.ObjectId `bson:",omitempty"`

	Created time.Time
	Updated time.Time
}

func NewWorker() (w *Worker) {
	w = new(Worker)
	w.CartaProfissional = map[CartaProfissional]Cartao{}
	return w
}

func (w *Worker) Save(s *mgo.Session) error {
	//we are creating

	if w.Id == bson.ObjectId("") {
		w.Created = time.Now()
		w.Updated = time.Now()
	} else { //we are updating
		w.Updated = time.Now()
	}
	_, err := WorkerCol(s).Upsert(bson.M{"_id": w.Id}, w)
	if err != nil {
		return err
	}
	return nil
}

func GetWorkerById(s *mgo.Session, id bson.ObjectId) (error, *Worker) {
	w := new(Worker)
	err := WorkerCol(s).FindId(id).One(w)
	if err != nil {
		return err, nil
	}
	return nil, w
}

func (w *Worker) Delete(s *mgo.Session) {
	WorkerCol(s).RemoveId(w.Id)

}
