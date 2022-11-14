package argStore

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mordor struct {
	host       string
	user       string
	pass       string
	dbName     string
	table      string
	collection *mongo.Collection
}

func NewMordor(host string, user string, pass string, db string, table string) *mordor {
	return &mordor{
		host:   host,
		user:   user,
		pass:   pass,
		dbName: db,
		table:  table,
	}
}

func (m *mordor) Get(ctx context.Context, field string, value interface{}, decoder interface{}) error {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return errors.New("Connection not setup")
	}

	if "key" == field {
		field = "_id"
	}

	if "_id" == field {
		id, idErr := primitive.ObjectIDFromHex(value.(string))

		if nil != idErr {
			return errors.New("Invalid key")
		}

		value = id
	}

	err := m.collection.FindOne(ctx, bson.M{field: value}).Decode(decoder)

	return err
}

func (m *mordor) GetMany(ctx context.Context, field string, value interface{}, limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error) {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return 0, errors.New("Connection not setup")
	}

	count, cErr := m.collection.EstimatedDocumentCount(ctx, nil)

	if nil != cErr {
		return 0, cErr
	}

	findOpts := options.Find()
	findOpts.SetLimit(limit)
	findOpts.SetSkip(offset)
	findOpts.SetSort(sort)
	cursor, err := m.collection.Find(ctx, bson.M{field: value}, findOpts)

	if nil != err {
		return 0, err
	}

	cursor.All(ctx, decoder)

	return count, nil
}

func (m *mordor) GetAll(ctx context.Context, limit int64, offset int64, sort interface{}, decoder interface{}) (int64, error) {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return 0, errors.New("Connection not setup")
	}

	count, cErr := m.collection.EstimatedDocumentCount(ctx, nil)

	if nil != cErr {
		return 0, cErr
	}

	findOpts := options.Find()
	findOpts.SetLimit(limit)
	findOpts.SetSkip(offset)
	findOpts.SetSort(sort)
	cursor, err := m.collection.Find(ctx, bson.D{}, findOpts)

	if nil != err {
		return 0, err
	}

	cursor.All(ctx, decoder)

	return count, nil
}

func (m *mordor) Write(ctx context.Context, data interface{}) (string, error) {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return "", errors.New("Connection not setup. Use mordor.Setup needs called first")
	}

	result, err := m.collection.InsertOne(ctx, data)

	if nil != err {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return "", errors.New("Unable to parse InsertID")
	}

	return id.Hex(), nil
}

func (m *mordor) Update(ctx context.Context, key string, newData interface{}) error {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return errors.New("Connection not setup. Use mordor.Setup needs called first")
	}

	id, idErr := primitive.ObjectIDFromHex(key)

	if nil != idErr {
		return errors.New("Invalid key")
	}

	_, err := m.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", newData},
		},
	)

	if nil != err {
		return err
	}

	return nil
}

func (m *mordor) Delete(ctx context.Context, key string) error {
	client, _ := m.init(ctx)
	defer client.Disconnect(ctx)

	if nil == m.collection {
		return errors.New("Connection not setup. Use mordor.Setup needs called first")
	}

	id, idErr := primitive.ObjectIDFromHex(key)

	if nil != idErr {
		return errors.New("Invalid key")
	}

	_, err := m.collection.DeleteOne(ctx, bson.M{"_id": id})

	if nil != err {
		return err
	}

	return nil
}

func (m *mordor) init(ctx context.Context) (*mongo.Client, error) {
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://" + m.user + ":" + m.pass + "@" + m.host + "/?authSource=admin&readPreference=primary&ssl=false"))
	// ctx, _ := context.WithTimeout(context.Background(), time.Second+10)
	clientErr := client.Connect(ctx)
	collection := client.Database(m.dbName).Collection(m.table)

	m.collection = collection

	return client, clientErr
}
