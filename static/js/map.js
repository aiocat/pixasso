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
const drawPixel = (context, x, y, c) => {
    context.fillStyle = colors[c]
    context.fillRect(x, y, 10, 10);
}

var canvas = document.getElementById("pixels")
var ctx = canvas.getContext("2d")
var color = 0

canvas.addEventListener("mousedown", (e) => {
    if (e.button != 0) return;

    const rect = canvas.getBoundingClientRect()
    const x = (Math.ceil((e.clientX - rect.left) / 10) * 10) - 10;
    const y = (Math.ceil((e.clientY - rect.top) / 10) * 10) - 10;

    websocket.send(JSON.stringify({
        "type": 1,
        "x": x,
        "y": y,
        "color": color
    }))
})

canvas.addEventListener("mousemove", (e) => {
    const rect = canvas.getBoundingClientRect()
    const x = (Math.ceil((e.clientX - rect.left) / 10) * 10) - 10;
    const y = (Math.ceil((e.clientY - rect.top) / 10) * 10) - 10;

    document.getElementById("x-axis").innerHTML = `<strong>X:</strong> ${x}`
    document.getElementById("y-axis").innerHTML = `<strong>Y:</strong> ${y}`
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
(async () => {
    let response = await fetch("/api/pixels")
    let respJson = await response.json()

    if (!respJson) return
    respJson.forEach(data => drawPixel(ctx, data.x, data.y, data.color))
})()

// Load colors
// black, gray, white, brown, red, orange, yellow, green, cyan, blue, purple, pink
var colors = ["#000", "#777", "#FFF", "#964B00", "#FF0000", "#FFA500", "#FFFF00", "#00FF00", "#00FFFF", "#0000FF", "#800080", "#FFC0CB"]

colors.forEach((value, index) => {
    let elem = document.createElement("span")
    elem.style.background = value
    elem.onclick = () => {
        color = index
        document.getElementById("color-status").style.background = value
    }

    document.getElementById("palette").append(elem)
})

// Load pixel slices
// Y 
for (let i = 0; i < 10000; i += 10) {
    ctx.beginPath()
    ctx.lineWidth = 0.1
    ctx.moveTo(i, 0)
    ctx.lineTo(i, 10000)
    ctx.stroke()
}

// X
for (let i = 0; i < 10000; i += 10) {
    ctx.beginPath()
    ctx.lineWidth = 0.1
    ctx.moveTo(0, i)
    ctx.lineTo(10000, i)
    ctx.stroke()
}