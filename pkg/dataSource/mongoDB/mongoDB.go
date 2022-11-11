package mongoDB

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "goskurikhin/pkg/dataSource"
)

type MongoDB struct {
	client *mongo.Client
}

func NewMongoDB() *MongoDB {
	return &MongoDB{
		client: nil,
	}
}

func (m *MongoDB) Open(dsn string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		return err
	}
	m.client = client
	return nil
}

func (m *MongoDB) Close() {
	m.client.Disconnect(context.TODO())
}

func (m *MongoDB) InsertLinks(longurl, shortname string) error {
	coll := m.client.Database("urlsdb").Collection("urls")

	doc := bson.D{{"longurl", longurl}, {"shortname", shortname}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDB) FindLongLink(shortname string) (string, error) {
	var doc bson.M
	coll := m.client.Database("urlsdb").Collection("urls")
	filter := bson.D{{"shortname", shortname}}
	err := coll.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		return "", err
	}
	return doc["longurl"].(string), nil
}

func (m *MongoDB) DoesExistShortname(shortname string) (bool, error) {
	var doc bson.M
	coll := m.client.Database("urlsdb").Collection("urls")
	filter := bson.D{{"shortname", shortname}}
	err := coll.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return true, err
	}
	return true, nil
}
