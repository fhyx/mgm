package mgm_test

import (
	"testing"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/internal/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDefConnection() {
	util.PanicErr(
		mgm.SetDefaultConfig(nil, "models", options.Client().ApplyURI("mongodb://root:12345@localhost:27017")),
	)
}

func resetCollection() {

	ctx, cancel := mgm.Ctx()
	defer cancel()

	_, err := mgm.Coll(&Doc{}).DeleteMany(ctx, bson.M{})
	_, err2 := mgm.Coll(&Person{}).DeleteMany(ctx, bson.M{})

	util.PanicErr(err)
	util.PanicErr(err2)
}

func seed() {
	docs := []any{
		NewDoc("Ali", 24),
		NewDoc("Mehran", 24),
		NewDoc("Reza", 26),
		NewDoc("Omid", 27),
	}

	ctx, cancel := mgm.Ctx()
	defer cancel()

	_, err := mgm.Coll(&Doc{}).InsertMany(ctx, docs)

	util.PanicErr(err)
}

func findDoc(t *testing.T) *Doc {
	found := &Doc{}

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, mgm.Coll(found).FindOne(ctx, bson.M{}).Decode(found))

	return found
}

type Doc struct {
	mgm.DefaultModel `bson:",inline"`

	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func NewDoc(name string, age int) *Doc {
	return &Doc{Name: name, Age: age}
}
