function showError(message, settingsError) {
  settingsError.style.display = "block";
  settingsError.innerHTML = message;
}

function setUserField(email, field, value, settingsError) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/users/" + email + "/" + field + "/set/" + value;
  fetch(url, {
    method: "post",
    body: JSON.stringify({}),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => {
      if (response.status != 200) {
        response.json().then((data) => {
          showError(data.msg, settingsError);
        })
      }
    });
}

function deleteUser(email, settingsError) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/users/" + email + "/delete";
  fetch(url, {
    method: "post",
    body: JSON.stringify({}),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => {
      if (response.status != 200) {
        response.json().then((data) => {
          showError(data.msg, settingsError);
        })
      } else {
        window.location.reload();
      }
    });
}

function setupEntries() {
  let settingsError = document.querySelector("#settings-error");

  let users = document.getElementsByClassName("user");
  for (let idx = 0; idx < users.length; idx++) {
    let user = users[idx];
    let isAdmin = user.querySelector(".is-admin").innerText == "true";
    let email = user.querySelector("td > .email").innerText;
    let isAdminEdit = user.querySelector("td > .is-admin-edit");
    isAdminEdit.checked = isAdmin;
    isAdminEdit.oncheckclick = function () {
      settingsSuccess.style.display = "none";
      settingsError.style.display = "none";
      setUserField(email, "is-admin", this.checked, settingsError);
    }
    let deleteButton = user.querySelector("td .delete");
    deleteButton.onclick = function () {
      settingsError.style.display = "none";
      deleteUser(email, settingsError);
    }
  }
}

document.addEventListener("DOMContentLoaded", function () {

  setupEntries();

});
