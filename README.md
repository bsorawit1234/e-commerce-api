# e-commerce-api

### Project Structure

```
simple-ecommerce/
├── main.go
├── controllers/
│   ├── orderController.go
│   └── productController.go
├── models/
│   ├── order.go
│   └── product.go
├── database/
│   └── database.go
├── go.mod
```

---

### Features

-   Product Management
    -   Create, Read, Update, Delete (CRUD) operations for products.
-   Order Processing
    -   Create orders with one or more products.
-   Inventory Management
    -   Update product stock levels upon order creation.
-   Data Validation and Error Handling
    -   Robust input validation and meaningful error messages.

---

### API Endpoints

Base URL

    http://localhost:8080

Products

1.  Get All Products
    -   Endpoint: GET /products
    -   Description: Retrieve a list of all products.
2.  Get Product by ID
    -   Endpoint: GET /products/:id
    -   Description: Retrieve a single product by its ID.
    -   Parameters:
        -   id: Product ID
3.  Create a New Product
    -   Endpoint: POST /products
    -   Description: Create a new product.
    -   Request Body:
        ```
        {
            "name": "Product Name",
            "description": "Product Description",
            "price": 100,
            "stock": 50
        }
        ```
4.  Update a Product

    -   Endpoint: PUT /products/:id
    -   Description: Update an existing product.
    -   Parameters:
        -   id: Product ID
    -   Request Body: (Any fields to update)
        ```
        {
            "price": 120,
            "stock": 45
        }
        ```

5.  Delete a Product
    -   Endpoint: DELETE /products/:id
    -   Description: Delete a product by its ID.
    -   Parameters: - id: Product ID
        Orders
6.  Create an Order

    -   Endpoint: POST /orders
    -   Description: Create a new order with one or more products.
    -   Request Body:

        ```
        {
            "items": [
                {
                    "product_id": 1,
                    "quantity": 2
                },
                {
                    "product_id": 3,
                    "quantity": 1
                }
            ]
        }
        ```

7.  Get All Orders
    -   Endpoint: GET /orders
    -   Description: Retrieve a list of all orders.
8.  Get Order by ID
    -   Endpoint: GET /orders/:id
    -   Description: Retrieve a single order by its ID.
    -   Parameters:
        -   id: Order ID
