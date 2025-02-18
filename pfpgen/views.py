from django.shortcuts import render, redirect
from django.http import HttpResponse
from django.conf import settings
from PIL import Image
import io
import os


NEOLIB_THICKNESS = 0.228125

def neolib(request):
    if request.method == "POST":
        file = request.FILES["image"]
        res = handle_upload(file)
        response = HttpResponse(res, content_type="image/png")
        response["Content-Disposition"] = f"attachment; filename=neolib_{file.name}"
        return response
    return render(request, "pfpgen/neolib.html", {})

def handle_upload(f):
    image = Image.open(io.BytesIO(f.read())).convert("RGBA")

    border_path = os.path.join(settings.STATIC_ROOT, "images/neoli_border.png")
    border = Image.open(border_path).convert("RGBA")
    
    image = image.resize([max(border.size)] * 2)

    border_size = int(0.1140625 * image.size[0])

    width, height = image.size
    border = border.resize((width + border_size * 2, height + border_size * 2))

    bordered_image = Image.new("RGBA", border.size, (0, 0, 0, 0))
    bordered_image.paste(image, (border_size, border_size), image)
    bordered_image = Image.alpha_composite(bordered_image, border)

    img_io = io.BytesIO()
    bordered_image.save(img_io, format="PNG")
    img_io.seek(0)
    return img_io.getvalue()
