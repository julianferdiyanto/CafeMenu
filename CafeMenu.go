package main

import "fmt"

const NMAX = 100
const KMAX = 50

type Menu struct {
	id        int
	nama      string
	kategori  string
	harga     int
	komposisi string
	tersedia  bool
}

var daftar [NMAX]Menu
var n int
var daftarKategori [KMAX]string
var nk int

func isiDataAwal() {
	daftar[0] = Menu{1, "Americano", "coffee", 18000, "kopi", true}
	daftar[1] = Menu{2, "Latte", "coffee", 25000, "kopi_susu", true}
	daftar[2] = Menu{3, "Cappuccino", "coffee", 28000, "kopi_susu", true}
	daftar[3] = Menu{4, "Matcha", "non-coffee", 22000, "bubuk_matcha", true}
	daftar[4] = Menu{5, "LemonTea", "non-coffee", 12000, "teh_lemon", true}
	n = 5
	daftarKategori[0] = "coffee"
	daftarKategori[1] = "non-coffee"
	nk = 2
}

func tambahMenu() {
	fmt.Println("\n=== TAMBAH MENU ===")
	fmt.Print("ID                   : ")
	fmt.Scan(&daftar[n].id)
	fmt.Print("Nama Menu            : ")
	fmt.Scan(&daftar[n].nama)
	fmt.Print("Kategori             : ")
	fmt.Scan(&daftar[n].kategori)
	fmt.Print("Harga                : ")
	fmt.Scan(&daftar[n].harga)
	fmt.Print("Komposisi            : ")
	fmt.Scan(&daftar[n].komposisi)
	fmt.Print("Tersedia (true/false): ")
	fmt.Scan(&daftar[n].tersedia)
	n++
	fmt.Println("Menu berhasil ditambahkan!")
}

func tampilMenu(T tabMenu, n int) {
	if n == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}
	fmt.Println()
	fmt.Printf("%-3s %-15s %-12s %-9s %-10s %s\n", "No", "Nama", "Kategori", "Harga", "Status", "Komposisi")
	var i int
	for i = 0; i < n; i++ {
		var status string
		if T[i].tersedia {
			status = "Tersedia"
		} else {
			status = "Habis"
		}
		fmt.Printf("%-3d %-15s %-12s %-9d %-10s %s\n", i+1, T[i].nama, T[i].kategori, T[i].harga, status, T[i].komposisi)
	}
	fmt.Println()
}

func ubahMenu(T *tabMenu, n int) {
	if n == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}
	tampilMenu(*T, n)
	var nomor int
	fmt.Print("Pilih nomor menu yang ingin diubah: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > n {
		fmt.Println("Nomor menu tidak valid.")
		return
	}
	var idx int = nomor - 1
	var s int
	fmt.Print("Nama menu baru (tanpa spasi) : ")
	fmt.Scan(&T[idx].nama)
	fmt.Print("Kategori baru                : ")
	fmt.Scan(&T[idx].kategori)
	fmt.Print("Harga baru                   : ")
	fmt.Scan(&T[idx].harga)
	fmt.Print("Komposisi baru (tanpa spasi) : ")
	fmt.Scan(&T[idx].komposisi)
	fmt.Print("Status (1=Tersedia, 0=Habis) : ")
	fmt.Scan(&s)
	if s == 1 {
		T[idx].tersedia = true
	} else {
		T[idx].tersedia = false
	}
	fmt.Println("Menu berhasil diubah.")
}
