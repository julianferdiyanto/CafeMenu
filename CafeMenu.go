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

func tampilMenu() {
	fmt.Println("\n=== DAFTAR MENU ===")
	if n == 0 {
		fmt.Println("Belum ada data menu.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Println("----------------------------")
		fmt.Println("ID        :", daftar[i].id)
		fmt.Println("Nama      :", daftar[i].nama)
		fmt.Println("Kategori  :", daftar[i].kategori)
		fmt.Println("Harga     :", daftar[i].harga)
		fmt.Println("Komposisi :", daftar[i].komposisi)
		fmt.Println("Tersedia  :", daftar[i].tersedia)
	}
}

func ubahMenu() {
	var id int
	fmt.Print("Masukkan ID menu yang diubah : ")
	fmt.Scan(&id)
	for i := 0; i < n; i++ {
		if daftar[i].id == id {
			fmt.Print("Nama baru            : ")
			fmt.Scan(&daftar[i].nama)
			fmt.Print("Kategori baru        : ")
			fmt.Scan(&daftar[i].kategori)
			fmt.Print("Harga baru           : ")
			fmt.Scan(&daftar[i].harga)
			fmt.Print("Komposisi baru       : ")
			fmt.Scan(&daftar[i].komposisi)
			fmt.Print("Tersedia (true/false): ")
			fmt.Scan(&daftar[i].tersedia)
			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func hapusMenu() {
	var id int
	fmt.Print("Masukkan ID menu yang dihapus : ")
	fmt.Scan(&id)
	for i := 0; i < n; i++ {
		if daftar[i].id == id {
			for j := i; j < n-1; j++ {
				daftar[j] = daftar[j+1]
			}
			n--
			fmt.Println("Menu berhasil dihapus!")
			return
		}
	}
	fmt.Println("Data tidak ditemukan.")
}

func kelolaKategori() {
	var pilih int
	for {
		fmt.Println("\n=== KELOLA KATEGORI ===")
		fmt.Println("1. Tambah kategori")
		fmt.Println("2. Ubah kategori")
		fmt.Println("3. Hapus kategori")
		fmt.Println("4. Tampilkan kategori")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			fmt.Print("Nama kategori baru : ")
			fmt.Scan(&daftarKategori[nk])
			nk++
			fmt.Println("Kategori berhasil ditambahkan!")
		} else if pilih == 2 {
			var lama string
			fmt.Print("Kategori yang diubah : ")
			fmt.Scan(&lama)
			for k := 0; k < nk; k++ {
				if daftarKategori[k] == lama {
					fmt.Print("Nama baru : ")
					fmt.Scan(&daftarKategori[k])
					fmt.Println("Kategori berhasil diubah!")
				}
			}
		} else if pilih == 3 {
			var nama string
			fmt.Print("Kategori yang dihapus : ")
			fmt.Scan(&nama)
			for k := 0; k < nk; k++ {
				if daftarKategori[k] == nama {
					for j := k; j < nk-1; j++ {
						daftarKategori[j] = daftarKategori[j+1]
					}
					nk--
					fmt.Println("Kategori berhasil dihapus!")
					break
				}
			}
		} else if pilih == 4 {
			fmt.Println("\n=== DAFTAR KATEGORI ===")
			for k := 0; k < nk; k++ {
				fmt.Printf("%d. %s\n", k+1, daftarKategori[k])
			}
		} else if pilih == 0 {
			return
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}