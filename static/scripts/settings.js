const typeLookup = {
  "int": "number",
  "str": "text",
  "bool": "checkbox"
}

function getTypeElement(type, id) {
  return `<input type="${typeLookup[type]}" id="${id}">`;
}

function setTypeElementValue(type, id, value) {
  let element = document.getElementById(id);
  switch (type) {
    case "int":
      element.value = value;
      break;
    case "str":
      element.value = value;
      break;
    case "bool":
      element.checked = value == "true";
      break;
  }
}

function getTypeElementValue(type, id) {
  let element = document.getElementById(id);
  let value = undefined;
  switch (type) {
    case "int":
      value = element.value;
      break;
    case "str":
      value = element.value;
      break;
    case "bool":
      value = element.checked ? "true" : "false";
      break;
  }
  return value;
}

function setEntryInputs() {
  let entries = document.getElementsByClassName("entry");
  for (let idx = 0; idx < entries.length; idx++) {
    let entry = entries.item(idx);
    let key = entry.querySelector(".key").innerText;
    let type = entry.querySelector(".type").innerText;
    let rawValue = entry.querySelector(".value").innerText;
    let inputWrapper = entry.querySelector(".input");
    inputWrapper.innerHTML = getTypeElement(type, "input-" + key);
    setTypeElementValue(type, "input-" + key, rawValue);
  }
}

document.addEventListener("DOMContentLoaded", function () {
  setEntryInputs();
});
