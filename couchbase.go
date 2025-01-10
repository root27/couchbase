package main

import (
	"github.com/couchbase/gocb/v2"
	"log"
	"os"
)

type Couchbase struct {
	*gocb.Cluster
}

var (
	connectionString = os.Getenv("CONNECTION_STRING")
	username         = os.Getenv("USER")
	password         = os.Getenv("PASSWORD")
	bucketName       = os.Getenv("BUCKET")
	collectionName   = os.Getenv("COLLECTION")
	scopeName        = os.Getenv("SCOPE")
)

func couchInit() *Couchbase {

	options := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	}

	if err := options.ApplyProfile(gocb.ClusterConfigProfileWanDevelopment); err != nil {
		log.Fatal(err)
	}

	// Initialize the Connection
	cluster, err := gocb.Connect(connectionString, options)
	if err != nil {
		log.Fatal(err)
	}

	return &Couchbase{
		cluster,
	}
}

// Insert Data

func (c *Couchbase) insertData(key string, data interface{}) error {
	collection := c.Bucket(bucketName).Scope(scopeName).Collection(collectionName)
	_, err := collection.Insert(key, data, nil)
	return err
}

// Get Data

func (c *Couchbase) getData(key string) (*Example, error) {
	collection := c.Bucket(bucketName).Scope(scopeName).Collection(collectionName)
	doc, err := collection.Get(key, nil)
	if err != nil {

		return nil, err

	}

	var exampleData Example

	if err := doc.Content(&exampleData); err != nil {
		return nil, err
	}

	return &exampleData, nil
}

// Update Data

func (c *Couchbase) updateData(key string, data interface{}) error {
	collection := c.Bucket(bucketName).Scope(scopeName).Collection(collectionName)
	_, err := collection.Replace(key, data, nil)
	return err
}

// Delete Data

func (c *Couchbase) deleteData(key string) error {
	collection := c.Bucket(bucketName).Scope(scopeName).Collection(collectionName)
	_, err := collection.Remove(key, nil)
	return err
}
