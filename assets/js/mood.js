$(() => {
  const maxMood = 11;
  $(".mood-color").each(function (index) {
    let mood = $(this).data("mood");
    let red = 0;
    let green = 0;
    let blue = 0;
    if (mood == -1) {
      red = 111;
      green = 111;
      blue = 111;
    } else {
      red = (255 * (9 - mood)) / maxMood;
      green = (255 * mood) / maxMood;
      blue = 0;
    }
    $(this).css("background-color", `rgb(${red}, ${green}, ${blue})`);
  });

  $(".mood-selected").addClass("border border-primary border-5");
});
