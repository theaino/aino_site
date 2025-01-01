$(() => {
	const maxMood = 11;
	$(".mood-card").each(function(index) {
		let mood = $(this).data("mood");
		console.log(mood);
		let red = 255 * (9 - mood) / maxMood;
		let green = 255 * mood / maxMood;
		$(this).css("background-color", `rgb(${red}, ${green}, 0)`);
	})

	$(".mood-selected").addClass("border border-primary border-5");
});
