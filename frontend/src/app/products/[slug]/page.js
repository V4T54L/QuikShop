"use client";

import React, { useEffect, useState } from "react";

async function fetchProductDetails(id) {
  const response = await fetch(`http://192.168.1.3:8080/products/${id}`);
  if (!response.ok) {
    throw new Error("Failed to fetch product details");
  }
  return response.json();
}

export default function ProductDetailsPage({ params }) {
  const { slug } = React.use(params); // Capture the ID from the URL
  const [product, setProduct] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentImageIndex, setCurrentImageIndex] = useState(0);

  useEffect(() => {
    const getProductDetails = async () => {
      try {
        const data = await fetchProductDetails(slug);
        setProduct(data);
      } catch (error) {
        setError("Failed to load product details.");
      } finally {
        setLoading(false);
      }
    };

    getProductDetails();
  }, [slug]); // Run the effect whenever the product ID changes

  const handleNextImage = () => {
    if (currentImageIndex < product.images.length - 1) {
      setCurrentImageIndex((prevIndex) => prevIndex + 1);
    }
  };

  const handlePrevImage = () => {
    if (currentImageIndex > 0) {
      setCurrentImageIndex((prevIndex) => prevIndex - 1);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div className="container mx-auto p-16 bg-black text-white">
      {/* Product Name & Description */}
      <div className="text-center">
        <h1 className="text-4xl font-extrabold mb-6 text-white">
          {product.name}
        </h1>
        <p className="text-lg text-gray-400 mb-8">{product.description}</p>
      </div>

      {/* Image Carousel */}
      <div className="relative w-full max-w-2xl mx-auto mb-8">
        {/* Carousel Images */}
        <div className="overflow-hidden rounded-lg shadow-lg">
          <img
            src={product.images[currentImageIndex]}
            alt={`${product.name} image ${currentImageIndex + 1}`}
            className="w-full h-64 object-cover transition-all duration-300"
          />
        </div>

        {/* Navigation Buttons */}
        <button
          onClick={handlePrevImage}
          disabled={currentImageIndex === 0}
          className="absolute left-4 top-1/2 transform -translate-y-1/2 text-white bg-black p-2 rounded-full opacity-75 hover:opacity-100 transition-all disabled:opacity-50"
        >
          &lt;
        </button>
        <button
          onClick={handleNextImage}
          disabled={currentImageIndex === product.images.length - 1}
          className="absolute right-4 top-1/2 transform -translate-y-1/2 text-white bg-black p-2 rounded-full opacity-75 hover:opacity-100 transition-all disabled:opacity-50"
        >
          &gt;
        </button>

        {/* Image Navigation Dots */}
        <div className="absolute bottom-4 left-1/2 transform -translate-x-1/2 flex space-x-2">
          {product.images.map((_, idx) => (
            <span
              key={idx}
              onClick={() => setCurrentImageIndex(idx)}
              className={`w-3 h-3 rounded-full cursor-pointer ${
                idx === currentImageIndex ? "bg-white" : "bg-gray-500"
              }`}
            />
          ))}
        </div>
      </div>

      {/* Product Price */}
      <div className="text-center mb-8">
        <p className="text-3xl font-bold text-white">${product.price}</p>
      </div>

      {/* Specifications */}
      <div className="bg-gray-900 p-6 rounded-lg shadow-md mb-8">
        <h3 className="text-xl font-semibold mb-4 text-white">
          Specifications
        </h3>
        <ul className="space-y-2 text-gray-400">
          {Object.entries(product.specifications).map(([key, value], idx) => (
            <li key={idx} className="flex justify-between">
              <span className="font-medium">{key}</span>
              <span>{value}</span>
            </li>
          ))}
        </ul>
      </div>

      {/* Reviews Section */}
      <div>
        <h3 className="text-2xl font-semibold mb-6 text-white">Reviews</h3>
        {product.reviews.length > 0 ? (
          product.reviews.map((review, idx) => (
            <div
              key={idx}
              className="border border-gray-700 p-6 rounded-lg shadow-md mb-6 bg-gray-800 hover:bg-gray-700 transition-all duration-300"
            >
              <div className="flex items-center space-x-2">
                <p className="text-lg text-white font-semibold">
                  {review.user}
                </p>
                <div className="flex space-x-1 text-yellow-400">
                  {[...Array(5)].map((_, i) => (
                    <svg
                      key={i}
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      fill={i < review.rating ? "currentColor" : "none"}
                      className="w-5 h-5"
                    >
                      <path
                        d="M12 17.27L18.18 21 16.54 13.97 22 9.24l-6.91-.58L12 2 9.91 8.66 3 9.24l5.46 4.73L5.82 21z"
                        stroke={i < review.rating ? "none" : "currentColor"}
                        strokeWidth="2"
                      />
                    </svg>
                  ))}
                </div>
              </div>

              <p className="text-sm text-gray-300 mt-2">{review.comment}</p>

              {/* Review Images */}
              {review.images && review.images.length > 0 && (
                <div className="flex space-x-4 mt-4">
                  {review.images.map((img, imgIdx) => (
                    <img
                      key={imgIdx}
                      src={img}
                      alt={`Review image ${imgIdx + 1}`}
                      className="w-32 h-32 object-cover rounded-lg shadow-md"
                    />
                  ))}
                </div>
              )}
            </div>
          ))
        ) : (
          <p className="text-gray-400">No reviews available.</p>
        )}
      </div>
    </div>
  );
}
