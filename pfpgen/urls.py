from django.urls import path
from pfpgen import views

urlpatterns = [
    path("neolib", views.neolib, name="neolib"),
]
