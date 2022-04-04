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

function drawPixel(context, x, y, color) {
    context.fillStyle = color || '#000';
  	context.fillRect(x, y, 5, 5);
}


var canvas = document.getElementById("pixels")
var ctx = canvas.getContext("2d")

canvas.addEventListener("mousedown", (e) => {
    const rect = canvas.getBoundingClientRect()
    const x = Math.round((e.clientX - rect.left) / 5) * 5;
    const y = Math.round((e.clientY - rect.top) / 5) * 5;

    drawPixel(ctx, x, y, "#000");
    console.log("x: " + x + " y: " + y);
})