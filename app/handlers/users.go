package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	appCtx "github.com/HichuYamichu/go-webserver-reference/app/context"
	"github.com/HichuYamichu/go-webserver-reference/app/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUsers : Returns JSON with all users
func GetUsers(reqCtx *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
	r.Header.Set("content-type", "application/json")

	var users []models.User
	filter := bson.M{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := reqCtx.DB.Collection("users")
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	return nil
}

// InsertUser : Inserts new user to database
func InsertUser(reqCtx *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := reqCtx.DB.Collection("users")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}

	w.WriteHeader(200)
	w.Write([]byte("OK"))
	return nil
}

// UpdateUser : Updates user info
func UpdateUser(reqCtx *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
	params := mux.Vars(r)
	id := params["id"]
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll := reqCtx.DB.Collection("users")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	filter := bson.M{"_id": _id}
	update := bson.M{
		"$set": user,
	}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
	return nil
}

// DeleteUser : Removes user from database
func DeleteUser(reqCtx *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
	params := mux.Vars(r)
	id := params["id"]
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": _id}
	coll := reqCtx.DB.Collection("users")
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return &AppError{
			Err:  err,
			Msg:  http.StatusText(http.StatusInternalServerError),
			Code: http.StatusInternalServerError}
	}
	w.WriteHeader(200)
	w.Write([]byte("OK"))
	return nil
}
