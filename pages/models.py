from django.db import models
from django.conf import settings
import markdown
from bs4 import BeautifulSoup


class Router(models.Model):
    name = models.CharField(max_length=255)
    specifications = models.FileField(upload_to="router_specifications")

    def __str__(self):
        return self.name


class Post(models.Model):
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=255)
    body = models.TextField()
    created_on = models.DateTimeField(auto_now_add=True)
    last_modified = models.DateTimeField(auto_now=True)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        md = markdown.Markdown(extensions=[
            "markdown.extensions.fenced_code",
            "markdown.extensions.codehilite",
            "mdx_math"
        ])
        self.html_body = md.convert(str(self.body))
        self.words = len(BeautifulSoup(self.html_body, "html.parser").get_text().split())
        self.read_time = self.words // settings.WORDS_PER_MINUTE

    def __str__(self):
        return self.title
