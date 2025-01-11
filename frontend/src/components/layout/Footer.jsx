import React from "react";
import { Button } from "@/components/ui/button"
import { Mail, LucideInstagram, LucideFacebook, LucideTwitter, LucideLinkedin } from "lucide-react"


const Footer = () => {
    const [email, setEmail] = React.useState("");
    async function handleSubmit(e) {
        let x
        e.preventDefault();
        await fetch("http://localhost:8080/subscription", {
            method: "POST",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json;charset=UTF-8"
            },
            body: JSON.stringify({ mailAddress: email })
        }).then(r => r.json())
            .then(json => {
                x = json
            })
        alert(x)
    }

    return (
        <footer className="bg-gray-800 text-white py-10">
            <div className="container mx-auto grid grid-cols-1 md:grid-cols-4 gap-8">
                <div>
                    <h2 className="text-2xl font-bold mb-4">Restaurant</h2>
                    <p className="text-sm text-gray-400">
                        Subscribe to be informed about events.
                    </p>
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-4">Quick Access</h3>
                    <ul className="space-y-2">
                        <li><a href="#" className="hover:text-gray-300">Home</a></li>
                        <li><a href="#" className="hover:text-gray-300">Menu</a></li>
                        <li><a href="#" className="hover:text-gray-300">About Us</a></li>
                    </ul>
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-4">Newsletter</h3>
                    <p className="text-sm text-gray-400 mb-4">
                        Subscribe to be informed about events.
                    </p>
                    <form onSubmit={handleSubmit}>
                        <input
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
                            type="email"
                            placeholder="Email Address"
                            className="w-full px-4 py-2 rounded bg-gray-800 text-gray-200 mb-2 focus:outline-none"
                        />
                        <Button variant="destructive" className="w-full"><Mail/> Subscribe </Button>
                    </form>
                </div>

                <div>
                    <h3 className="text-xl font-semibold mb-4">Follow Us</h3>
                    <div className="flex space-x-4">
                        <a href="#" className="hover:text-gray-300"><LucideFacebook/>Facebook</a>
                        <a href="#" className="hover:text-gray-300"><LucideTwitter/>Twitter</a>
                        <a href="#" className="hover:text-gray-300"><LucideInstagram/>Instagram</a>
                        <a href="#" className="hover:text-gray-300"><LucideLinkedin/>LinkedIn</a>
                    </div>
                </div>
            </div>

            <div className="mt-10 border-t border-gray-700 pt-4 text-center text-sm text-gray-400">
                © 2024 All Rights Reserved.
            </div>
        </footer>
    );
};

export default Footer;
