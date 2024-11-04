# Aplikasi Manajemen Perpustakaan
Aplikasi ini adalah sistem manajemen perpustakaan sederhana yang dibangun menggunakan bahasa pemrograman Go. Pengguna dapat menambah, menghapus, dan menampilkan buku yang ada di perpustakaan. Aplikasi ini dirancang untuk memberikan pengalaman pengguna yang interaktif melalui terminal dengan penggunaan warna pada teks untuk meningkatkan antarmuka pengguna.

## Fitur
- **Tambah Buku**: Menambahkan buku baru ke dalam perpustakaan dengan memeriksa duplikasi berdasarkan ISBN.
- **Hapus Buku**: Menghapus buku dari perpustakaan berdasarkan ISBN.
- **Tampilkan Buku**: Menampilkan daftar semua buku yang ada di perpustakaan.
- **Validasi Input**: Memastikan input dari pengguna valid, termasuk panjang dan format yang sesuai.

## Cara Menggunakan
1. Setelah menjalankan aplikasi, Anda akan melihat menu utama dengan pilihan:
```
   **Aplikasi Manajemen Perpustakaan**
   1. Tambah Buku
   2. Hapus Buku
   3. Tampilkan Buku
   4. Keluar
```
2. Masukkan nomor menu yang ingin Anda pilih dan tekan Enter.
3. Ikuti petunjuk untuk menambah atau menghapus buku, termasuk memasukkan judul, penulis, dan ISBN.
4. Aplikasi akan memberikan umpan balik tentang apakah operasi berhasil atau tidak.

## Penanganan Kesalahan
Aplikasi ini menggunakan mekanisme `recover` untuk menangani kesalahan yang tidak terduga. Jika terjadi kesalahan, pesan kesalahan akan ditampilkan dengan warna merah di terminal.