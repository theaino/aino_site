from django.db import models
from django.db.models.signals import post_delete, pre_save
from django.dispatch import receiver
from django.conf import settings
from django.urls import reverse
from markdownx.utils import markdownify
from markdownx.models import MarkdownxField
from bs4 import BeautifulSoup
import os


class Router(models.Model):
    file = models.FileField(upload_to="router_specifications")

    def __str__(self):
        return self.file.path.split("/")[-1]


@receiver(post_delete, sender=Router)
def delete_specification_file(_sender, instance, **kwargs):
    if instance.file and os.path.isfile(instance.file.path):
        os.remove(instance.file.path)


@receiver(pre_save, sender=Router)
def delete_old_file_on_update(_sender, instance, **kwargs):
    if not instance.pk:
        return False
    try:
        old_file = Router.objects.get(pk=instance.pk).file
    except Router.DoesNotExist:
        return False
    new_file = instance.file
    if old_file != new_file and old_file and os.path.isfile(old_file.path):
        os.remove(old_file.path)


class Post(models.Model):
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=255)
    body = MarkdownxField()
    public = models.BooleanField()
    like_ips = models.JSONField(editable=False, default=dict)
    created_on = models.DateTimeField(auto_now_add=True)
    last_modified = models.DateTimeField(auto_now=True)

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.html_body = markdownify(str(self.body))
        self.words = len(
            BeautifulSoup(self.html_body, "html.parser").get_text().split()
        )
        self.read_time = self.words // settings.WORDS_PER_MINUTE

    def get_absolute_url(self):
        return reverse("post", args=[str(self.pk)])

    def __str__(self):
        return self.title
