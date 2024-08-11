from django.shortcuts import render
from pages.models import Post
from django.conf import settings

def home(request):
    posts = Post.objects.all().order_by("-created_on")[:settings.POSTS_PER_PAGE]
    context = {
        "posts": posts,
    }
    return render(request, "pages/home.html", context)

def about(request):
    return render(request, "pages/about.html", {})

def posts(request, page=0):
    query = request.GET.get("q", "")
    do_search = query != ""
    posts = Post.objects.all()
    if do_search:
        posts = posts.filter(title__icontains=query) | posts.filter(description__icontains=query) | posts.filter(body__icontains=query)
    posts = posts.order_by("-created_on")
    post_count = len(posts)
    posts = posts[page*settings.POSTS_PER_PAGE:(page+1)*settings.POSTS_PER_PAGE]
    context = {
        "page": page,
        "page_first": 0,
        "page_last": post_count // settings.POSTS_PER_PAGE,
        "page_before": page > 0,
        "page_after": (page+1)*settings.POSTS_PER_PAGE < post_count,
        "posts": posts,
        "query": query
    }
    return render(request, "pages/posts.html", context)

def post(request, pk):
    post = Post.objects.get(pk=pk)
    context = {
        "post": post,
    }
    return render(request, "pages/post.html", context)
