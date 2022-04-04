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
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

// Handle create user route
func HandlePostUser(c *fiber.Ctx) error {
	// Check content type
	if c.Get("content-type", "") != "application/json" {
		return c.Status(400).JSON(Error{"content must be application/json"})
	}

	// Parse body
	user := User{}
	err := json.Unmarshal(c.Body(), &user)

	if err != nil {
		return c.Status(400).JSON(Error{"invalid json type"})
	}

	// Create account
	err = CreateAccount(user)
	if err != nil {
		return c.Status(400).JSON(Error{err.Error()})
	}

	return c.SendStatus(204)
}
