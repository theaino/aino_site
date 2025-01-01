from django.db import models

class Mood(models.Model):
    value = models.IntegerField()
    date = models.DateField(editable=True)

    def __str__(self):
        return f"{self.value} | {str(self.date)}"
