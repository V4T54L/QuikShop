"use client"

import { useState } from 'react';
import Link from 'next/link';
import { FiMapPin, FiSearch, FiShoppingCart } from 'react-icons/fi';

export default function Navbar() {
  const [cartCount, setCartCount] = useState(3); // Example cart count

  return (
    <nav className="bg-gray-900 text-white shadow-md sticky top-0 z-50">
      <div className="container mx-auto flex items-center justify-between px-6 py-4 gap-20">
        {/* Logo */}
        <div className="text-2xl font-bold mb-2 sm:mb-0">
          <Link href="/" className="hover:text-gray-400 transition-colors">
            MyShop
          </Link>
        </div>

        {/* Location */}
        <div className="flex items-center space-x-2 mb-2 sm:mb-0 animate-fade-in">
          <FiMapPin className="text-gray-400 text-2xl" />
          <span className="text-gray-300 hover:text-white transition-colors text-lg">Your Location</span>
        </div>

        {/* Search Box */}
        <div className="relative flex-1 mx-4 mb-2 sm:mb-0">
          <input
            type="text"
            placeholder="Search..."
            className="w-full px-4 py-2 rounded-lg bg-gray-800 text-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-600"
          />
          <FiSearch className="absolute top-1/2 right-3 transform -translate-y-1/2 text-gray-400 hover:text-white transition-colors" />
        </div>

        {/* Cart Icon */}
        <div className="relative">
          <FiShoppingCart className="text-2xl hover:text-gray-400 transition-transform transform hover:scale-110 cursor-pointer" />
          {cartCount > 0 && (
            <span className="absolute -top-2 -right-2 bg-red-600 text-white text-xs font-bold w-5 h-5 flex items-center justify-center rounded-full animate-pulse">
              {cartCount}
            </span>
          )}
        </div>
      </div>

      {/* Responsive Styling */}
      <style jsx>{`
        @keyframes fade-in {
          from {
            opacity: 0;
            transform: translateX(-10px);
          }
          to {
            opacity: 1;
            transform: translateX(0);
          }
        }

        .animate-fade-in {
          animation: fade-in 0.5s ease-in-out;
        }
      `}</style>
    </nav>
  );
}
