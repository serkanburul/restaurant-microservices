from django.contrib import admin
from .models import Category, Food

@admin.register(Category)
class CategoryAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'created_at')
    search_fields = ('name',)
    ordering = ('-created_at',)

@admin.register(Food)
class FoodAdmin(admin.ModelAdmin):
    list_display = ('id', 'name', 'price', 'category')
    search_fields = ('name', 'explanation')
    list_filter = ('category',)
    ordering = ('id',)

admin.site.site_header = 'Restaurant'
admin.site.site_title = 'Restaurant'
admin.site.index_title = 'Restaurant'