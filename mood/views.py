from django.shortcuts import render, redirect
from django.utils.timezone import localdate
from datetime import datetime, timedelta
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

def timeline(request):
    week_aligned_days = []
    moods = Mood.objects.all().order_by("date")
    current_date = moods[0].date
    current_date -= timedelta(days=current_date.weekday())
    end_date = moods[len(moods) - 1].date
    end_date += timedelta(6 - end_date.weekday())
    while current_date <= end_date:
        if current_date.weekday() == 0:
            week_aligned_days.append([])
        mood = moods.filter(date=current_date)
        if len(mood) == 0:
            mood = None
        else:
            mood = mood[0]
        week_aligned_days[-1].append(mood)
        current_date += timedelta(days=1)
    print(week_aligned_days)
    return render(request, "mood/timeline.html", {"weeks": week_aligned_days})
