package main
import "fmt"

const NMAX int = 100
type menu struct {
	nama      string
	kategori  string
	harga     int
	komposisi string
	tersedia  bool
}
type tabMenu [NMAX]menu
type tabKategori [NMAX]string

func tambahMenu(T *tabMenu, n *int) {
	if *n >= NMAX {
		fmt.Println("Data menu sudah penuh.")
		return
	}
	var baru menu
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