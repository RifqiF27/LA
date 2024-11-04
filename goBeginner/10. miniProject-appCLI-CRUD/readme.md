## E-Wallet CLI

E-Wallet CLI adalah aplikasi dompet digital sederhana yang dibangun menggunakan Golang. Aplikasi ini memiliki fitur seperti pengecekan saldo, top-up, transfer, riwayat transaksi, serta pengaturan akun pengguna.

### Fitur Utama

1. **Login & Register** - Autentikasi pengguna.
2. **Check Balance** - Menampilkan saldo saat ini.
3. **Top-Up** - Menambahkan saldo ke akun.
4. **Transfer** - Mengirim saldo ke akun pengguna lain.
5. **Transaction History** - Menampilkan riwayat transaksi pengguna.
6. **Edit Account** - Mengubah data akun seperti nama, nomor telepon, dan kata sandi.
7. **Logout** - Keluar dari sesi pengguna.
8. **Exit** - Keluar dari aplikasi.

### Flowchart

### Flowchart Login dan Register

```plaintext
                                    +-------------------+
                                    |   Start Program   |
                                    +---------+---------+
                                              |
                                     +--------v--------+
                                     |   Login Menu    |
                                     +--------+--------+
                                              |
                              +---------------+----------------+
                              |                                |
                      +-------v-------+                 +------v-------+
                      |   Register    |                 |    Login     |
                      +-------+-------+                 +------^-------+
                              |                                |
                      +-------v--------+          +------------v--------------+
                      | Input New Data |          | Check Username & Password |
                      +-------+--------+          +------------v--------------+
                              |                                |
                      +-------v--------+                +------v-------+
                      | Save New User  |                |  Valid User? |
                      +-------+--------+                +------^-------+
                              |                                |
                      +-------v--------+          +------------v--------------+
                      | Return to Menu |          | Access Main Menu if Valid |
                      +----------------+          |   Show Error if Invalid   |
                                                  +------------v--------------+

```

### Flowchart Main Menu

```plaintext
                                    +-------------------+
                                    |   Start Program   |
                                    +---------+---------+
                                              |
                                     +--------v--------+
                                     |   Main Menu     |
                                     +--------+--------+
                                              |
                                              |
                             +----------------+-----------------+
                             |                                  |
                      +------v--------+                 +-------v--------+
                      | Check Balance |                 |    Top-Up      |
                      +------+--------+                 +-------+--------+
                             |                                  |
                      +------v--------+              +----------v----------+
                      |    Transfer   |              | Transaction History |
                      +-------+-------+              +----------v----------+
                              |                                 |
                      +-------v--------+                 +------v--------+
                      | Edit Account   |                 |    Logout     |
                      +--------+-------+                 +-------+-------+
                               |                                 |
                      +--------v-------+              +----------v----------+
                      |     Exit       |              |   Return to Login   |
                      +----------------+              +---------------------+
                               |         
                      +--------v-------+ 
                      |       End      | 
                      +----------------+ 

                                

                      
```

---

### Penjelasan Alur

1. **Login Menu**:

   - Saat aplikasi dijalankan, pengguna diarahkan ke _Login Menu_ untuk memilih opsi _Login_ atau _Register_.

2. **Register**:

   - Jika pengguna memilih untuk _Register_, maka pengguna akan diminta untuk memasukkan data baru, seperti nama, nomor telepon, dan kata sandi.
   - Setelah data dimasukkan, aplikasi akan menyimpan pengguna baru tersebut.
   - Pengguna kemudian diarahkan kembali ke _Login Menu_.

3. **Login**:
   - Jika pengguna memilih untuk _Login_, aplikasi akan meminta username dan password.
   - Aplikasi kemudian akan memvalidasi username dan password.
   - Jika valid, pengguna diarahkan ke _Main Menu_.
   - Jika tidak valid, pesan kesalahan akan ditampilkan, dan pengguna diminta untuk mencoba lagi.

Alur ini membantu untuk memastikan bahwa hanya pengguna yang telah terdaftar dapat mengakses fitur utama aplikasi, sementara pengguna baru dapat membuat akun terlebih dahulu.

### Petunjuk Penggunaan

1. **Login atau Registrasi**
   - Saat aplikasi dibuka, pengguna diminta untuk login atau membuat akun baru.
2. **Navigasi Menu Utama**

   - Setelah login, pilih salah satu opsi dari menu utama:
     - `1. Check Balance`: Lihat saldo saat ini.
     - `2. Top-Up`: Tambahkan saldo ke akun.
     - `3. Transfer`: Kirim saldo ke akun lain.
     - `4. History Transaction`: Lihat daftar transaksi sebelumnya.
     - `5. Edit Account`: Perbarui nama, nomor telepon, atau kata sandi.
     - `6. Logout`: Keluar dari sesi dan kembali ke layar login.
     - `7. Exit`: Menutup aplikasi. (menyimpan session ketika running app tidak perlu login lagi)

3. **Kesalahan Input**
   - Jika pengguna memasukkan input yang tidak valid, aplikasi akan menampilkan pesan kesalahan dan meminta input ulang.

### Struktur Folder

- `main.go` - File utama yang menjalankan aplikasi.
- `views` - Mengelola antarmuka menu dan interaksi pengguna.
- `services` - Menyimpan logika layanan seperti login, transfer, dan lainnya.
- `models` - Mengatur data dan logika terkait akun, pengguna, dan transaksi.
- `utils` - Utilitas seperti membaca file, validasi input, dan manajemen error.

### Catatan

- Saldo minimum untuk top-up dan transfer adalah 5000.
- Nomor telepon dan nama harus mengikuti format yang sesuai (misalnya, nama tidak mengandung angka).

---
