from rest_framework import status
from rest_framework.renderers import JSONRenderer
from rest_framework.response import Response
from rest_framework.views import APIView
from .models import Food, Category
from .serializers import FoodSerializer, CategorySerializer

class FoodListView(APIView):
    renderer_classes = [JSONRenderer]
    def get(self, request):
        category = request.query_params.get('category', None)
        if category:
            foods = Food.objects.filter(category__name=category)
        else:
            foods = Food.objects.all()

        serializer = FoodSerializer(foods, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

class CategoryListView(APIView):
    renderer_classes = [JSONRenderer]

    def get(self, request):
        categories = Category.objects.all()
        serializer = CategorySerializer(categories, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

class FeaturedFoodListView(APIView):
    renderer_classes = [JSONRenderer]

    def get(self, request):
        foods = Food.objects.filter(isFeatured=True)
        serializer = FoodSerializer(foods, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)
