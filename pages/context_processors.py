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
                "github": "https://github.com/theaino",
                "github_site": "https://github.com/theaino/aino_site",
                "instagram": "https://instagram.com/aino.spring",
                "email": "info@aino-spring.com",
                }
            }

def load_sites(request):
    return {
            "SITES": {
                "searxng": "https://search.aino-spring.com"
                }
            }
