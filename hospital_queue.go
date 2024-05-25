package main

import (
	"container/list"
	"fmt"
)

// Struktur untuk menyimpan data pasien
type Pasien struct {
	id   int
	nama string
}

// Struktur untuk menyimpan data dokter
type Dokter struct {
	id   int
	nama string
}

// Daftar pasien dan dokter menggunakan package container/list
var pasienQueue = list.New()
var dokterList []Dokter

// Fungsi untuk menambahkan pasien ke dalam antrian
func daftarPasien(id int, nama string) {
	pasienQueue.PushBack(Pasien{id: id, nama: nama})
	fmt.Println("Pasien berhasil didaftarkan:", nama)
}

// Fungsi untuk menambahkan dokter ke dalam daftar dokter
func daftarDokter(id int, nama string) {
	dokterList = append(dokterList, Dokter{id: id, nama: nama})
	fmt.Println("Dokter berhasil didaftarkan:", nama)
}

// Fungsi untuk menampilkan antrian pasien
func tampilkanAntrian() {
	fmt.Println("Antrian Pasien:")
	for e := pasienQueue.Front(); e != nil; e = e.Next() {
		pasien := e.Value.(Pasien)
		fmt.Printf("ID: %d, Nama: %s\n", pasien.id, pasien.nama)
	}
}

// Fungsi untuk mengeluarkan pasien dari antrian
func panggilPasien() {
	if pasienQueue.Len() > 0 {
		front := pasienQueue.Front()
		pasien := front.Value.(Pasien)
		pasienQueue.Remove(front)
		fmt.Println("Memanggil Pasien:", pasien.nama)
	} else {
		fmt.Println("Tidak ada pasien dalam antrian.")
	}
}

func main() {
	var pilihan int

	for {
		// Menampilkan menu utama
		fmt.Println("\nSistem Antrian Rumah Sakit")
		fmt.Println("1. Daftar Pasien")
		fmt.Println("2. Daftar Dokter")
		fmt.Println("3. Tampilkan Antrian Pasien")
		fmt.Println("4. Panggil Pasien")
		fmt.Println("5. Keluar")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var id int
			var nama string
			fmt.Print("Masukkan ID pasien: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan nama pasien: ")
			fmt.Scan(&nama)
			daftarPasien(id, nama)
		case 2:
			var id int
			var nama string
			fmt.Print("Masukkan ID dokter: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan nama dokter: ")
			fmt.Scan(&nama)
			daftarDokter(id, nama)
		case 3:
			tampilkanAntrian()
		case 4:
			panggilPasien()
		case 5:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
