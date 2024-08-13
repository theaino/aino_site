from django.contrib import admin
from pages.models import Post, Router
from markdownx.admin import MarkdownxModelAdmin


class RouterAdmin(admin.ModelAdmin):
    pass


class PostAdmin(MarkdownxModelAdmin):
    view_on_site = True

admin.site.register(Post, PostAdmin)
admin.site.register(Router, RouterAdmin)
