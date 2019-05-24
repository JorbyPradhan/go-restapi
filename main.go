package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Person struct {
	ID 		primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string 		   `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string 		   `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client
func CreatePersonEndPoint(response http.ResponseWriter, request *http.Request){
	response.Header().Add("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ :=context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func main()  {
	fmt.Println("Starting")
	ctx, _ := context.WithTimeout(context.Background(),10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	router :=mux.NewRouter()
	router.HandleFunc("/person",CreatePersonEndPoint).Methods("POST")
	http.ListenAndServe(":12345", router)
}
/*import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

type Person struct{
	ID string `json:"id"`
	Name string `json:"name"`
	NRC string `json:"nrc"`
}
type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var person []Person
var posts []Post

func getPerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func createPersonData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var pers Person
	_= json.NewDecoder(r.Body).Decode(&pers)
	pers.ID = strconv.Itoa(rand.Intn(100))
	person = append(person,pers)
	json.NewEncoder(w).Encode(&pers)

}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	json.NewEncoder(w).Encode(&post)
}
func getPersonDetail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range person{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}
func updatePerson(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range person{
		if item.ID == params["id"]{
			person = append(person[:index], person[index+1:]...)
			var pers Person
			_= json.NewDecoder(r.Body).Decode(&pers)
			pers.ID = params["id"]
			person= append(person,pers)
			json.NewEncoder(w).Encode(&pers)
			return
		}
	}
	json.NewEncoder(w).Encode(person)
}
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {

			// using slice
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}
func deletePerson(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range person{
		if item.ID == params["id"]{
			person = append(person[:index], person[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(person)

}
func main() {
	router := mux.NewRouter()
	//slice
	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
*/