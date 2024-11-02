# Service Inventory CLI

Service Inventory CLI adalah aplikasi Command-Line Interface (CLI) yang membantu dalam mengelola data inventori, transaksi, kategori, dan lokasi penyimpanan barang. Aplikasi ini mendukung fitur login, pengelolaan item, pencatatan transaksi, dan logout.

## Fitur

1. **Login dan Logout**  
   - **`login`**: Endpoint untuk login pengguna.
   - **`logout`**: Endpoint untuk logout pengguna.

2. **Pengelolaan Barang**
   - **`get-item`**: Melihat daftar barang.
   - **`add-item`**: Menambah barang baru.
   - **`update-item`**: Memperbarui stok barang.

3. **Pencatatan Transaksi**
   - **`add-transaction`**: Menambah transaksi keluar/masuk barang dengan informasi waktu, jumlah, dan keterangan.
   - **`get-transaction`**: Melihat daftar transaksi.

4. **Pengelolaan Kategori dan Lokasi**
   - **`add-category`**: Menambah kategori baru untuk barang.
   - **`add-location`**: Menambah lokasi penyimpanan barang.

## Cara Penggunaan

Setelah menjalankan aplikasi, masukkan endpoint yang ingin diakses sesuai daftar di bawah ini:


| Endpoint            | Deskripsi                                     |
|---------------------|-----------------------------------------------|
| **login**           | Login pengguna.                               |
| **logout**          | Logout pengguna.                              |
| **get-item**        | Menampilkan daftar barang.                    |
| **add-item**        | Menambahkan barang baru.                      |
| **update-stock**     | Memperbarui stok barang.                      |
| **add-transaction** | Mencatat transaksi keluar/masuk barang.       |
| **get-transaction** | Melihat daftar transaksi.                     |
| **add-category**    | Menambahkan kategori barang baru.             |
| **add-location**    | Menambahkan lokasi penyimpanan baru.          |

### Contoh Penggunaan

1. **Login**  
   Masukkan `login` untuk melakukan login.

***login untuk admin***
```json
{
  "username": "admin1",
  "password": "hashedpassword1"
}
```
   ***login untuk staff***
```json
{
  "username": "staff1",
  "password": "hashedpassword2"
}
```

2. **Cek Barang**  
   Masukkan `get-item` untuk melihat daftar barang beserta jumlah stok yang tersedia, kemudian ikuti petunjuk untuk mengatur page, filter stok kurang dari 10 dan pencarian berdasarkan nama barang.
```json
{
    "page": 1,
    "limit": 5,
    "filter_stock": false,
    "search_name": ""
}
```

3. **Tambah Item Baru**  
   Masukkan `add-item`, kemudian ikuti petunjuk untuk menambahkan barang ke inventori.
```json
{
	"item_code": "ITEM021",
	"item_name": "New Projector",
	"stock": 20,
	"category_id": 1,
	"location_id": 1
}
```
4. **Update Stok Barang**  
   Masukkan `Update-stock`, kemudian ikuti petunjuk untuk memperbahurui stok barang yang ada di inventori.
```json
{
	"item_code": "ITEM021",
	"stock": 30
}
```

5. **Transaksi Barang Keluar/Masuk**  
   Masukkan `add-transaction` untuk mencatat transaksi keluar atau masuk barang dengan informasi waktu, jumlah, dan keterangan transaksi.

```json
{
    "item_id": 21,
    "transaction_type": "in",
    "quantity": 70,
    "notes": "shipment to branch",
    "user_id": 1
}
```

6. **Cek Trasaksi**  
   Masukkan `get-transaction` untuk melihat daftar transaksi keluar atau masuk barang dengan informasi waktu, jumlah, dan keterangan transaksi. kemudian ikuti petunjuk untuk melakukan pencarian berdasarkan nama barang.
```json
{
    "item_name": ""
}
```
7. **Tambah Kategori Baru**  
   Masukkan `add-category`, kemudian ikuti petunjuk untuk menambahkan kategori ke inventori.
```json
{
    "category_name": "material"
}
```
8. **Tambah Lokasi Baru**  
   Masukkan `add-location`, kemudian ikuti petunjuk untuk menambahkan lokasi ke inventori.
```json
{
    "location_name": "warehouse D"
}
```
9. **Logout**  
   Masukkan `logout` untuk keluar dari sistem.

## Catatan

- Pastikan database PostgreSQL sudah berjalan dan konfigurasi sudah sesuai.