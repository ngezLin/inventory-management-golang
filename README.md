1. Create Product
POST http://localhost:8080/products

{
  "name": "Produk 1",
  "description": "produk satu",
  "price": 100000,
  "category": "C001"
}

Hasil
{
    "data": {
        "ID": 1,
        "CreatedAt": "2025-04-12T22:45:01.442+08:00",
        "UpdatedAt": "2025-04-12T22:45:01.442+08:00",
        "DeletedAt": null,
        "name": "Produk 1",
        "description": "produk satu",
        "price": 100000,
        "category": "C001",
        "image_url": ""
    }
}

2. Lihat semua Products
GET http://localhost:8080/products

Hasil
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2025-04-12T15:28:15.824+08:00",
            "UpdatedAt": "2025-04-12T22:46:24.7+08:00",
            "DeletedAt": null,
            "name": "produk 1",
            "description": "produk satu",
            "price": 100000,
            "category": "C001",
            "image_url": ""
        },
        {
            "ID": 2,
            "CreatedAt": "2025-04-12T22:45:01.442+08:00",
            "UpdatedAt": "2025-04-12T22:45:01.442+08:00",
            "DeletedAt": null,
            "name": "Produk 2",
            "description": "produk kedua",
            "price": 100000,
            "category": "C002",
            "image_url": ""
        }
    ]
}

3. Melihat Product Berdasarkan kategori(ID)
GET http://localhost:8080/products/1



4. Melihat Product berdasarkan ID
GET http://localhost:8080/products/category/C002

Hasil
{
    "data": [
        {
            "ID": 3,
            "CreatedAt": "2025-04-12T22:45:01.442+08:00",
            "UpdatedAt": "2025-04-12T22:48:52.874+08:00",
            "DeletedAt": null,
            "name": "Produk 2",
            "description": "produk kedua",
            "price": 100000,
            "category": "C002",
            "image_url": "/uploads/3_download.jpeg"
        }
    ]
}

5. Update barang berdasarkan ID
PUT http://localhost:8080/products/1
{
  "name": "barang satu baru",
  "description": "Mouse hehe",
  "price": 300000,
  "category": "Peripheral"
}

Hasil
{
    "data": {
        "ID": 1,
        "CreatedAt": "2025-04-12T15:28:15.824+08:00",
        "UpdatedAt": "2025-04-12T23:29:53.21+08:00",
        "DeletedAt": null,
        "name": "barang satu baru",
        "description": "Mouse hehe",
        "price": 300000,
        "category": "Peripheral",
        "image_url": "/uploads/1_download.jpeg"
    }
}

6. Delete Product Berdasarkan ID
DEL http://localhost:8080/products/1

Hasil
{
    "message": "Product deleted"
}

7. Upload Gambar ke product berdasarkan ID
POST http://localhost:8080/products/3/image
key: image, value: file

hasil
{
    "image_url": "/uploads/3_download.jpeg",
    "message": "Image uploaded successfully"
}

8. Melihat gambar dari product ID
GET http://localhost:8080/products/3/image

Hasil
{
    "image_url": "/uploads/3_download.jpeg"
}

9. Membuat Inventory
POST http://localhost:8080/inventory
{
  "product_id": 1,
  "quantity": 100,
  "location": "Gudang A"
}

Hasil
{
    "data": {
        "ID": 2,
        "ProductID": 1,
        "Quantity": 100,
        "Location": "Gudang A"
    }
}

10. Melihat inventory dari product id
GET http://localhost:8080/inventory/1

Hasil 
{
    "data": {
        "ID": 1,
        "ProductID": 1,
        "Quantity": 100,
        "Location": "Gudang A"
    }
}

11. Delete Inventory
DELETE http://localhost:8080/inventory/1

Hasil
{
    "message": "Inventory deleted successfully"
}

12. Update Inventory
PUT http://localhost:8080/inventory/1
{
  "product_id": 1,
  "quantity": 110,
  "location": "Gudang A"
}

Hasil
{
    "data": {
        "ID": 2,
        "ProductID": 1,
        "Quantity": 110,
        "Location": "Gudang A"
    }
}

13. Membuat Order
POST http://localhost:8080/orders
{
  "product_id": 1,
  "quantity": 10
}

Hasil
{
    "data": {
        "ID": 1,
        "CreatedAt": "2025-04-12T23:43:52.277+08:00",
        "UpdatedAt": "2025-04-12T23:43:52.277+08:00",
        "DeletedAt": null,
        "product_id": 1,
        "quantity": 10,
        "order_date": "2025-04-12T23:43:52.277+08:00",
        "Product": {
            "ID": 1,
            "CreatedAt": "2025-04-12T23:41:59.319+08:00",
            "UpdatedAt": "2025-04-12T23:42:46.326+08:00",
            "DeletedAt": null,
            "name": "barang satu baru",
            "description": "Mouse hehe",
            "price": 300000,
            "category": "Peripheral",
            "image_url": "/uploads/1_download.jpeg"
        }
    }
}

14. Melihat order berdasarkan order id
GET http://localhost:8080/orders/1

Hasil
{
    "data": {
        "ID": 1,
        "CreatedAt": "2025-04-12T23:43:52.277+08:00",
        "UpdatedAt": "2025-04-12T23:43:52.277+08:00",
        "DeletedAt": null,
        "product_id": 1,
        "quantity": 10,
        "order_date": "2025-04-12T23:43:52.277+08:00",
        "Product": {
            "ID": 1,
            "CreatedAt": "2025-04-12T23:41:59.319+08:00",
            "UpdatedAt": "2025-04-12T23:42:46.326+08:00",
            "DeletedAt": null,
            "name": "barang satu baru",
            "description": "Mouse hehe",
            "price": 300000,
            "category": "Peripheral",
            "image_url": "/uploads/1_download.jpeg"
        }
    }
}

15. Delete order by id
DELETE http://localhost:8080/orders/1

Hasil
{
    "message": "Order deleted successfully"
}

16. Update order by id
PUT http://localhost:8080/orders/2
{
  "product_id": 3,
  "quantity": 10
}

Hasil
{
    "data": {
        "ID": 2,
        "CreatedAt": "2025-04-13T00:15:12.615+08:00",
        "UpdatedAt": "2025-04-13T00:22:09.891+08:00",
        "DeletedAt": null,
        "product_id": 3,
        "quantity": 10,
        "order_date": "2025-04-13T00:15:12.615+08:00",
        "Product": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "name": "",
            "description": "",
            "price": 0,
            "category": "",
            "image_url": ""
        }
    }
}


link postman
https://api122-5422.postman.co/workspace/Dibimbing~b3b8789f-1523-4763-b10f-6c86b392f45d/collection/40938914-4a913581-f5ec-41b8-9317-4288c2faffc2?action=share&creator=40938914