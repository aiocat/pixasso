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
	"crypto/sha1"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// User struct
type User struct {
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Token       string `json:"-" bson:"token"`
	Captcha     string `json:"captcha" bson:"-"`
	LastPixelAt int64  `json:"-" bson:"last_pixel_at"`
}

// Create new user account
func CreateAccount(user *User) error {
	// Check username and password
	if len((*user).Username) < 3 || len((*user).Username) > 24 || !alphaOnly((*user).Username) {
		return errors.New("bad username format")
	} else if len((*user).Password) < 8 || len((*user).Password) > 72 {
		return errors.New("bad password format")
	} else if !captchaChecker((*user).Captcha) {
		return errors.New("invalid captcha")
	}

	// Check username exists
	users := DATABASE.GetCollection("users")

	if users.FindOne(context.Background(), bson.M{"username": (*user).Username}).Err() == nil {
		return errors.New("username already exists")
	}

	// Prepare user
	(*user).LastPixelAt = 0
	(*user).Token = fmt.Sprintf("%x", sha1.Sum([]byte(user.Username+user.Password+user.Captcha))) // Set user token
	(*user).Password = fmt.Sprintf("%x", sha1.Sum([]byte(user.Password)))                         // Encrpyt user password

	// Insert user
	_, err := users.InsertOne(context.Background(), *user)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Auth an account
func AuthAccount(user *User) (string, error) {
	// Check username and password
	if len((*user).Username) < 3 || len((*user).Username) > 24 || !alphaOnly((*user).Username) {
		return "", errors.New("bad username format")
	} else if len((*user).Password) < 8 || len((*user).Password) > 72 {
		return "", errors.New("bad password format")
	}

	// Get user
	users := DATABASE.GetCollection("users")
	user.Password = fmt.Sprintf("%x", sha1.Sum([]byte((*user).Password)))
	result := users.FindOne(context.Background(), bson.M{"username": (*user).Username, "password": (*user).Password})

	if result.Err() != nil {
		return "", errors.New("user not found")
	}

	result.Decode(user)
	return (*user).Token, nil
}
