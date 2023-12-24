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
    let inputWrapper = entry.querySelector(".input");
    inputWrapper.innerHTML = getTypeElement(type, "input-" + key);
    setTypeElementValue(type, "input-" + key, rawValue);
    keys.push([key, type]);
  }
  return keys;
}

function updateKeys(map) {
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
        if (response.status != 400) {
          response.json().then((data) => {
            console.log(data.msg);
          })
        } else {
          console.log("success");
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

document.addEventListener("DOMContentLoaded", function () {

  let save = document.querySelector("#save");
  let keys = setEntryInputs();

  save.onclick = function () {
    let map = getSettingsMap(keys);
    updateKeys(map);
  };

});

