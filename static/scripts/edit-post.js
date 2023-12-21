function setValues(title, abstract, contents, public) {
  let url = location.protocol + "//" + location.host + "/api/posts/" + _post_id + "/edit";
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
  }).then((response) => {
      notify("Successfully edited post");
    });
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
  let abstract = document.querySelector("#abstract-edit");
  let contents = document.querySelector("#contents-edit");
  let public = document.querySelector("#public-edit");
  let finish = document.querySelector("#finish-edit");

  let isPublicPost = _post_public;

  updatePublicButton(public, isPublicPost);

  public.onclick = function () {
    isPublicPost = !isPublicPost;
    updatePublicButton(public, isPublicPost);
  }

  finish.onclick = function () {
    setValues(title.innerText, abstract.innerText, contents.value, isPublicPost);
  }

});
