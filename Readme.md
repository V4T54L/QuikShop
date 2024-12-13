# 🚀 QuikShop

**QuikShop** is a modern eCommerce platform designed to deliver a fast and seamless shopping experience. Built with a robust tech stack, this application provides a scalable solution for both small businesses and large retailers.

## 🌟 Features

- **Instant Checkout:** Streamlined checkout process ensures quick transactions.
- **Responsive Design:** Beautifully crafted with Tailwind CSS for optimal display on any device.
- **Dynamic Product Management:** Effortlessly manage product listings with rich metadata.
- **Advanced Search & Filtering:** Help customers find products effortlessly.
- **Secure Payments:** Integration with various payment gateways for secure transactions.
- **User Reviews & Ratings:** Foster customer feedback to build trust and improve product visibility.

## 🛠️ Tech Stack

- **Frontend:** 
  - React
  - Redux
  - Tailwind CSS
- **Backend:**
  - Golang
  - PostgreSQL
  - Redis
- **State Management:** Redux

## 💻 Getting Started

### 📋 Prerequisites

Make sure you have the following installed on your machine:

- [Node.js](https://nodejs.org) (v14.x or later)
- [Go](https://golang.org/dl/) (v1.16 or later)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/download)

### 🛠️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/V4T54L/QuikShop.git
   cd quikshop
   ```

2. **Set up the frontend:**

   ```bash
   cd frontend
   npm install
   ```

3. **Set up the backend:**

   ```bash
   cd backend
   go mod tidy
   ```

4. **Configure the environment variables:**

   Create a `.env` file in the backend directory copy the contents of .env.example into it.

5. **Run the database migrations:**

   Make sure your PostgreSQL server is running and execute the migration commands from the backend.

6. **Start the servers:**

   - **Frontend:**

     ```bash
     cd frontend
     npm start
     ```

   - **Backend:**

     ```bash
     cd backend
     go run main.go
     ```

### 🧪 Running Tests

- Run frontend tests:

  ```bash
  cd frontend
  npm test
  ```

- Run backend tests:

  ```bash
  cd backend
  go test ./...
  ```
---

