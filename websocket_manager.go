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
	"time"

	"github.com/gofiber/websocket/v2"
)

var CONNECTIONS []*Socket // List all avaible sock connections

// Socket struct
type Socket struct {
	Connection *websocket.Conn `json:"-"`
	Pinged     bool            `json:"-"`
	Deleted    bool            `json:"-"`
	Token      string          `json:"id"`
	CreatedAt  time.Time       `json:"created_at"`
}

// WebsocketMessage struct
type WebSocketMessage struct {
	Type  uint8  `json:"type,omitempty" bson:"-"`
	X     uint   `json:"x" bson:"x"`
	Y     uint   `json:"y" bson:"y"`
	Color uint8  `json:"color" bson:"color"`
	From  string `json:"from" bson:"-"`
}

// Write a message to sock instance
func (s *Socket) WriteMessage(messageType int, message []byte) error {
	return s.Connection.WriteMessage(messageType, message)
}

// Remove sock from connections list
func (s *Socket) Destroy() {
	for index, sock := range CONNECTIONS {
		if sock.Token == s.Token {
			CONNECTIONS[index] = CONNECTIONS[len(CONNECTIONS)-1]
			CONNECTIONS = CONNECTIONS[:len(CONNECTIONS)-1]

			return
		}
	}
}

// Start checker (if sock is alive)
func (s *Socket) StartPingChecker() {
	go func() {
		for {
			time.Sleep(time.Duration(30) * time.Second)

			if s.Pinged {
				s.Pinged = false
			} else {
				s.Deleted = true
				return
			}
		}
	}()
}
