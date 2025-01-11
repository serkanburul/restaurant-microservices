import React from 'react';
import {BrowserRouter} from 'react-router-dom';
import Navbar from './components/layout/Navbar';
import Footer from './components/layout/Footer';
import AppRoutes from './routes';

function App() {
    return (
        <BrowserRouter>
            <div className="flex flex-col min-h-screen relative bg-gray-100">
                <Navbar />
                <main id="main" className="flex-grow">
                    <AppRoutes />
                </main>
                <Footer/>
            </div>
        </BrowserRouter>
    );
}

export default App;