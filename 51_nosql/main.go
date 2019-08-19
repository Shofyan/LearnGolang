package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.Background()


type student struct {
	Name string `bson:name`
	Grade int `bson:grade`
}

func connect() (*mongo.Database, error){
	clientConnection := options.Client()
	clientConnection.ApplyURI("mongodb://localhost:27017")
	client , err := mongo.NewClient(clientConnection)
	if err != nil {
		return nil, err
	}

	return client.Database("LearnGolang"),nil
}


func insert(){
	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	_, err= db.Collection("student").InsertOne(ctx,student{"wick",23})
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Collection("student").InsertOne(ctx,student{"jhon",17})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("insert success!")
}


func find(){
	db, err := connect()
	if err != nil {
		log.Fatalln(err)
	}

	csr,err := db.Collection("student").Find(ctx,bson.M{"name":"jhon"})
	if err != nil {log.Fatal(err)}
	defer csr.Close(ctx)
	result := make([]student, 0)
	for csr.Next(ctx) {
		var row student
		err:= csr.Decode(&row)
		if err != nil{
			log.Fatalln(err)
		}
		result = append(result,row)
	}

	if len(result) > 0 {
		fmt.Println("name: ",result[0].Name)
		fmt.Println("grade: ",result[0].Grade)
	}
}

func update(){
	db , err := connect()
	if err != nil {
		log.Fatalln(err)
	}
	var selector = bson.M{"name":"jhon"}
	var changes = student{"jhon wick",2}
	_,err  = db.Collection("student").UpdateOne(ctx,selector, bson.M{"$set":changes})
	if err != nil { log.Fatalln(err)}
	 fmt.Println("update success")
}


func remove(){
	db , err := connect()
	if err != nil {
		log.Fatalln(err)
	}
	var selector = bson.M{"name":"jhon"}
	_,err  = db.Collection("student").DeleteOne(ctx,selector})
	if err != nil { log.Fatalln(err)}
	fmt.Println("delete success")
}


func main (){
insert()
find()
update()
remove()
}
