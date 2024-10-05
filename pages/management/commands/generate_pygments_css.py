import os
from django.core.management.base import BaseCommand
from django.conf import settings


class Command(BaseCommand):
    help = "Generate Pygments CSS for code highlighting"

    def handle(self, *args, **kwargs):
        output_dir = os.path.join("static", "styles")
        output_file = os.path.join(output_dir, "pygments.css")

        if not os.path.exists(output_dir):
            os.makedirs(output_dir)

        css_content = settings.PYGMENTS_FORMATTER.get_style_defs(
            "." + settings.PYGMENTS_CSS_CLASS
        )

        with open(output_file, "w") as f:
            f.write(css_content)

        self.stdout.write(
            self.style.SUCCESS(f"Pygments CSS generated at {output_file}")
        )
