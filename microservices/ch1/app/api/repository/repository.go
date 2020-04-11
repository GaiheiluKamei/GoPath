package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type TZConversion struct {
	TimeZone       string `bson:"timeZone" json:"timeZone"`
	TimeDifference string `bson:"timeDifference" json:"timeDifference"`
}

type Repository struct {
	dbSession    *mgo.Session
	dbServer     string
	dbDatabase   string
	dbCollection string
}

func NewRepository(dbServer, dbDatabase, dbCollection string) *Repository {
	dbSession, err := mgo.Dial(dbServer)
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		dbSession: dbSession,
		dbServer: dbServer,
		dbDatabase: dbDatabase,
		dbCollection: dbCollection,
	}
}

func (repo *Repository) Close() {
	repo.dbSession.Close()
}

func (repo *Repository) newCollection() *mgo.Collection {
	dbSession := repo.dbSession.Clone()
	return dbSession.DB(repo.dbDatabase).C(repo.dbCollection)
}

func (repo *Repository) FindAll() ([]TZConversion, error) {
	var tzcs []TZConversion
	return tzcs, repo.newCollection().Find(bson.M{}).All(&tzcs)
}

func (repo *Repository) FindByTimeZone(tz string) (TZConversion, error) {
	var tzc TZConversion
	return tzc, repo.newCollection().Find(bson.M{"timeZone": tz}).One(&tzc)
}

func (repo *Repository) Insert(tzc TZConversion) error {
	return repo.newCollection().Insert(&tzc)
}

func (repo *Repository) Delete(tzc TZConversion) error {
	return repo.newCollection().Remove(bson.M{"timeZone": tzc.TimeZone})
}

func (repo *Repository) Update(tz string, tzc TZConversion) error {
	return repo.newCollection().Update(bson.M{"timeZone": tz}, &tzc)
}