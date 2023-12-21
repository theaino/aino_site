function setValues(title, abstract, contents, public) {
  let host = location.protocol + "//" + location.host;
  let url = host + "/api/new-post";
  fetch(url, {
    method: "POST",
    body: JSON.stringify({
      title: title,
      abstract: abstract,
      contents: contents,
      public: public
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((response) => response.json())
    .then((data) => {
      window.location.replace(host + "/posts/" + data.id);
    })
}

function updatePublicButton(button, public) {
  if (public) {
    button.innerHTML = "Public";
  } else {
    button.innerHTML = "Private";
  }
}

document.addEventListener("DOMContentLoaded", function () {

  let title = document.querySelector("#title-edit");
  if (screen.availHeight > screen.availWidth) {
    title = document.querySelector("#mobile-title-edit")
  }
  let abstract = document.querySelector("#abstract-edit");
  let contents = document.querySelector("#contents-edit");
  let public = document.querySelector("#public-edit");
  let finish = document.querySelector("#finish-edit");

  let isPublicPost = false;

  updatePublicButton(public, isPublicPost);

  public.onclick = function () {
    isPublicPost = !isPublicPost;
    updatePublicButton(public, isPublicPost);
  }

  finish.onclick = function () {
    setValues(title.innerText, abstract.innerText, contents.innerText, isPublicPost);
  }

});
