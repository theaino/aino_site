def load_nav_obj(request):
    return {
        "NAV_PAGES": [{
            "name": "home",
            "url": "home",
            "urls": ["home"],
        },
        {
            "name": "posts",
            "url": "posts",
            "urls": ["post", "posts"],
        }]
    }

def load_contact(request):
    return {
        "CONTACT": {
            "github": "https://github.com/AinoSpring",
            "instagram": "https://instagram.com/ainospring",
            "email": "info@aino-spring.com",
        }
    }