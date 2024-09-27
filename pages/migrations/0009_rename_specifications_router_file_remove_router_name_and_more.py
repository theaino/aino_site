# Generated by Django 5.1 on 2024-09-27 12:26

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('pages', '0008_post_like_ips'),
    ]

    operations = [
        migrations.RenameField(
            model_name='router',
            old_name='specifications',
            new_name='file',
        ),
        migrations.RemoveField(
            model_name='router',
            name='name',
        ),
        migrations.AlterField(
            model_name='post',
            name='like_ips',
            field=models.JSONField(default=dict, editable=False),
        ),
    ]
