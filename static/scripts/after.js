function updateCheckboxButton(checkbox) {
  if (checkbox.checked) {
    checkbox.innerHTML = "<i class=\"fa-solid fa-check fa-xl\"></i>";
  } else {
    checkbox.innerHTML = "<i class=\"fa-solid fa-xmark fa-xl\"></i>";
  }
}

document.addEventListener("DOMContentLoaded", function () {
  
  let checkboxes = document.getElementsByClassName("checkbox");
  for (let idx = 0; idx < checkboxes.length; idx++) {
    checkboxes[idx].onclick = function () {
      this.checked = !this.checked;
      updateCheckboxButton(this);
    };
    updateCheckboxButton(checkboxes[idx]);
  }

});
