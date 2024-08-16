from django.db import models
from django.conf import settings
from django.urls import reverse
from markdownx.utils import markdownify
from markdownx.models import MarkdownxField
from bs4 import BeautifulSoup


class Router(models.Model):
    name = models.CharField(max_length=255)
    specifications = models.FileField(upload_to="router_specifications")

    def __str__(self):
        return self.name


class Post(models.Model):
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=255)
    body = MarkdownxField()
    public = models.BooleanField()
    created_on = models.DateTimeField(auto_now_add=True)
    last_modified = models.DateTimeField(auto_now=True)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.html_body = markdownify(str(self.body))
        self.words = len(BeautifulSoup(self.html_body, "html.parser").get_text().split())
        self.read_time = self.words // settings.WORDS_PER_MINUTE

    def get_absolute_url(self):
        return reverse("post", args=[str(self.pk)])

    def __str__(self):
        return self.title
