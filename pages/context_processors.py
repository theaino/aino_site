def load_nav_obj(request):
    return {
        "NAV_PAGES": [{
            "name": "home",
            "url": "home",
            "urls": ["home"],
        },
        {
            "name": "about",
            "url": "about",
            "urls": ["about"],
        },
        {
            "name": "posts",
            "url": "posts",
            "urls": ["post", "posts", "posts_category"],
        }]
    }
