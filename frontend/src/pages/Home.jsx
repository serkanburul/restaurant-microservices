import React, {useEffect, useState} from 'react';
import { Link } from 'react-router-dom';
import { Button } from "@/components/ui/button";

const Home = () => {
    let [menuItems, setMenuItems]= useState([])

    useEffect(() => {
        fetch('http://localhost:8080/content/featured')
            .then(response => response.json())
            .then(data => setMenuItems(data))
            .catch((error) => console.log(error));
    },[])

    return (
        <>
            <div className={"pt-20 pb-80 flex items-center justify-center"} style={{
                backgroundImage: 'url("/a.png")',
                backgroundSize: 'cover',
                backgroundPosition: "center",
                backgroundRepeat: "no-repeat"
            }}>
                <div className="w-2/5 text-center text-white justify-center border-4 border-accent-foreground
                            rounded-2xl p-8 h-full border-white">
                    <h2 className="text-white text-4xl font-bold text-gray-700 font-poppins">RESTAURANT</h2>
                    <p className="text-white text-gray-700 font-poppins my-4">
                        Lorem ipsum dolor sit amet, consectetur adipiscing elit. In auctor, sapien ac fermentum luctus, nibh ligula eleifend orci,
                        quis fermentum nisi elit non massa. Proin sed sagittis dui, consectetur mollis nisi. Quisque eleifend elit a leo posuere porta.
                        Vivamus euismod lacinia nisl in dapibus. Nam ut mauris ac ante semper egestas et in velit. Maecenas finibus tortor nec leo pellentesque faucibus.
                        Proin commodo congue tincidunt. Quisque vitae faucibus magna, vel bibendum sem.
                    </p>
                    <Link to="/reservation">
                        <Button variant="secondary">
                            Book a Table
                        </Button>
                    </Link>
                </div>
            </div>

            <section className="container mx-auto py-12">
                <h2 className="text-3xl font-bold text-center mb-8 text-black">Featured Flavors</h2>
                <div className="grid md:grid-cols-3 gap-6">
                    {menuItems.slice(0, 3).map(item => (
                        <div
                            key={item.id}
                            className=" shadow-md rounded-lg overflow-hiddenx border-4 border-gray-100"
                        >
                            <div className="p-6">
                                <h3 className="text-xl font-semibold mb-2">{item["name"]}</h3>
                                <p className="text-gray-600 mb-4">{item["explanation"]}</p>
                                <p className="text-red-600 font-bold">{item["price"]}</p>
                            </div>
                        </div>
                    ))}
                </div>
                <div className="text-center mt-8">
                    <Link to="/menu#">
                        <Button>All Menu</Button>
                    </Link>
                </div>
            </section>
        </>

    );
};

export default Home;