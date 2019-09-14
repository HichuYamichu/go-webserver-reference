package user

import (
	"log"
	"net/http"
)

// GetUsers : Returns JSON with all users
func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) (int, error) {
	log.Printf("main")
	w.Write([]byte("hello\ngo\n"))

	// err = json.NewEncoder(w).Encode(users)
	// if err != nil {
	// 	return &AppError{
	// 		Err:  err,
	// 		Msg:  http.StatusText(http.StatusInternalServerError),
	// 		Code: http.StatusInternalServerError}
	// }
	return 200, nil
}

// // InsertUser : Inserts new user to database
// func InsertUser(CTX *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
// 	var user models.User
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	coll := CTX.DB.Collection("users")
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	_, err = coll.InsertOne(ctx, user)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}

// 	w.WriteHeader(200)
// 	w.Write([]byte("OK"))
// 	return nil
// }

// // UpdateUser : Updates user info
// func UpdateUser(CTX *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	var user models.User
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	coll := CTX.DB.Collection("users")
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	filter := bson.M{"_id": _id}
// 	update := bson.M{
// 		"$set": user,
// 	}
// 	_, err = coll.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	w.WriteHeader(200)
// 	w.Write([]byte("OK"))
// 	return nil
// }

// // DeleteUser : Removes user from database
// func DeleteUser(CTX *appCtx.Context, w http.ResponseWriter, r *http.Request) *AppError {
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	_id, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	filter := bson.M{"_id": _id}
// 	coll := CTX.DB.Collection("users")
// 	_, err = coll.DeleteOne(ctx, filter)
// 	if err != nil {
// 		return &AppError{
// 			Err:  err,
// 			Msg:  http.StatusText(http.StatusInternalServerError),
// 			Code: http.StatusInternalServerError}
// 	}
// 	w.WriteHeader(200)
// 	w.Write([]byte("OK"))
// 	return nil
// }