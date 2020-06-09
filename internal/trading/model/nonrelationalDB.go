package model

import (
	"gopkg.in/mgo.v2"
)

type mongoDB struct {
	session *mgo.Session
	db      *mgo.Database
}

func newMongoDB() (*mongoDB, error) {
	session, err := mgo.Dial("mongodb://root:123456@0.0.0.0:27017")
	if err != nil {
		return nil, err
	}
	// session.SetMode(mgo.Monotonic, true)

	db := session.DB("exchange")
	return &mongoDB{
		session: session,
		db:      db,
	}, nil
}

func (m *mongoDB) insert(collection string, data interface{}) error {
	err := m.db.C(collection).Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoDB) close() {
	m.session.Close()
}
