function interpRedGreen(t) {
  t = 1 - Math.max(0, Math.min(1, t));

  const redHSL = { h: 0, s: 1, l: 0.5 };
  const greenHSL = { h: 120, s: 1, l: 0.5 };

  const h = redHSL.h + (greenHSL.h - redHSL.h) * t;
  const s = redHSL.s + (greenHSL.s - redHSL.s) * t;
  const l = redHSL.l + (greenHSL.l - redHSL.l) * t;

  const rgb = hslToRgb(h, s, l);

  return rgb;
}

function hslToRgb(h, s, l) {
  h /= 360;
  let r, g, b;

  if (s === 0) {
    r = g = b = l;
  } else {
    const hueToRgb = (p, q, t) => {
      if (t < 0) t += 1;
      if (t > 1) t -= 1;
      if (t < 1 / 6) return p + (q - p) * 6 * t;
      if (t < 1 / 2) return q;
      if (t < 2 / 3) return p + (q - p) * (2 / 3 - t) * 6;
      return p;
    };

    const q = l < 0.5 ? l * (1 + s) : l + s - l * s;
    const p = 2 * l - q;
    r = hueToRgb(p, q, h + 1 / 3);
    g = hueToRgb(p, q, h);
    b = hueToRgb(p, q, h - 1 / 3);
  }

  return [r, g, b];
}

$(() => {
  const maxMood = 11;
  $(".mood-color").each(function (_index) {
    let mood = $(this).data("mood");
    let red = 0;
    let green = 0;
    let blue = 0;
    if (mood == -1) {
      red = 111;
      green = 111;
      blue = 111;
    } else {
      //red = 255 * (maxMood - mood) / maxMood;
      //green = 255 * mood / maxMood;
      //blue = 0;
      let color = interpRedGreen((maxMood - mood) / maxMood);
      console.log(color);
      red = color[0] * 255;
      green = color[1] * 255;
      blue = color[2] * 255;
    }
    $(this).css("background-color", `rgb(${red}, ${green}, ${blue})`);
  });

  $(".mood-selected").addClass("border border-primary border-5");
});
