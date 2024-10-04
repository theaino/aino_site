from markdown.inlinepatterns import InlineProcessor
from markdown.extensions import Extension
import xml.etree.ElementTree as etree


class ImgInlineProcessor(InlineProcessor):
    def handleMatch(self, m, data):
        alt_text = m.group(1)
        url = m.group(2)
        has_size = len(m.groups()) > 2
        width = m.group(3) if has_size else ""
        height = m.group(4) if has_size else ""

        el = etree.Element("img")
        el.set("src", url)
        el.set("alt", alt_text)

        styles = []
        if width:
            styles.append(f"width:{width}")
        if height:
            styles.append(f"height:{height}")

        if not (width or height):
            styles.append(f"width:100%")

        el.set("style", ";".join(styles))

        return el, m.start(0), m.end(0)


class ImgExtension(Extension):
    def extendMarkdown(self, md):
        IMAGE_RE = r"!\[(.*)\]\(([^\s]+)(?: +(.*)x(.*))?\)"
        md.inlinePatterns.register(ImgInlineProcessor(IMAGE_RE, md), "sized-image", 175)


def makeExtension(*args, **kwargs):
    return ImgExtension(*args, **kwargs)
