## `.env`

```
DJANGO_SECRET_KEY
DJANGO_DB_NAME
DJANGO_DB_HOST
DJANGO_DB_PORT
DJANGO_DB_USER
DJANGO_DB_PASSWORD
DJANGO_DEBUG
```

Generate secret using
```sh
echo 'from django.core.management.utils import get_random_secret_key;print(get_random_secret_key())' | django-admin shell
```

## Setup

### Syntax highlighting

Generate css

```sh
python3 manage.py generate_pygments_css
```
