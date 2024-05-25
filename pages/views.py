from django.shortcuts import render
from pages.models import Post, Category
from django.conf import settings
from bs4 import BeautifulSoup
import markdown

def home(request):
    posts = Post.objects.all().order_by("-created_on")[:settings.POSTS_PER_PAGE]
    context = {
        "posts": posts,
    }
    return render(request, "pages/home.html", context)

def about(request):
    return render(request, "pages/about.html", {})

def posts(request, page=0):
    category = request.GET.get("category")
    has_category = category is not None and category.isdigit()
    if has_category:
        posts = Post.objects.filter(
            categories__pk=int(category)
        )
    else:
        posts = Post.objects.all()
    posts = posts.order_by("-created_on")
    post_count = len(posts)
    posts = posts[page*settings.POSTS_PER_PAGE:(page+1)*settings.POSTS_PER_PAGE]
    categories = Category.objects.all()
    context = {
        "page": page,
        "page_first": 0,
        "page_last": post_count // settings.POSTS_PER_PAGE,
        "page_before": page > 0,
        "page_after": (page+1)*settings.POSTS_PER_PAGE < post_count,
        "posts": posts,
        "categories": categories,
    }
    if has_category:
        context["selected_category"] = Category.objects.get(pk=int(category))
    return render(request, "pages/posts.html", context)

def post(request, pk):
    md = markdown.Markdown(extensions=[
        "markdown.extensions.fenced_code",
        "markdown.extensions.codehilite",
        "mdx_math",
    ])
    post = Post.objects.get(pk=pk)
    post.body = md.convert(post.body)
    post.words = len(BeautifulSoup(post.body, "html.parser").get_text().split())
    post.read_time = post.words // settings.WORDS_PER_MINUTE
    context = {
        "post": post,
    }
    return render(request, "pages/post.html", context)
