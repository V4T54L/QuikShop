"use client"

import Link from "next/link";
import { useState } from "react";
import { FiChevronDown } from "react-icons/fi";

const categories = [
  {
    name: "Electronics",
    items: [
      { name: "Laptops", href: "/category/laptops" },
      { name: "Mobile Phones", href: "/category/mobile-phones" },
      { name: "Headphones", href: "/category/headphones" },
      { name: "Cameras", href: "/category/cameras" },
      { name: "Smart Watches", href: "/category/smart-watches" },
    ],
  },
  {
    name: "Fashion",
    items: [
      { name: "Clothing", href: "/category/clothing" },
      { name: "Shoes", href: "/category/shoes" },
      { name: "Accessories", href: "/category/accessories" },
      { name: "Watches", href: "/category/watches" },
      { name: "Bags", href: "/category/bags" },
    ],
  },
  {
    name: "Home Appliances",
    items: [
      { name: "Refrigerators", href: "/category/refrigerators" },
      { name: "Microwaves", href: "/category/microwaves" },
      { name: "Washing Machines", href: "/category/washing-machines" },
      { name: "Air Conditioners", href: "/category/air-conditioners" },
      { name: "Vacuum Cleaners", href: "/category/vacuum-cleaners" },
    ],
  },
  {
    name: "Books",
    items: [
      { name: "Fiction", href: "/category/fiction" },
      { name: "Non-Fiction", href: "/category/non-fiction" },
      { name: "Comics", href: "/category/comics" },
      { name: "Biographies", href: "/category/biographies" },
      { name: "Educational", href: "/category/educational" },
    ],
  },
  {
    name: "Sports & Fitness",
    items: [
      { name: "Fitness Equipment", href: "/category/fitness-equipment" },
      { name: "Outdoor Gear", href: "/category/outdoor-gear" },
      { name: "Sportswear", href: "/category/sportswear" },
      { name: "Footwear", href: "/category/footwear" },
      { name: "Accessories", href: "/category/accessories" },
    ],
  },
];

export default function Home() {
  const [openIndex, setOpenIndex] = useState(null);

  const toggleDropdown = (index) => {
    setOpenIndex(openIndex === index ? null : index);
  };

  const closeDropdown = () => {
    setOpenIndex(null);
  };

  return (
    <div className="bg-gray-100 py-4">
      <div className="container mx-auto px-6">
        <div className="flex justify-center space-x-6">
          {categories.map((category, index) => (
            <div key={index} className="relative group">
              <button
                className="flex items-center space-x-2 text-gray-700 hover:text-gray-900 text-lg font-medium focus:outline-none"
                onClick={() => toggleDropdown(index)}
                aria-expanded={openIndex === index}
                aria-controls={`dropdown-${index}`}
              >
                <span>{category.name}</span>
                <FiChevronDown
                  className={`text-gray-500 ${
                    openIndex === index ? "text-gray-700" : "group-hover:text-gray-700"
                  } transition-colors`}
                />
              </button>

              {/* Dropdown menu */}
              {openIndex === index && (
                <div
                  id={`dropdown-${index}`}
                  className="absolute left-0 top-full mt-2 bg-white rounded-lg shadow-lg z-10"
                  onMouseLeave={closeDropdown}
                >
                  <ul className="py-2 px-4">
                    {category.items.map((item, i) => (
                      <li
                        key={i}
                        className="text-gray-600 hover:text-gray-800 cursor-pointer py-1 px-2 hover:bg-gray-100 rounded"
                      >
                        <a
                          href={item.href}
                          className="block"
                          onClick={() => closeDropdown()} // Close dropdown on navigation
                        >
                          {item.name}
                        </a>
                      </li>
                    ))}
                  </ul>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

