package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Vzaldat/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//this is connection clause, enter your own credentials like password and check the connection type
//The following section of lines(till the next comment) is just an example of the connection strings
const connectionString = "mongodb+srv://vishaldatta2002:vzal123@cluster1.xxpakzp.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// to call the reference of the collections we have
var collection *mongo.Collection

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connection to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption) // when context type you are unsure of what to do, background operation

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success!!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready to execute")

}

// MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}
func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println("modified count: ", res.ModifiedCount)
}

func deleteOneMovie(movieID string) int64 {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	deleteCount, err1 := collection.DeleteOne(context.Background(), filter)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println("Movie got deleted with delete count : ", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all the movies

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err1 := cursor.Decode(&movie); err1 != nil {
			log.Fatal(err1)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//Actual controller - file

func GetAllMyMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func Deleteam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	var cnt int64 = deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(cnt)
}

func Deleteallms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

//PUT CAPITALS FOR ALL EXPORT FUNCTIONS
