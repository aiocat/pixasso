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
var captcha = document.getElementById("captcha");

document.getElementById("create").onclick = () => {
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
    
    captcha.execute();
};

captcha.addEventListener("verified", async ({ token }) => {
    let response = await fetch("/users", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            captcha: token,
            username: username.value,
            password: password.value,
        }),
    });

    if (response.ok) {
        window.location.href = "/signin";
    } else {
        let responseJson = await response.json();
        document.getElementById("error").innerText = responseJson.error;
    }
});

captcha.addEventListener("error", ({ error }) => {
    document.getElementById("error").innerText = error;
});