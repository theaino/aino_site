function deletePost(id) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/posts/" + id + "/delete";
  fetch(url, {
    method: "POST",
    body: JSON.stringify({}),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => {
      window.location.replace(host + "/posts")
    });
}

document.addEventListener("DOMContentLoaded", function () {

  if (_authed) {
    let deleteButton = document.querySelector("#delete");
    deleteButton.onclick = function () {
      if (confirm("Are you sure you want to delete this post?")) {
        deletePost(_post_id);
      }
    }
  }

});
