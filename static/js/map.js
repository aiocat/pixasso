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

// Canvas section
const drawPixel = (context, x, y, color) => {
    if (color == 0) {
        context.fillStyle = "#000"
    }
  	context.fillRect(x, y, 5, 5);
}

var canvas = document.getElementById("pixels")
var ctx = canvas.getContext("2d")

canvas.addEventListener("mousedown", (e) => {
    const rect = canvas.getBoundingClientRect()
    const x = Math.round((e.clientX - rect.left) / 5) * 5;
    const y = Math.round((e.clientY - rect.top) / 5) * 5;

    websocket.send(JSON.stringify({
        "type": 1,
        "x": x,
        "y": y,
        "color": 0
    }))
})

// Connect to websocket
var websocket = new WebSocket(`ws://127.0.0.1:3000/ws/${localStorage.getItem("token")}`)

websocket.onclose = () => {
    console.log("Closed")
}

websocket.onopen = () => {
    setInterval(() => {
        websocket.send(JSON.stringify({
            "type": 0,
            "from": localStorage.getItem("token")
        }))
    }, 20000)
}

websocket.onmessage = ({ data }) => {
    let dataAsJson = JSON.parse(data)

    if (dataAsJson.type == 1) {
        drawPixel(ctx, dataAsJson.x, dataAsJson.y, dataAsJson.color)
    } else if (dataAsJson.error) {
        console.log(dataAsJson)
    }
}

// Load pixels
(async() => {
    let response = await fetch("/api/pixels")
    let respJson = await response.json()

    respJson.forEach(data => drawPixel(ctx, data.x, data.y, data.color))
})()