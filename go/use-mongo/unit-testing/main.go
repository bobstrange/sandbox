package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gookit/color.v1"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	opts := options.Client().ApplyURI("mongodb://user:password@localhost:37017")
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("tasker").Collection("tasks")
}

// Task task entity
type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func createTask(task *Task) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func getAll() ([]*Task, error) {
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error) {
	var tasks []*Task

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return tasks, err
	}

	for cur.Next(ctx) {
		var t Task
		if err := cur.Decode(&t); err != nil {
			return tasks, err
		}
		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil {
		return tasks, err
	}

	cur.Close(ctx)

	if len(tasks) == 0 {
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil
}

func printTasks(tasks []*Task) {
	for i, t := range tasks {
		if t.Completed {
			color.Green.Printf("%d: %s\n", i+1, t.Text)
		} else {
			color.Yellow.Printf("%d: %s\n", i+1, t.Text)
		}
	}
}

func addCommand() *cli.Command {
	return &cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task to the list",
		Action: func(c *cli.Context) error {
			text := c.Args().First()
			if text == "" {
				return errors.New("Cannot add an empty task")
			}

			task := &Task{
				ID:        primitive.NewObjectID(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Text:      text,
				Completed: false,
			}
			return createTask(task)
		},
	}
}

func allCommand() *cli.Command {
	return &cli.Command{
		Name:    "all",
		Aliases: []string{"l"},
		Usage:   "list all tasks",
		Action: func(c *cli.Context) error {
			tasks, err := getAll()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Print("Nothing to see here.\nRun `add task` to add a task")
					return nil
				}
				return err
			}
			printTasks(tasks)
			return nil
		},
	}
}

func main() {
	app := &cli.App{
		Name:  "tasker",
		Usage: "A simple CLI program to manage your tasks",
		Commands: []*cli.Command{
			addCommand(),
			allCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
