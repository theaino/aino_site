from django.contrib import admin
from mood.models import Mood

@admin.register(Mood)
class RouterAdmin(admin.ModelAdmin):
    pass

