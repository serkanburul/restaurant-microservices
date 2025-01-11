import React from 'react';
import {Link, useLocation} from 'react-router-dom';
import { Button } from "@/components/ui/button";
import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuList
} from "@/components/ui/navigation-menu";

const Navbar = () => {
    const linkClass = "text-white px-3 py-2 hover:bg-black hover:backdrop-blur hover:bg-opacity-80 rounded-md transition-colors duration-300 font-poppins font-bold";
    const linkClass2 = "text-black px-3 py-2 hover:bg-white hover:backdrop-blur hover:bg-opacity-80 rounded-md transition-colors duration-300 font-poppins font-bold";
    const location = useLocation()
    let isHome = location.pathname === "/" || location.pathname === "/home";

    return (
        <nav className= "shadow-md p-4 fixed top-0 left-0 right-0 z-50 bg-opacity-80 ">
            <div className="container mx-auto flex justify-between items-center">
                <Link to="/" className="text-2xl font-bold text-red-600">
                    RESTAURANT
                </Link>
                <NavigationMenu>
                    <NavigationMenuList>
                        <NavigationMenuItem>
                            <Link to="/" className={isHome ? linkClass : linkClass2}>
                                HOME
                            </Link>
                        </NavigationMenuItem>
                        <NavigationMenuItem>
                            <Link to="/menu" className={isHome ? linkClass : linkClass2}>
                                MENU
                            </Link>
                        </NavigationMenuItem>
                    </NavigationMenuList>
                </NavigationMenu>
                <Link to="/reservation">
                    <Button variant="secondary" className="rounded-3xl p-5">BOOK A TABLE</Button>
                </Link>
            </div>
        </nav>
    );
};

export default Navbar;
