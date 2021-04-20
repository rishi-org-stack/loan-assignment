package model

import (
	"context"
	"fmt"
	"os"

	// "log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson" 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	uri        string
	client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func Instantiate() *DB {
	godotenv.Load(".env")
	Db := &DB{}
	Db.uri = os.Getenv("MONGODB_URL")
	return Db
}

func (db *DB) Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.uri))
	if err != nil {
		fmt.Printf("error in connection : %v\n", err)
	}
	db.client = client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db.client.Connect(ctx)
	defer cancel()
	return db
}

func (db *DB) CreateDb(databse string ) *DB{
	db.Database = db.client.Database(databse)
	return db
	// db.Collection = db.Database.Collection(databse)
}
func (db *DB) LinkToCollection(coll string) *DB{
	db.Collection = db.Database.Collection(coll)
	return db
}
func (db *DB) Insert(p interface{}) interface{} {
	coll := db.Collection
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	i, e := coll.InsertOne(ctx, &p)
	fmt.Println(e)
	return i.InsertedID
}

func (db *DB) Get(params string, val interface{}, p interface{}) error {
	coll := db.Collection
	filter := bson.D{{params, val}}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	r := coll.FindOne(ctx, filter).Decode(p)
	if r != nil {
		return r
	}
	defer cancel()
	return nil
}

func (db *DB) UpdateaDocument(find map[string]interface{}, set map[string]interface{}) (count int64, err error) {
	coll := db.Collection
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	for i := range set {
		res, erro := coll.UpdateOne(ctx, bson.M{(find["identity"]).(string): find["val"]}, bson.D{
			{
				"$set", bson.D{{i, set[i]}},
			},
		})
		count = res.ModifiedCount
		err = erro
	}
	return
}
func (db *DB) AddaDocument(find map[string]interface{}, set map[string]interface{}) ( err error) {
	coll := db.Collection
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	for i := range set {
		_, erro := coll.UpdateOne(ctx, bson.M{(find["identity"]).(string): find["val"]}, bson.D{
			{"$push", bson.D{{i, set[i]}},},
		})
		// count = res.ModifiedCount
		err = erro
	}
	return
}
func (db *DB) DeleteADocument(find map[string]interface{})(count int64,err error) {
	coll := db.Collection

	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	res,erro:=coll.DeleteOne(ctx, bson.M{find["identity"].(string): find["val"]})
	count =res.DeletedCount
	err  =erro
	return
}

