package db

import (
	"github.com/kosuda/golang-web/config"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"strconv"
)

var _session *mgo.Session
var dbname string

func init() {
	log.Print("mongo db init")
	c := config.Configuration.Mongo

	session, err := mgo.Dial("mongodb://" + c.Host + ":" + strconv.Itoa(c.Port))

	if err != nil {
		log.Fatal("init error:", err.Error())
	}

	_session = session
	dbname = c.Dbname
}

// Find all or by selector function
func Find(colname string, query interface{}, result interface{}) error {
	col := _session.DB(dbname).C(colname)

	err := col.Find(query).All(result)

	if err != nil {
		return err
	}

	return nil
}

// FindByID function (result must be one)
func FindByID(colname string, id string, result interface{}) error {
	col := _session.DB(dbname).C(colname)

	err := col.Find(bson.M{"id": id}).One(result)

	if err != nil {
		return err
	}

	return nil
}

// Upsert function
func Upsert(colname string, id string, data interface{}) error {
	col := _session.DB(dbname).C(colname)

	if id == "" {
		id = bson.NewObjectId().Hex()
		log.Printf("create _id: %s", id)
	}

	if _, err := col.Upsert(map[string]string{"_id": id}, data); err != nil {
		return err
	}

	return nil
}

// Update function
func Update(colname string, query interface{}, data interface{}) error {
	col := _session.DB(dbname).C(colname)

	if err := col.Update(query, bson.M{"$set": data}); err != nil {
		return err
	}

	return nil
}

// UpdateByID function
func UpdateByID(colname string, id string, data interface{}) error {
	col := _session.DB(dbname).C(colname)

	if err := col.UpdateId(id, bson.M{"$set": data}); err != nil {
		return err
	}

	return nil
}

// Remove all function
func Remove(colname string, query interface{}) error {
	col := _session.DB(dbname).C(colname)

	if err := col.Remove(query); err != nil {
		return err
	}

	return nil
}

// RemoveByID function
func RemoveByID(colname string, id string) error {
	col := _session.DB(dbname).C(colname)

	if err := col.RemoveId(id); err != nil {
		return err
	}

	return nil
}
