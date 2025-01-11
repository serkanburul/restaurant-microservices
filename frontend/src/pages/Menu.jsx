import React, {useEffect, useState} from 'react';
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { menuCategories } from '../data/menuItems';

const Menu = () => {
    const [selectedCategory, setSelectedCategory] = useState('');
    const [menuItems, setMenuItems] = useState([]);

    useEffect(() => {
        fetch(`http://localhost:8080/content/foods?category=${selectedCategory}`).then((response) => {
            if (!response.ok) {
                throw new Error("Failed to fetch foods: " + response.status)
            }
            return response.json();
        })
            .then(data => setMenuItems(data))
            .catch(error => console.log(error));
    }, [selectedCategory]);


    return (
        <div className="container mx-auto py-24">
            <h1 className="text-4xl font-bold text-center mb-8">Menu</h1>

            <div className="flex justify-center space-x-4 mb-8">
                <button
                    onClick={() => setSelectedCategory('')}
                    className={`px-4 py-2 rounded ${
                        selectedCategory === ''
                            ? 'bg-red-600 text-white'
                            : 'bg-gray-200 text-gray-800'
                    }`}
                >
                    All
                </button>
                {menuCategories.map(category => (
                    <button
                        key={category}
                        onClick={() => setSelectedCategory(category)}
                        className={`px-4 py-2 rounded ${
                            selectedCategory === category
                                ? 'bg-red-600 text-white'
                                : 'bg-gray-200 text-gray-800'
                        }`}
                    >
                        {category}
                    </button>
                ))}
            </div>

            <div className="grid md:grid-cols-3 gap-6">
                {menuItems.map(item => (
                    <Card key={item.id}>
                        <CardHeader>
                            <CardTitle>{item.name}</CardTitle>
                        </CardHeader>
                        <CardContent>
                            <img className="mx-auto w-56 rounded-xl" src={"http://localhost:8080/content"+item.image}/>
                            <p className="text-gray-600 mb-2">{item["explanation"]}</p>
                            <div className="font-bold text-red-600">{item.price}</div>
                        </CardContent>
                    </Card>
                ))}
            </div>
        </div>
    );
};

export default Menu;