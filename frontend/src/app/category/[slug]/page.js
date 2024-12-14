"use client"; // Required for client-side features like useState
import React, { useState, useEffect } from "react";

export default function CategoryPage({ params }) {
  const { slug } = React.use(params); // Unwrap params using React.use()

  const fetchCategoryData = async (slug, page = 1) => {
    const limit = 5; // Number of products per page
    const offset = (page - 1) * limit; // Calculate offset based on the current page
    const response = await fetch(
      `http://192.168.1.3:8080/products?limit=${limit}`
    );

    console.log("Response: ", response);
    if (!response.ok) {
      throw new Error("Failed to fetch data");
    }
    return response.json();
  };

  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const data = await fetchCategoryData(slug, currentPage);
        // if (data && data.products && data.products.length > 0) {
        //   setProducts((prev) => [...prev, ...data.products]); // Append new products to the existing list
        // } else {
        //   setHasMore(false); // No more products available
        setProducts(data);
      } catch (error) {
        console.error("Error fetching data: ", error);
      } finally {
        setLoading(false);
      }
    };

    if (slug) {
      fetchData(); // Only fetch data if slug exists
    }
  }, [slug, currentPage]); // Re-run the effect when slug or currentPage changes

  const handleNextPage = () => {
    if (hasMore) {
      setCurrentPage((prevPage) => prevPage + 1); // Increment the page number to load more products
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage((prevPage) => prevPage - 1); // Decrease the page number
    }
  };

  if (loading && currentPage === 1) {
    return <div className="text-center mt-16 text-xl">Loading...</div>;
  }

  console.log(products);

  return (
    <div className="container mx-auto py-16">
      <h1 className="text-4xl font-extrabold mb-8 capitalize text-center text-gray-800 tracking-wide">
        {slug}
      </h1>

      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
        {products?.map((product, idx) => (
          <div
            key={idx}
            className="group relative bg-white border rounded-xl shadow-lg overflow-hidden transition-transform duration-300 transform hover:scale-105 hover:shadow-2xl"
          >
            <a href={`/products/${product.id}`}>
              <div className="overflow-hidden">
                <img
                  src={product.thumbnail}
                  alt={product.name}
                  className="w-full h-48 object-cover transition-transform duration-300 group-hover:scale-110"
                />
              </div>

              <div className="p-4">
                <h2 className="text-lg font-semibold text-gray-800 group-hover:text-gray-900 transition-colors">
                  {product.name}
                </h2>
                <p className="text-gray-500 mt-1 text-sm">
                  {product.description || "Description not available"}
                </p>
                <p className="text-xl font-bold text-gray-800 mt-2 group-hover:text-indigo-600 transition-colors">
                  {product.price}
                </p>
              </div>

              <div className="absolute top-2 right-2 bg-indigo-500 text-white text-xs px-3 py-1 rounded-full shadow-lg group-hover:bg-indigo-600 transition-all">
                {product.category || "Category"}
              </div>
            </a>
          </div>
        ))}
      </div>

      <div className="flex justify-center items-center mt-8 space-x-4">
        <button
          onClick={handlePreviousPage}
          disabled={currentPage === 1}
          className={`px-4 py-2 text-sm font-medium border rounded-lg ${
            currentPage === 1
              ? "bg-gray-200 text-gray-400 cursor-not-allowed"
              : "bg-indigo-500 text-white hover:bg-indigo-600"
          }`}
        >
          Previous
        </button>
        <span className="text-gray-700 text-sm font-medium">
          Page {currentPage}
        </span>
        <button
          onClick={handleNextPage}
          disabled={!hasMore}
          className={`px-4 py-2 text-sm font-medium border rounded-lg ${
            !hasMore
              ? "bg-gray-200 text-gray-400 cursor-not-allowed"
              : "bg-indigo-500 text-white hover:bg-indigo-600"
          }`}
        >
          Next
        </button>
      </div>
    </div>
  );
}
