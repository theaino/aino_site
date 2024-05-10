function showError(signupError, message) {
  signupError.style.display = "block";
  signupError.innerText = message;
}

function requestSignup(name, email, password, signupError) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/signup";
  fetch(url, {
    method: "POST",
    body: JSON.stringify({
      name: name,
      email: email,
      password: password
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => {
      if (response.status === 200) {
        window.location.replace(host + "/home")
        return;
      }
      response.json().then((data) => {
        let code = data.code;
        switch (code) {
          case 1:
            showError(signupError, "You have entered an invalid email");
            break;
          case 2:
            showError(signupError, "A user with this name already exists");
            break;
          case 3:
            showError(signupError, "A user with this email already exists");
            break;
          default:
            showError(signupError, "An unexpected error occured")
        }
      });
    });
}

const emailRegex = /^(([^<>()[\]\.,;:\s@\"]+(\.[^<>()[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;

document.addEventListener("DOMContentLoaded", function () {

  let name = document.querySelector("#name");
  let email = document.querySelector("#email");
  let password = document.querySelector("#password");
  let signup = document.querySelector("#signup");
  let signupError = document.querySelector("#signup-error");
  name.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
      event.preventDefault();
      email.focus();
    }
  });
  email.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
      event.preventDefault();
      password.focus();
    }
  });
  password.addEventListener("keypress", function (event) {
    if (event.key === "Enter") {
      event.preventDefault();
      signup.click()
    }
  });

  signup.onclick = function () {
    signupError.display = "none";
    if (name.value === "" || email.value === "" || password.value === "") {
      return;
    }
    if (!email.value.match(emailRegex)) {
      showError(signupError, "You have entered an invalid email");
      return;
    }
    requestSignup(name.value, email.value, password.value, signupError);
  };

});

