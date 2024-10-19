### README: Aplikasi Manajemen Makanan

#### Deskripsi
Aplikasi ini merupakan program manajemen makanan berbasis CLI (Command Line Interface). Pengguna dapat menambahkan, mencari, menghapus, memperbarui, membeli, dan menampilkan menu makanan. Aplikasi ini juga dilengkapi dengan sistem diskon otomatis ketika pembelian melebihi nominal tertentu.

#### Fitur
1. **Tambah Menu**: Menambahkan menu makanan baru ke daftar.
2. **Cari Menu**: Mencari menu berdasarkan nama makanan.
3. **Hapus Menu**: Menghapus menu makanan dari daftar.
4. **Update Menu**: Memperbarui informasi makanan yang ada.
5. **Beli Makanan**: Membeli makanan dan menghitung total harga dengan kemungkinan diskon.
6. **Tampilkan Menu**: Menampilkan daftar makanan yang tersedia.


#### Cara Menggunakan
1. **Clone atau download project**.
2. **Buka terminal** dan arahkan ke direktori project.
3. **Jalankan perintah `go run main.go`** untuk memulai aplikasi.
4. **Pilih opsi dari menu utama** untuk melakukan operasi yang diinginkan.

Berikut deskripsi dari masing-masing menu:

- **Tambah Menu**: Anda akan diminta untuk memasukkan nama, jenis, harga, dan jumlah stok makanan. Menu yang berhasil ditambahkan akan langsung tersimpan.
  
- **Cari Menu**: Masukkan nama makanan yang ingin dicari, kemudian informasi makanan akan ditampilkan jika ditemukan.

- **Hapus Menu**: Masukkan nama makanan yang ingin dihapus, jika ditemukan, makanan tersebut akan dihapus dari daftar menu.

- **Update Menu**: Anda akan diminta untuk memperbarui informasi makanan (nama, jenis, harga, stok). Menu makanan akan diperbarui jika ditemukan.

- **Beli Makanan**: Pilih makanan yang ingin dibeli dengan memasukkan namanya, tentukan jumlah, dan aplikasi akan menghitung total harga termasuk diskon (jika memenuhi syarat).

- **Tampilkan Menu**: Menampilkan semua menu makanan beserta detailnya (status, jumlah, harga, dll).

#### Logika Diskon
- **Diskon 30%** jika total pembelian mencapai **Rp500.000** atau lebih.
- **Diskon 10%** jika total pembelian mencapai **Rp200.000** dengan pembelian lebih dari 2 item.
- PPN (Pajak Pertambahan Nilai) sebesar 11% ditambahkan setelah diskon diterapkan.

#### Flowchart
```
                  +-----------------------------------------------------------+
                  |                       Mulai Program                       |
                  +-----------------------------------------------------------+
                                    |                                |
                                    V                                |
                       +------------------------+                    V
                 +---->| Tampilkan Menu Utama   | ----------> Apakah Pilihan Benar? -------+
                 |     +------------------------+                    |                     |
                 |                   |                               |                Tidak Valid
                 |                   V                               V                     |
                 |      +-------------------------+       +---------------------+          |
           tidak |      | Input Pilihan Pengguna  |       | Lakukan Aksi Sesuai |          |
                 |      +-------------------------+       | Dengan Pilihan      |<---------+
                 |                   |                    +---------------------+
                 |                   V                               |
                 |     +---------------------------------+           V
                 +-----| Apakah Pilihan Keluar?          |<----------+
                       +---------------------------------+                
                                    | Ya                                  
                                    V
                  +----------------------------------------------------------+
                  |                        Akhir Program                     |
                  +----------------------------------------------------------+
```

#### Penjelasan Flowchart
1. Program dimulai dengan menampilkan menu utama yang berisi beberapa pilihan seperti tambah menu, cari menu, beli makanan, dll.
2. Pengguna memasukkan input pilihan.
3. Jika input valid, aksi yang sesuai dengan pilihan akan dilakukan.
4. Jika pengguna memilih untuk keluar, program akan berhenti. Jika tidak, program akan kembali ke menu utama.

#### Dependencies
- **GoLang**: Pastikan Go sudah terinstall di sistem Anda (versi yang direkomendasikan minimal 1.18).


---
