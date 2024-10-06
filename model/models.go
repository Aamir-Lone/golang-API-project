package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Connect to MongoDB
func ConnectDB() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://Aamirlone:Aamirlone@cluster0.1zwiz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// Define a model for Todo
type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty"`
	Completed bool               `json:"completed" bson:"completed"`
}

// Insert a new todo
func InsertTodo(todo Todo) (*mongo.InsertOneResult, error) {
	collection := client.Database("tododb").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, todo)
	return result, err
}

// Get all todos
func GetAllTodos() ([]Todo, error) {
	collection := client.Database("tododb").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []Todo
	for cursor.Next(ctx) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

// Get a single todo by ID
func GetTodoByID(id string) (Todo, error) {
	collection := client.Database("tododb").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Todo{}, err
	}

	var todo Todo
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	return todo, err
}

// Update a todo by ID
func UpdateTodoByID(id string, todo Todo) error {
	collection := client.Database("tododb").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":     todo.Title,
			"completed": todo.Completed,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// Delete a todo by ID
func DeleteTodoByID(id string) error {
	collection := client.Database("tododb").Collection("todos")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
