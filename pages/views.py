from django.shortcuts import render
from pages.models import Post

def home(request):
    return render(request, "pages/home.html", {})

def about(request):
    return render(request, "pages/about.html", {})

def posts(request):
    posts = Post.objects.all().order_by("-created_on")
    context = {
        "posts": posts,
    }
    return render(request, "pages/posts.html", context)

def post(request, pk):
    post = Post.objects.get(pk=pk)
    context = {
        "post": post,
    }
    return render(request, "pages/post.html", context)
