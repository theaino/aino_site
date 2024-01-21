function setUserField(email, field, value) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/users/" + email + "/" + field + "/set/" + value;
  fetch(url, {
    method: "post",
    body: JSON.stringify({}),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => {
      console.log(response.status);
    });
}

function setupEntries() {
  let users = document.getElementsByClassName("user");
  for (let idx = 0; idx < users.length; idx++) {
    let user = users[idx];
    let isAdmin = user.querySelector(".is-admin").innerText == "true";
    let email = user.querySelector(".email").innerText;
    let isAdminEdit = user.querySelector(".is-admin-edit");
    isAdminEdit.checked = isAdmin;
    isAdminEdit.oncheckclick = function () {
      setUserField(email, "is-admin", this.checked);
    }
  }
}

document.addEventListener("DOMContentLoaded", function () {
  
  setupEntries();

});
