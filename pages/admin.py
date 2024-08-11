from django.contrib import admin
from pages.models import Post, Router


class RouterAdmin(admin.ModelAdmin):
    pass


class PostAdmin(admin.ModelAdmin):
    view_on_site = True

admin.site.register(Post, PostAdmin)
admin.site.register(Router, RouterAdmin)
