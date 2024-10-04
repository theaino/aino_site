function set_theme(value) {
	$("html").attr("data-bs-theme", value);
}

$(() => {
	let theme = Cookies.get("theme") || "dark";
	set_theme(theme);

	let toggle = $("#theme-toggle");
	toggle.on("click", () => {
		console.log(theme);
		theme = theme == "dark" ? "light" : "dark";
		Cookies.set("theme", theme);
		set_theme(theme);
	});
});
