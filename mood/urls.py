from django.urls import path
from mood import views

urlpatterns = [
    path("", views.index, name="index"),
    path("<int:value>", views.submit, name="submit"),
    path("timeline", views.timeline, name="timeline"),
]
