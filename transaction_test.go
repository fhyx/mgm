package mgm_test

import (
	"testing"

	"github.com/kamva/mgm/v3"
	"github.com/kamva/mgm/v3/internal/util"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Note: to run Transaction tests, the MongoDB daemon must run as replica set, not as a standalone daemon.
// To convert it [see this](https://docs.mongodb.com/manual/tutorial/convert-standalone-to-replica-set/)
func TestTransactionCommit(t *testing.T) {
	setupDefConnection()
	resetCollection()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	})

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(ctx, bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(1), count)
}

func TestTransactionAbort(t *testing.T) {
	setupDefConnection()
	resetCollection()
	//seed()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	ctx, cancel := mgm.Ctx()
	defer cancel()

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(ctx, bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}

func TestTransactionWithCtx(t *testing.T) {
	setupDefConnection()
	resetCollection()
	//seed()

	ctx, cancel := mgm.Ctx()
	defer cancel()

	d := &Doc{Name: "check", Age: 10}

	err := mgm.TransactionWithCtx(ctx, func(session mongo.Session, sc mongo.SessionContext) error {

		err := mgm.Coll(d).CreateWithCtx(sc, d)

		if err != nil {
			return err
		}

		return session.AbortTransaction(sc)
	})

	util.AssertErrIsNil(t, err)
	count, err := mgm.Coll(d).CountDocuments(ctx, bson.M{})

	util.AssertErrIsNil(t, err)
	require.Equal(t, int64(0), count)
}
