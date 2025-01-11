from django.conf import settings
from django.conf.urls.static import static
from django.urls import path
from .views import FoodListView

urlpatterns = [
    path('foods', FoodListView.as_view(), name='food-list'),
    path('featured', FoodListView.as_view(), name='featured')
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)