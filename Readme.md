# Aplikasi Antrian Rumah Sakit

Aplikasi Antrian Rumah Sakit adalah aplikasi sederhana yang dibuat menggunakan bahasa pemrograman Go. Aplikasi ini membantu dalam manajemen antrian pasien di rumah sakit atau klinik.

## Fitur

1. **Daftar Pasien**
   - Menambahkan pasien baru ke dalam antrian.
2. **Daftar Dokter**
   - Menambahkan dokter baru ke dalam daftar dokter.
3. **Tampilkan Antrian Pasien**
   - Menampilkan daftar semua pasien yang ada dalam antrian.
4. **Panggil Pasien**
   - Mengeluarkan pasien dari antrian dan memanggil pasien.
5. **Keluar**
   - Keluar dari aplikasi.

## Prasyarat

Pastikan Anda telah menginstal Go di sistem Anda. Anda dapat mengunduh dan menginstal Go dari [golang.org](https://golang.org/).

## Cara Menggunakan

1. **Clone repository ini:**
  ```sh
git clone https://github.com/guns-joy/hospital_queue.git
```
2. **Masuk direktori**
```sh
cd hospital_queue
```

3. **Jalankan program**
   ```sh
   go run hospital_queue.go
   ```
   
## Preview menu 

<pre>
   Sistem Antrian Rumah Sakit
1. Daftar Pasien
2. Daftar Dokter
3. Tampilkan Antrian Pasien
4. Panggil Pasien
5. Keluar
Masukkan pilihan: 
</pre>

## Penjelasan kode 

<pre>
Penjelasan Kode:
   
1. Struktur Pasien: Menyimpan data pasien, termasuk ID dan nama.
   
2. Struktur Dokter: Menyimpan data dokter, termasuk ID dan nama.
   
3. Daftar pasien dan dokter: Menggunakan package container/list untuk antrian pasien dan slice untuk daftar dokter.
   
4. Fungsi daftarPasien: Menambahkan pasien ke dalam antrian. 

5. Fungsi daftarDokter: Menambahkan dokter ke dalam daftar.
   
6. Fungsi tampilkanAntrian: Menampilkan semua pasien dalam antrian.
   
7. Fungsi panggilPasien: Mengeluarkan pasien dari antrian dan memanggil pasien.
   
8. Fungsi main: Menampilkan menu utama dan memproses pilihan pengguna.
</pre>


## Struktur program 

<pre>
 Program ini menggunakan struktur data sederhana untuk menyimpan informasi pasien dan dokter. Berikut adalah penjelasan singkat dari setiap bagian program:

1. Struktur Pasien: Menyimpan data pasien, termasuk ID dan nama.
   
2. Struktur Dokter: Menyimpan data dokter, termasuk ID dan nama.
   
3. Daftar pasien dan dokter: Menggunakan package container/list untuk antrian pasien dan slice untuk daftar dokter.

4. Fungsi daftarPasien: Menambahkan pasien ke dalam antrian.
   
5. Fungsi daftarDokter: Menambahkan dokter ke dalam daftar.
   
6. Fungsi tampilkanAntrian: Menampilkan semua pasien dalam antrian.
   
7. Fungsi panggilPasien: Mengeluarkan pasien dari antrian dan memanggil pasien.
   
8. Fungsi main: Menampilkan menu utama dan memproses pilihan pengguna.
</pre>
