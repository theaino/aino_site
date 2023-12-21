const titles = ["Hi, I'm Aino!", "Hello, World!", "Hi, I write code."];

function prefix(words){
  if (!words[0] || words.length == 1) return words[0] || "";
  let i = 0;
  while(words[0][i] && words.every(w => w[i] === words[0][i])) {
    i++;
  }
  return i;
}

let currentTitle = "";
let newTitle = "";
let remove = true;
let titleElement = undefined;

function randomTitle() {
  let idx = Math.floor(Math.random() * titles.length);
  return titles[idx];
}

function nextCharacter() {
  let timeout = 150;

  if (titleElement.innerHTML.length === prefix([currentTitle, newTitle])) {
    remove = false;
  }

  if (titleElement.innerHTML.length === prefix([currentTitle, newTitle]) + 1 && remove) {
    timeout = 500;
  }

  if (remove) {
    titleElement.innerHTML = titleElement.innerHTML.slice(0, -1);
  } else {
    titleElement.innerHTML = newTitle.substring(0, titleElement.innerHTML.length + 1);
  }


  if (titleElement.innerHTML === newTitle) {
    titleElement.classList.remove("paused");
    nextTitle();
  } else {
    setTimeout(function () {
      nextCharacter();
    }, timeout);
  }
}

function chooseNewTitle() {
  if (newTitle === "") {
    newTitle = randomTitle();
  }
  currentTitle = newTitle;
  do {
    newTitle = randomTitle();
  } while (newTitle === currentTitle);
}

function nextTitle() {
  chooseNewTitle();
  titleElement.innerHTML = currentTitle;
  remove = true;
  setTimeout(function () {
    titleElement.classList.add("paused");
    nextCharacter();
  }, 5000);
}

document.addEventListener("DOMContentLoaded", function () {

  titleElement = document.querySelector("#title");

  nextTitle();

});
