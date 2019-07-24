package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/HichuYamichu/go-webserver-reference/app/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUsers : Returns JSON with all users
func GetUsers(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	r.Header.Set("content-type", "application/json")

	var users []models.User
	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := db.Collection("users")
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err = json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
}

// InsertUser : Inserts new user to database
func InsertUser(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := db.Collection("users")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(body, &user); err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// UpdateUser : Updates user info
func UpdateUser(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := db.Collection("users")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(body, &user); err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": user,
	}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// DeleteUser : Removes user from database
func DeleteUser(db *mongo.Database, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": _id}
	coll := db.Collection("users")
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
