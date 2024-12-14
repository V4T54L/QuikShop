package store

import (
	"backend/internals/models"
	"context"
	"math/rand"
)

type product struct {
	id             int
	name           string
	description    string
	price          int
	thumbnail      string
	images         []string
	Specifications map[string]string
	category       string
}

type review struct {
	user    string
	comment string
	rating  int
	images  []string
}

type MockProductStore struct {
	products []product
	reviews  []review
}

func NewMockProductStore() *MockProductStore {
	products := []product{
		{
			id:             1,
			name:           "Wireless Headphones",
			description:    "High-quality noise-cancelling wireless headphones with a long battery life.",
			price:          150,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Battery Life": "20 hours", "Weight": "250g"},
			category:       "Electronics",
		},
		{
			id:             2,
			name:           "Smartwatch",
			description:    "Stylish smartwatch with fitness tracking and customizable watch faces.",
			price:          200,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Water Resistance": "IP68", "Display Size": "1.4 inches"},
			category:       "Wearables",
		},
		{
			id:             3,
			name:           "4K Ultra HD TV",
			description:    "Experience stunning picture quality with this 4K Ultra HD Smart TV.",
			price:          700,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Screen Size": "55 inches", "HDR": "Yes"},
			category:       "Home Appliances",
		},
		{
			id:             4,
			name:           "Gaming Laptop",
			description:    "Powerful gaming laptop with high performance graphics and fast processing speed.",
			price:          1200,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"RAM": "16GB", "Storage": "1TB SSD"},
			category:       "Computers",
		},
		{
			id:             5,
			name:           "Electric Kettle",
			description:    "Quick boil electric kettle with auto shut-off feature.",
			price:          50,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Capacity": "1.7L", "Material": "Stainless Steel"},
			category:       "Kitchen Appliances",
		},
		{
			id:             6,
			name:           "Bluetooth Speaker",
			description:    "Portable Bluetooth speaker with superior sound quality and long battery life.",
			price:          80,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Battery Life": "15 hours", "Waterproof": "Yes"},
			category:       "Audio",
		},
		{
			id:             7,
			name:           "Air Fryer",
			description:    "Healthy air fryer that cooks food with little to no oil.",
			price:          120,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Capacity": "4L", "Power": "1500W"},
			category:       "Kitchen Appliances",
		},
		{
			id:             8,
			name:           "Robot Vacuum",
			description:    "Intelligent robot vacuum cleaner with mapping technology.",
			price:          300,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Battery": "2600mAh", "Cleaning Modes": "Multiple"},
			category:       "Home Appliances",
		},
		{
			id:             9,
			name:           "Fitness Tracker",
			description:    "Accurate fitness tracker with heart rate monitoring and sleep tracking.",
			price:          75,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Battery Life": "7 days", "Compatibility": "iOS & Android"},
			category:       "Wearables",
		},
		{
			id:             10,
			name:           "Portable SSD",
			description:    "Fast portable SSD for secure data storage on the go.",
			price:          100,
			thumbnail:      "https://via.placeholder.com/150",
			images:         []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
			Specifications: map[string]string{"Storage Capacity": "500GB", "Transfer Speed": "1000MB/s"},
			category:       "Computers",
		},
	}

	reviews := []review{
		{
			user:    "Alice",
			comment: "Amazing sound quality and battery life!",
			rating:  5,
			images:  []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
		},
		{
			user:    "Bob",
			comment: "Good fitness features but a bit bulky.",
			rating:  3,
			images:  []string{"https://via.placeholder.com/600x400"},
		},
		{
			user:    "Charlie",
			comment: "Love the picture quality and smart features!",
			rating:  4,
			images:  []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
		},
		{
			user:    "Diana",
			comment: "Great for gaming but quite heavy.",
			rating:  4,
			images:  []string{},
		},
		{
			user:    "Ethan",
			comment: "Works well, makes cooking easier!",
			rating:  5,
			images:  []string{"https://via.placeholder.com/600x400"},
		},
		{
			user:    "Fiona",
			comment: "Very convenient for outdoor use, love it!",
			rating:  5,
			images:  []string{"https://via.placeholder.com/600x400", "https://via.placeholder.com/600x400"},
		},
		{
			user:    "George",
			comment: "Vacuums well but gets stuck often.",
			rating:  3,
			images:  []string{},
		},
		{
			user:    "Hannah",
			comment: "Just what I needed for data backup.",
			rating:  4,
			images:  []string{"https://via.placeholder.com/600x400"},
		},
		{
			user:    "Ian",
			comment: "Battery lasts long and has nice features.",
			rating:  4,
			images:  []string{},
		},
		{
			user:    "Jasmine",
			comment: "Comfortable and stylish! Highly recommend.",
			rating:  5,
			images:  []string{"https://via.placeholder.com/600x400"},
		},
	}

	return &MockProductStore{products: products, reviews: reviews}
}

func (s *MockProductStore) getProductSummmary(idx int) *models.ProductSummary {
	product := models.ProductSummary{}
	productData := &s.products[idx]

	product.ID = productData.id
	product.Name = productData.name
	product.Price = productData.price
	product.ThumbnailURI = productData.thumbnail
	product.Rating = (rand.Int() % 5) * 100
	product.Category = productData.category

	return &product
}

func (s *MockProductStore) getProductDetail(idx, nReviews int) *models.ProductDetail {
	product := models.ProductDetail{}
	productData := &s.products[idx]
	product.ID = productData.id
	product.Name = productData.name
	product.Description = productData.description
	product.Price = productData.price
	product.ImageURIs = productData.images
	product.Specifications = productData.Specifications
	product.Category = productData.category

	reviews := make([]models.Review, nReviews)
	rating := 0
	for i := 0; i < nReviews; i++ {
		n := rand.Int() % len(s.reviews)
		reviewData := s.reviews[n]
		reviews[i].User = reviewData.user
		reviews[i].Comment = reviewData.comment
		reviews[i].Rating = reviewData.rating
		reviews[i].ImageURIs = reviewData.images
		rating += reviewData.rating
	}

	product.Reviews = reviews
	product.Rating = rating * 100 / len(reviews)

	return &product
}

// func (s *MockProductStore) GetFeaturedProducts(ctx context.Context) ([]models.ProductSummary, error) {
// 	nProducts := 15
// 	products := make([]models.ProductSummary, 0, nProducts)
// 	for i := 0; i < nProducts; i++ {
// 		products = append(products, *s.getProductSummmary(rand.Int() % len(s.products)))
// 	}
// 	return products, nil
// }

func (s *MockProductStore) SearchProducts(ctx context.Context, query string, start, limit int) ([]models.ProductSummary, error) {
	products := make([]models.ProductSummary, 0, limit)
	for i := 0; i < limit; i++ {
		products = append(products, *s.getProductSummmary(rand.Int() % len(s.products)))
	}
	return products, nil
}

func (s *MockProductStore) GetProductByID(ctx context.Context, productID int) (*models.ProductDetail, error) {
	return s.getProductDetail(productID, rand.Int()%20), nil
}

func (s *MockProductStore) GetProductsByIDs(ctx context.Context, ids []int) ([]models.ProductSummary, error) {
	products := make([]models.ProductSummary, 0, len(ids))
	for _, productId := range ids {
		products = append(products, *s.getProductSummmary(productId))
	}
	return products, nil
}
