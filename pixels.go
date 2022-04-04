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

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Handle map pixel route
func HandlePixelApi(c *fiber.Ctx) error {
	pixels := DATABASE.GetCollection("pixels")
	cursor, err := pixels.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(Error{"Database error"})
	}

	defer cursor.Close(context.Background())
	var pixelsSlice []*WebSocketMessage

	// Iterate results
	for cursor.Next(context.Background()) {
		pixel := WebSocketMessage{}
		err := cursor.Decode(&pixel)

		if err != nil {
			return c.Status(500).JSON(Error{"Database error"})
		}

		pixelsSlice = append(pixelsSlice, &pixel)
	}

	return c.JSON(pixelsSlice)
}
