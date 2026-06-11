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

func tambahMenu(T *tabMenu, n *int) {
	if *n >= NMAX {
		fmt.Println("Data menu sudah penuh.")
		return
	}
	var baru Menu
	var s int
	fmt.Print("Nama menu (tanpa spasi)      : ")
	fmt.Scan(&baru.nama)
	fmt.Print("Kategori (mis. coffee)       : ")
	fmt.Scan(&baru.kategori)
	fmt.Print("Harga                        : ")
	fmt.Scan(&baru.harga)
	fmt.Print("Komposisi (tanpa spasi)      : ")
	fmt.Scan(&baru.komposisi)
	fmt.Print("Status (1 = Tersedia, 0 = Habis) : ")
	fmt.Scan(&s)
	if s == 1 {
		baru.tersedia = true
	} else {
		baru.tersedia = false
	}
	T[*n] = baru
	*n = *n + 1
	fmt.Println("Menu berhasil ditambahkan.")
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