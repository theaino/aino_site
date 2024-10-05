import Cookies from "js-cookie";

function set_theme(value) {
  $("html").attr("data-bs-theme", value);
  Cookies.set("theme", value);
}

$(() => {
  let theme = Cookies.get("theme") || "dark";
  set_theme(theme);

	let toggle = $("#theme-toggle");
	toggle.on("click", () => {
		theme = theme == "dark" ? "light" : "dark";
		set_theme(theme);
	});
});
