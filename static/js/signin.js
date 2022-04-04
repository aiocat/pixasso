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

var username = document.getElementById("username");
var password = document.getElementById("password");

document.getElementById("auth").onclick = async() => {
    if (username.value.length < 3) {
        document.getElementById("error").innerText = "Username is too short";
        return;
    } else if (username.value.length > 24) {
        document.getElementById("error").innerText = "Username is too long";
        return;
    } else if (password.value.length < 8) {
        document.getElementById("error").innerText = "Password is too short";
        return;
    } else if (password.value.length > 72) {
        document.getElementById("error").innerText = "Password is too long";
        return;
    }
    
    let response = await fetch("/api/users/auth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username: username.value,
            password: password.value,
        }),
    });

    let responseJson = await response.json();

    if (response.ok) {
        localStorage.setItem("token", responseJson.token)
        window.location.href = "/";
    } else {
        document.getElementById("error").innerText = responseJson.message;
    }
};