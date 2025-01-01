from django.shortcuts import render, redirect
from django.utils.timezone import localdate
from mood.models import Mood

today = localdate

def find_mood(date=None):
    if date is None:
        date = today()
    moods = Mood.objects.filter(date=date)
    return None if len(moods) == 0 else moods[0]

def index(request):
    if not request.user.is_superuser:
        return
    mood = find_mood()
    return render(request, "mood/index.html", {"mood": mood})


def submit(request, value):
    if not request.user.is_superuser:
        return
    mood = find_mood()
    if mood is None:
        mood = Mood.objects.create(date=today(), value=value)
    mood.value = value
    mood.save()
    return redirect("index")
