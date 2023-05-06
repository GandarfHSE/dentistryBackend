import requests

image = {"image_data": open("testimg.jpg", 'rb')}
payload = {
    "image_extension": "jpg"
}
requests.put(url="http://localhost:8083/image/upload", files=image, data=payload)
