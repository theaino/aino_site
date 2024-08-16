from django.shortcuts import render
from pages.models import Post
from django.conf import settings
from django.http import HttpResponse


def get_client_ip(request):
    x_forwarded_for = request.META.get("HTTP_X_FORWARDED_FOR")
    if x_forwarded_for:
        ip = x_forwarded_for.split(",")[0]
    else:
        ip = request.META.get("REMOTE_ADDR")
    return ip


def home(request):
    posts = Post.objects.all().order_by("-created_on")
    if not request.user.is_authenticated:
        posts = posts.filter(public=True)
    posts = posts[:settings.POSTS_PER_PAGE]
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
    if not request.user.is_authenticated:
        posts = posts.filter(public=True)
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
    ip = get_client_ip(request)
    context = {
        "post": post,
        "liked": ip in post.like_ips.keys() and post.like_ips[ip],
        "like_count": sum([1 if x else 0 for x in post.like_ips.values()])
    }
    return render(request, "pages/post.html", context)


def like_post(request, pk, like):
    post = Post.objects.get(pk=pk)
    ip = get_client_ip(request)
    post.like_ips[ip] = like != 0
    post.save()
    return HttpResponse(sum([1 if x else 0 for x in post.like_ips.values()]))
