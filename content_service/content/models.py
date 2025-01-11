from django.db import models

class Category(models.Model):
    name = models.CharField(max_length=255)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return self.name


class Food(models.Model):
    name = models.CharField(max_length=255)
    image = models.ImageField(upload_to='food_images/')
    explanation = models.CharField(max_length=255)
    price = models.PositiveIntegerField()
    category = models.ForeignKey(Category, on_delete=models.CASCADE, related_name='foods')
    isFeatured = models.BooleanField(default=False)

    def __str__(self):
        return self.name
