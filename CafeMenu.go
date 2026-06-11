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

func cariMenu() {
	var pilih int
	var kategori string
	fmt.Println("\n=== CARI MENU ===")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&pilih)
	fmt.Print("Masukkan kategori : ")
	fmt.Scan(&kategori)

	if pilih == 1 {
		var ketemu bool
		fmt.Println("\n=== HASIL SEQUENTIAL SEARCH ===")
		for i := 0; i < n; i++ {
			if daftar[i].kategori == kategori {
				fmt.Println("-", daftar[i].nama, "- Rp", daftar[i].harga)
				ketemu = true
			}
		}
		if !ketemu {
			fmt.Println("Menu tidak ditemukan.")
		}
	} else if pilih == 2 {
		var temp Menu
		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				if daftar[i].kategori > daftar[j].kategori {
					temp = daftar[i]
					daftar[i] = daftar[j]
					daftar[j] = temp
				}
			}
		}
		low := 0
		high := n - 1
		var ketemu bool
		fmt.Println("\n=== HASIL BINARY SEARCH ===")
		for low <= high {
			mid := (low + high) / 2
			if daftar[mid].kategori == kategori {
				fmt.Println("-", daftar[mid].nama, "- Rp", daftar[mid].harga)
				ketemu = true
				break
			} else if daftar[mid].kategori < kategori {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		if !ketemu {
			fmt.Println("Menu tidak ditemukan.")
		}
	} else {
		fmt.Println("Metode tidak tersedia.")
	}
}

func urutkanMenu() {
	var pilih int
	fmt.Println("\n=== URUTKAN MENU ===")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		var min int
		var temp Menu
		for i := 0; i < n-1; i++ {
			min = i
			for j := i + 1; j < n; j++ {
				if daftar[j].harga < daftar[min].harga {
					min = j
				}
			}
			temp = daftar[i]
			daftar[i] = daftar[min]
			daftar[min] = temp
		}
		fmt.Println("Data diurutkan dengan Selection Sort.")
		tampilMenu()
	} else if pilih == 2 {
		var temp Menu
		var j int
		for i := 1; i < n; i++ {
			temp = daftar[i]
			j = i - 1
			for j >= 0 && daftar[j].harga > temp.harga {
				daftar[j+1] = daftar[j]
				j--
			}
			daftar[j+1] = temp
		}
		fmt.Println("Data diurutkan dengan Insertion Sort.")
		tampilMenu()
	} else {
		fmt.Println("Metode tidak tersedia.")
	}
}

func statistik() {
	fmt.Println("\n+++  AsharCafe  +++")
	if n == 0 {
		fmt.Println("Belum ada data menu.")
		fmt.Println("+++++++++++++++++++")
		return
	}
	fmt.Println("Jumlah menu per kategori:")
	for k := 0; k < nk; k++ {
		jumlah := 0
		for i := 0; i < n; i++ {
			if daftar[i].kategori == daftarKategori[k] {
				jumlah++
			}
		}
		fmt.Println("-", daftarKategori[k], ":", jumlah, "menu")
	}
	total := 0
	for i := 0; i < n; i++ {
		total += daftar[i].harga
	}
	fmt.Println("Total semua menu     :", n, "menu")
	fmt.Println("Rata-rata harga menu : Rp", total/n)
	fmt.Println("+++++++++++++++++++")
}

func main() {
	isiDataAwal()
	var pilih int
	for {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Tambah menu")
		fmt.Println("2. Tampilkan menu")
		fmt.Println("3. Ubah menu")
		fmt.Println("4. Hapus menu")
		fmt.Println("5. Kelola kategori")
		fmt.Println("6. Cari menu")
		fmt.Println("7. Urutkan menu")
		fmt.Println("8. Statistik")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahMenu()
		} else if pilih == 2 {
			tampilMenu()
		} else if pilih == 3 {
			ubahMenu()
		} else if pilih == 4 {
			hapusMenu()
		} else if pilih == 5 {
			kelolaKategori()
		} else if pilih == 6 {
			cariMenu()
		} else if pilih == 7 {
			urutkanMenu()
		} else if pilih == 8 {
			statistik()
		} else if pilih == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini!")
			break
		} else {
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}