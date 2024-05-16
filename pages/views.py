from django.shortcuts import render
from pages.models import Post, Category

def home(request):
    return render(request, "pages/home.html", {})

def about(request):
    return render(request, "pages/about.html", {})

def posts(request):
    posts = Post.objects.all().order_by("-created_on")
    categories = Category.objects.all()
    context = {
        "posts": posts,
        "categories": categories,
    }
    return render(request, "pages/posts.html", context)

def post(request, pk):
    post = Post.objects.get(pk=pk)
    context = {
        "post": post,
    }
    return render(request, "pages/post.html", context)

def posts_category(request, pk):
    posts = Post.objects.filter(
        categories__pk=pk
    ).order_by(
        "-created_on"
    )
    categories = Category.objects.all()
    selected_category = Category.objects.get(pk=pk)
    context = {
        "posts": posts,
        "selected_category": selected_category,
        "categories": categories,
    }
    return render(request, "pages/posts.html", context)
