package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Title       string    `json:"title" bson:"title,omitempty"`
	Description string    `json:"description" bson:"description,omitempty"`
	Done        bool      `json:"done" bson:"done,omitempty"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

var client *mongo.Client

func New(mongoClient *mongo.Client) *Todo {
	client = mongoClient
	return &Todo{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("todos-db").Collection(collection)
}

func (t *Todo) InsertTodo(entry Todo) error {
	collection := returnCollectionPointer("todos")
	_, err := collection.InsertOne(context.TODO(), Todo{
		Title:       entry.Title,
		Description: entry.Description,
		Done:        entry.Done,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		log.Printf("Error inserting todo: %v", err)
		return err
	}

	return nil
}

func (t *Todo) GetTodos() ([]Todo, error) {
	collection := returnCollectionPointer("todos")
	var todos []Todo

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Printf("Error finding todos: %v", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *Todo) GetTodoByID(id string) (*Todo, error) {
	collection := returnCollectionPointer("todos")
	var todo Todo

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &Todo{}, err
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(&todo)
	if err != nil {
		log.Printf("Error finding todo by ID: %v", err)
		return &Todo{}, err
	}

	return &todo, nil
}

func (t *Todo) UpdateTodo(id string, updatedTodo Todo) error {
	collection := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	updatedTodo.UpdatedAt = time.Now()

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": mongoID}, bson.M{"$set": updatedTodo})
	if err != nil {
		log.Printf("Error updating todo: %v", err)
		return err
	}

	return nil
}

func (t *Todo) DeleteTodo(id string) error {
	collection := returnCollectionPointer("todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": mongoID})
	if err != nil {
		log.Printf("Error deleting todo: %v", err)
		return err
	}

	return nil
}
