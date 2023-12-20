function setValues(title, abstract, contents) {
  let url = location.protocol + "//" + location.host + "/api/posts/" + _post_id + "/edit";
  fetch(url, {
    method: "POST",
    body: JSON.stringify({
      title: title,
      abstract: abstract,
      contents: contents
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  });
  notify("Successfully edited post");
}

document.addEventListener("DOMContentLoaded", function () {

  let title = document.querySelector("#title-edit");
  let abstract = document.querySelector("#abstract-edit");
  let contents = document.querySelector("#contents-edit");
  let finish = document.querySelector("#finish-edit");

  finish.onclick = function () {
    setValues(title.innerText, abstract.innerText, contents.value);
  }

});
