const typeLookup = {
  "int": "number",
  "str": "text",
  "bool": "checkbox\" class=\"checkbox\""
}

function getTypeElement(type, id) {
  switch (type) {
    case "int":
      return "<input type=\"number\" id=\"" + id + "\">";
    case "str":
      return "<input type=\"text\" id=\"" + id + "\">";
    case "bool":
      return "<button class=\"square checkbox\" id=\"" + id + "\"></button>"
  }
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
      updateCheckboxButton(element);
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
  let keys = [];
  let entries = document.getElementsByClassName("entry");
  for (let idx = 0; idx < entries.length; idx++) {
    let entry = entries.item(idx);
    let key = entry.querySelector(".key").innerText;
    let type = entry.querySelector(".type").innerText;
    let rawValue = entry.querySelector(".value").innerText;
    let defaultValue = entry.querySelector(".default-value").innerText;
    let inputWrapper = entry.querySelector(".input");
    let reset = entry.querySelector(".reset button");
    inputWrapper.innerHTML = getTypeElement(type, "input-" + key);
    setTypeElementValue(type, "input-" + key, rawValue);
    reset.onclick = function () {
      setTypeElementValue(type, "input-" + key, defaultValue);
    };
    keys.push([key, type]);
  }
  return keys;
}

function updateKeys(map, settingsError, settingsSuccess) {
  let host = location.protocol + "//" + location.host;
  Object.keys(map).forEach((key) => {
    let value = map[key];
    console.log(key, value);
    let url = host + "/api/settings/" + key + "/set/" + value;
    fetch(url, {
      method: "POST",
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
          showSuccess("Successfully saved settings", settingsSuccess);
        }
      });
  });
}

function getSettingsMap(keys) {
  map = {};
  keys.forEach(([key, type]) => {
    map[key] = getTypeElementValue(type, "input-" + key);
  });
  return map;
}

function showError(message, settingsError) {
  settingsError.style.display = "block";
  settingsError.innerHTML = message;
}

function showSuccess(message, settingsSuccess) {
  settingsSuccess.style.display = "block";
  settingsSuccess.innerHTML = message;
}

document.addEventListener("DOMContentLoaded", function () {

  let settingsError = document.querySelector("#settings-error");
  let settingsSuccess = document.querySelector("#settings-success");
  let save = document.querySelector("#save");
  let keys = setEntryInputs();

  save.onclick = function () {
    settingsError.style.display = "none";
    settingsSuccess.style.display = "none";
    let map = getSettingsMap(keys);
    updateKeys(map, settingsError, settingsSuccess);
  };

});

