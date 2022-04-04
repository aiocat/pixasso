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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Simple error struct to send as JSON
type Error struct {
	Message string
}

// Captcha result struct
type HCaptchaResult struct {
	Success bool `json:"success"`
}

// Simple function to check if hcaptcha token is valid
func captchaChecker(resp string) bool {
	if resp == "" {
		return false
	}

	request, err := http.NewRequest("POST", "https://hcaptcha.com/siteverify", bytes.NewBuffer([]byte("response="+resp+"&secret="+HCAPTCHA_SECRET)))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	result := new(HCaptchaResult)
	err = json.Unmarshal(body, &result)

	if err != nil {
		panic(err)
	}

	return result.Success
}

// Simple function to check if a string is only contains alpha characters
func alphaOnly(str string) bool {
	for _, char := range str {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyz_", strings.ToLower(string(char))) {
			return false
		}
	}
	return true
}
