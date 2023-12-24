let timeoutId = undefined;

function updateCheckboxButton(checkbox) {
  if (checkbox.checked) {
    checkbox.innerHTML = "<i class=\"fa-solid fa-check fa-xl\"></i>";
  } else {
    checkbox.innerHTML = "<i class=\"fa-solid fa-xmark fa-xl\"></i>";
  }
}

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

