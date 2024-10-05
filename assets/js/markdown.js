const copy_button_label = "Copy";

$(() => {
	let elements = $("pre");

	elements.each((_idx, element) => {
		if (navigator.clipboard) {
			let button = $("<button></button>");
			$(element).children("code").after(button);

			button.text(copy_button_label);
			button.addClass("copy-code-button");

			button.on("click", async () => {
				await copy_code(element);
			});
		}
	});

	async function copy_code(element) {
		let code = $(element).children("code");
		let text = $(code).text();
		let lines = text.split("\n");
		lines = lines.map((line, idx) => {
			return line.slice(idx.toString().length);
		});
		text = lines.join("\n");

		await navigator.clipboard.writeText(text);
	}
});

$(() => {
	$(".markdown a").addClass("text-reset");
});
