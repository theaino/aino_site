import os
from django.core.management.base import BaseCommand
from pygments.formatters import HtmlFormatter

class Command(BaseCommand):
    help = "Generate Pygments CSS for code highlighting"

    def handle(self, *args, **kwargs):
        output_dir = os.path.join("static", "styles")
        output_file = os.path.join(output_dir, "pygments.css")

        if not os.path.exists(output_dir):
            os.makedirs(output_dir)

        formatter = HtmlFormatter(cssfile="pygments.css", style="one-dark")
        css_content = formatter.get_style_defs(".codehilite")

        with open(output_file, "w") as f:
            f.write(css_content)

        self.stdout.write(self.style.SUCCESS(f"Pygments CSS generated at {output_file}"))
