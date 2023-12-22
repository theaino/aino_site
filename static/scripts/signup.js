function requestSignup(name, email, password) {
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
        window.location.replace(host + "/login")
      }
    });
}

document.addEventListener("DOMContentLoaded", function () {

  let name = document.querySelector("#name");
  let email = document.querySelector("#email");
  let password = document.querySelector("#password");
  let signup = document.querySelector("#signup")

  signup.onclick = function () {
    requestSignup(name.value, email.value, password.value);
  };

});

