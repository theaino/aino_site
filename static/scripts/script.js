let timeoutId = undefined;

function notify(value) {
  if (typeof(timeoutId) == Number) {
    clearTimeout(timeoutId);
  }
  let notifyElement = document.querySelector("#navbar-notify");
  notifyElement.innerText = value;
  timeoutId = setTimeout(function () {
    notifyElement.innerText = "";
  }, 3000);
}

