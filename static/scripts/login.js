function showError(loginError, message) {
  loginError.style.display = "block";
  loginError.innerText = message;
}

function requestLogin(email, password, loginError) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/login";
  fetch(url + "?" + new URLSearchParams({
    "email": email,
    "password": password
  }), {
    method: "GET",
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => response.json())
    .then((data) => {
      if (data.authed) {
        window.location.replace(host + "/home");
      } else {
        showError(loginError, "You have entered an invalid email or password");
      }
    });
}

const emailRegex = /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;

document.addEventListener("DOMContentLoaded", function () {

  let email = document.querySelector("#email");
  let password = document.querySelector("#password");
  let login = document.querySelector("#login")
  let loginError = document.querySelector("#login-error");
  email.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
      event.preventDefault();
      password.focus();
    }
  });
  password.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
      event.preventDefault();
      login.click();
    }
  });

  login.onclick = function () {
    loginError.display = "none";
    if (email.value === "" || password.value === "") {
      return;
    }
    if (!email.value.match(emailRegex)) {
      showError(loginError, "You have entered an invalid email");
      return;
    }
    requestLogin(email.value, password.value, loginError);
  };

});

