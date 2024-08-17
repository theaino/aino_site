from django.contrib import admin
from pages.models import Post, Router
from markdownx.admin import MarkdownxModelAdmin


@admin.register(Router)
class RouterAdmin(admin.ModelAdmin):
    pass


@admin.register(Post)
class PostAdmin(MarkdownxModelAdmin):
    view_on_site = True
