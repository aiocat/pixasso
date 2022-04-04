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
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Websocket handler
func WebSocket(c *websocket.Conn) {
	// c.Locals is added to the *websocket.Conn
	if !c.Locals("allowed").(bool) {
		c.Close()
		return
	}

	givenToken := c.Params("token", "")
	users := DATABASE.GetCollection("users")

	if givenToken == "" || users.FindOne(context.Background(), bson.M{"token": givenToken}).Err() != nil {
		c.Close()
		return
	}

	// Set read limit
	c.SetReadLimit(int64(1024))

	// New sock instance
	sock := Socket{
		Connection: c,
		Pinged:     false,
		Deleted:    false,
		Token:      givenToken,
		CreatedAt:  time.Now(),
	}
	sock.StartPingChecker()

	// Close sock instance end of the function
	defer sock.Destroy()

	// Add to connections
	CONNECTIONS = append(CONNECTIONS, &sock)

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		messageType int
		message     []byte
		wsError     error
	)
	for {
		// Read message
		if messageType, message, wsError = c.ReadMessage(); wsError != nil {
			log.Println("Read error:", wsError)
			break
		}

		// Decode message
		wsMessage := WebSocketMessage{}
		err := json.Unmarshal(message, &wsMessage)

		if err != nil {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"Invalid json format\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		}

		// Check if a ping message
		if wsMessage.Type == 0 {
			sock.Pinged = true
			continue
		} else if sock.Deleted { // Remove connection if sock instance deleted
			sock.Destroy()
			break
		}

		// Decode user
		user := User{}
		foundUser := users.FindOne(context.Background(), bson.M{"token": sock.Token})
		foundUser.Decode(&user)

		// Check if is in cooldown
		if time.Now().Unix()-user.LastPixelAt < 30 {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"You are in cooldown\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		}

		// Check if cords are true
		if wsMessage.X%5 != 0 || wsMessage.Y%5 != 0 || wsMessage.X > 1000 || wsMessage.Y > 1000 {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"Invalid cords\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		} else if wsMessage.Color > 12 {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"Invalid color\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		}

		// Prepare websocket message
		wsMessage.From = sock.Token

		// Insert into database
		pixels := DATABASE.GetCollection("pixels")
		_, err = pixels.ReplaceOne(context.Background(), bson.M{"x": wsMessage.X, "y": wsMessage.Y}, wsMessage, options.Replace().SetUpsert(true))

		if err != nil {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"Database error\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		}

		users.FindOneAndUpdate(context.Background(), bson.M{"token": sock.Token}, bson.M{"$set": bson.M{"last_pixel_at": time.Now().Unix()}})

		marshalledMessage, err := json.Marshal(wsMessage)

		if err != nil {
			// Write error message
			if wsError = sock.WriteMessage(messageType, []byte("{\"error\":\"Unknown error\"}")); wsError != nil {
				log.Println("Write error:", wsError)
				break
			}
			continue
		}

		for _, instance := range CONNECTIONS {
			if instance.Deleted {
				instance.Destroy()
				continue
			}

			if wsError = instance.WriteMessage(messageType, marshalledMessage); wsError != nil {
				log.Println("Write error:", wsError)
			}
		}
	}
}
