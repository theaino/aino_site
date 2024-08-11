from django.contrib import admin
from pages.models import Post, Router

class PostAdmin(admin.ModelAdmin):
    pass

admin.site.register(Post, PostAdmin)
admin.site.register(Router)
