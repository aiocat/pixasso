// Copyright (C) 2022 aiocat
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DATABASE = Database{"", "pixasso", nil} // Global variable to access database
)

// Public mongodb database struct
type Database struct {
	MongoUrl, DatabaseName string
	MongoClient            *mongo.Client
}

// Start database connection
func (db *Database) StartConnection(database_uri string) error {
	db.MongoUrl = database_uri
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.MongoUrl))

	if err != nil {
		return err
	}

	// Ping database
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return err
	}

	db.MongoClient = client
	return nil
}

// Get mongo collection from database
func (db *Database) GetCollection(collection string) *mongo.Collection {
	return db.MongoClient.Database(db.DatabaseName).Collection(collection)
}
