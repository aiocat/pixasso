<!--
 Copyright (C) 2022 aiocat
 
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as
 published by the Free Software Foundation, either version 3 of the
 License, or (at your option) any later version.
 
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.
 
 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
-->

<div align="center">

![Logo](/static/img/pixasso.png)
# Pixasso
Anarchic pixel art site with only one rule.

</div>

## Rule
You can insert only one pixel in 2 second.

## Technologies

- **Programming Language**: Go
- **Database**: MongoDB
- **Server**: Gofiber
- **Captcha Service**: HCaptcha
- **Front-end**: HTML, CSS & JS
- **Hosting**: Heroku

## Hosting

Create a `.env` file and add these key-values:

- **MONGO_URL**: MongoDB database connection url.
- **HCAPTCHA_SECRET**: HCaptcha secret key.
- **PORT**: Port to serve.

Install Go programming language and run with: `go run .`

View demo here: https://pixasso-app.herokuapp.com

## Routes

- `GET /`
- `GET /signin`
- `GET /signup`
- `POST /api/users`
- `POST /api/users/auth`
- `GET /api/pixels`
- `WEBSOCKET /ws/:token`

## License
Pixasso is distributed under AGPLv3 license. for more information:

- https://raw.githubusercontent.com/aiocat/pixasso/main/LICENSE