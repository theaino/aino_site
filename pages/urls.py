from django.urls import path
from pages import views

urlpatterns = [
    path("", views.home, name="home"),
    path("about/", views.about, name="about"),
    path("posts/", views.posts, name="posts"),
    path("posts/<int:pk>/", views.post, name="post"),
]
