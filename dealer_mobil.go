package main

import (
	"fmt"
	"sort"
)

type Mobil struct {
	Nama      string
	Pabrikan  string
	Tahun     int
	Penjualan int
}

type Pabrikan struct {
	Nama   string
	Mobils []Mobil
}

var pabrikanList []Pabrikan

// Dummy data
func tambahDataDummy() {
	pabrikanList = append(pabrikanList, Pabrikan{
		Nama: "Toyota",
		Mobils: []Mobil{
			{"Corolla", "Toyota", 2020, 5000},
			{"Camry", "Toyota", 2021, 3000},
		},
	})
	pabrikanList = append(pabrikanList, Pabrikan{
		Nama: "Honda",
		Mobils: []Mobil{
			{"Civic", "Honda", 2022, 3500},
			{"Accord", "Honda", 2021, 2500},
		},
	})
	pabrikanList = append(pabrikanList, Pabrikan{
		Nama: "Ford",
		Mobils: []Mobil{
			{"Focus", "Ford", 2020, 1500},
			{"Mustang", "Ford", 2021, 1000},
		},
	})

	fmt.Println("Data dummy berhasil ditambahkan!")
}

// Fungsi Menambah Pabrikan
func tambahPabrikan() {
	var nama string
	fmt.Print("Masukkan nama pabrikan: ")
	fmt.Scanln(&nama)
	pabrikanList = append(pabrikanList, Pabrikan{Nama: nama})
	fmt.Println("Pabrikan berhasil ditambahkan!")
}

// Fungsi Menambah Mobil
func tambahMobil() {
	var namaPabrikan, namaMobil string
	var tahun, penjualan int

	fmt.Print("Masukkan nama pabrikan: ")
	fmt.Scanln(&namaPabrikan)

	for i, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			fmt.Print("Masukkan nama mobil: ")
			fmt.Scanln(&namaMobil)
			fmt.Print("Masukkan tahun: ")
			fmt.Scanln(&tahun)
			fmt.Print("Masukkan jumlah penjualan: ")
			fmt.Scanln(&penjualan)

			pabrikanList[i].Mobils = append(pabrikanList[i].Mobils, Mobil{
				Nama:      namaMobil,
				Pabrikan:  namaPabrikan,
				Tahun:     tahun,
				Penjualan: penjualan,
			})
			fmt.Println("Mobil berhasil ditambahkan!")
			return
		}
	}

	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi Tampilkan Mobil
func tampilkanMobil() {
	var namaPabrikan string
	fmt.Print("Masukkan nama pabrikan: ")
	fmt.Scanln(&namaPabrikan)

	for _, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			fmt.Printf("Mobil dari pabrikan %s:\n", namaPabrikan)
			for _, m := range p.Mobils {
				fmt.Printf("- %s (%d), Penjualan: %d\n", m.Nama, m.Tahun, m.Penjualan)
			}
			return
		}
	}

	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi Total Penjualan Pabrikan
func totalPenjualan(p Pabrikan) int {
	total := 0
	for _, m := range p.Mobils {
		total += m.Penjualan
	}
	return total
}

// Fungsi Edit Pabrikan
func editPabrikan() {
	var namaLama, namaBaru string
	fmt.Print("Masukkan nama pabrikan yang ingin diubah: ")
	fmt.Scanln(&namaLama)

	for i, p := range pabrikanList {
		if p.Nama == namaLama {
			fmt.Print("Masukkan nama baru untuk pabrikan: ")
			fmt.Scanln(&namaBaru)
			pabrikanList[i].Nama = namaBaru
			fmt.Println("Nama pabrikan berhasil diubah!")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi Edit Mobil
func editMobil() {
	var namaPabrikan, namaMobilLama, namaMobilBaru string
	var tahunBaru, penjualanBaru int

	fmt.Print("Masukkan nama pabrikan: ")
	fmt.Scanln(&namaPabrikan)

	for i, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			fmt.Print("Masukkan nama mobil yang ingin diubah: ")
			fmt.Scanln(&namaMobilLama)

			for j, m := range p.Mobils {
				if m.Nama == namaMobilLama {
					fmt.Print("Masukkan nama baru mobil: ")
					fmt.Scanln(&namaMobilBaru)
					fmt.Print("Masukkan tahun baru: ")
					fmt.Scanln(&tahunBaru)
					fmt.Print("Masukkan jumlah penjualan baru: ")
					fmt.Scanln(&penjualanBaru)

					pabrikanList[i].Mobils[j] = Mobil{
						Nama:      namaMobilBaru,
						Pabrikan:  namaPabrikan,
						Tahun:     tahunBaru,
						Penjualan: penjualanBaru,
					}
					fmt.Println("Data mobil berhasil diubah!")
					return
				}
			}
			fmt.Println("Mobil tidak ditemukan.")
			return
		}
	}
	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi untuk Menghapus Pabrikan
func hapusPabrikan() {
	var namaPabrikan string
	fmt.Print("Masukkan nama pabrikan yang ingin dihapus: ")
	fmt.Scanln(&namaPabrikan)

	for i, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			pabrikanList = append(pabrikanList[:i], pabrikanList[i+1:]...)
			fmt.Println("Pabrikan berhasil dihapus!")
			return
		}
	}

	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi untuk Menghapus Mobil
func hapusMobil() {
	var namaPabrikan, namaMobil string
	fmt.Print("Masukkan nama pabrikan mobil yang ingin dihapus: ")
	fmt.Scanln(&namaPabrikan)

	for i, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			fmt.Print("Masukkan nama mobil yang ingin dihapus: ")
			fmt.Scanln(&namaMobil)

			for j, m := range p.Mobils {
				if m.Nama == namaMobil {
					pabrikanList[i].Mobils = append(pabrikanList[i].Mobils[:j], pabrikanList[i].Mobils[j+1:]...)
					fmt.Println("Mobil berhasil dihapus!")
					return
				}
			}
			fmt.Println("Mobil tidak ditemukan.")
			return
		}
	}

	fmt.Println("Pabrikan tidak ditemukan.")
}

// Fungsi Submenu Pabrikan dengan Opsi Hapus
func submenuPabrikan() {
	for {
		fmt.Println("1. Tambah Pabrikan")
		fmt.Println("2. Edit Pabrikan")
		fmt.Println("3. Hapus Pabrikan")  // Tambah opsi hapus pabrikan
		fmt.Println("4. Keluar ke Menu Utama")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahPabrikan()
		case 2:
			editPabrikan()
		case 3:
			hapusPabrikan()  // Menjalankan fungsi hapus pabrikan
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi Submenu Mobil dengan Opsi Hapus
func submenuMobil() {
	for {
		fmt.Println("1. Tambah Mobil")
		fmt.Println("2. Edit Mobil")
		fmt.Println("3. Hapus Mobil")  // Tambah opsi hapus mobil
		fmt.Println("4. Keluar ke Menu Utama")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahMobil()
		case 2:
			editMobil()
		case 3:
			hapusMobil()  // Menjalankan fungsi hapus mobil
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}


// Fungsi Submenu Pencarian Mobil
func submenuPencarianMobil() {
	for {
		fmt.Println("\nSubmenu Pencarian Mobil:")
		fmt.Println("1. Cari Mobil Berdasarkan Nama Pabrikan")
		fmt.Println("2. Cari Mobil Berdasarkan Nama Mobil")
		fmt.Println("3. Kembali ke Menu Utama")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tampilkanMobilBerdasarkanPabrikan()
		case 2:
			tampilkanMobilBerdasarkanNama()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi Tampilkan Mobil Berdasarkan Nama Pabrikan
func tampilkanMobilBerdasarkanPabrikan() {
	var namaPabrikan string
	fmt.Print("Masukkan nama pabrikan untuk pencarian: ")
	fmt.Scanln(&namaPabrikan)

	found := false
	for _, p := range pabrikanList {
		if p.Nama == namaPabrikan {
			fmt.Printf("Mobil dari pabrikan %s:\n", namaPabrikan)
			for _, m := range p.Mobils {
				fmt.Printf("- %s (%d), Penjualan: %d\n", m.Nama, m.Tahun, m.Penjualan)
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Pabrikan tidak ditemukan.")
	}
}

// Fungsi Tampilkan Mobil Berdasarkan Nama Mobil
func tampilkanMobilBerdasarkanNama() {
	var namaMobil string
	fmt.Print("Masukkan nama mobil untuk pencarian: ")
	fmt.Scanln(&namaMobil)

	found := false
	for _, p := range pabrikanList {
		for _, m := range p.Mobils {
			if m.Nama == namaMobil {
				fmt.Printf("Mobil ditemukan: %s (Pabrikan: %s, Tahun: %d, Penjualan: %d)\n", m.Nama, p.Nama, m.Tahun, m.Penjualan)
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		fmt.Println("Mobil tidak ditemukan.")
	}
}

// Fungsi Submenu Penjualan
func submenuPenjualan() {
	for {
		fmt.Println("\nSubmenu Penjualan:")
		fmt.Println("1. Tampilkan Top 3 Mobil dengan Penjualan Tertinggi")
		fmt.Println("2. Tampilkan Top 3 Pabrikan dengan Penjualan Tertinggi")
		fmt.Println("3. Kembali ke Menu Utama")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			top3MobilBerdasarkanPenjualan()
		case 2:
			top3PabrikanBerdasarkanPenjualan()
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi Menampilkan Top 3 Mobil Berdasarkan Penjualan
func top3MobilBerdasarkanPenjualan() {
	// Flatten the list of all cars into a single list
	var allMobils []Mobil
	for _, p := range pabrikanList {
		allMobils = append(allMobils, p.Mobils...)
	}

	// Sort the cars by total sales in descending order
	sort.Slice(allMobils, func(i, j int) bool {
		return allMobils[i].Penjualan > allMobils[j].Penjualan
	})

	// Display the top 3
	fmt.Println("Top 3 Mobil dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < len(allMobils); i++ {
		fmt.Printf("%d. %s (Pabrikan: %s, Penjualan: %d)\n", i+1, allMobils[i].Nama, allMobils[i].Pabrikan, allMobils[i].Penjualan)
	}
}

// Fungsi Menampilkan Top 3 Pabrikan Berdasarkan Penjualan
func top3PabrikanBerdasarkanPenjualan() {
	// Sort pabrikanList by total sales
	sort.SliceStable(pabrikanList, func(i, j int) bool {
		return totalPenjualan(pabrikanList[i]) > totalPenjualan(pabrikanList[j])
	})

	// Display the top 3
	fmt.Println("Top 3 Pabrikan dengan Penjualan Tertinggi:")
	for i := 0; i < 3 && i < len(pabrikanList); i++ {
		fmt.Printf("%d. %s (Total Penjualan: %d)\n", i+1, pabrikanList[i].Nama, totalPenjualan(pabrikanList[i]))
	}
}

// Fungsi Submenu Sorting
func submenuSorting() {
	for {
		fmt.Println("\nSubmenu Sorting:")
		fmt.Println("1. Sort Mobil Berdasarkan Nama")
		fmt.Println("2. Sort Mobil Berdasarkan Penjualan")
		fmt.Println("3. Sort Pabrikan Berdasarkan Nama")
		fmt.Println("4. Sort Pabrikan Berdasarkan Penjualan")
		fmt.Println("5. Kembali ke Menu Utama")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			sort.Slice(pabrikanList, func(i, j int) bool {
				return pabrikanList[i].Nama < pabrikanList[j].Nama
			})
			fmt.Println("Pabrikan telah diurutkan berdasarkan nama.")
		case 2:
			sort.SliceStable(pabrikanList, func(i, j int) bool {
				return totalPenjualan(pabrikanList[i]) > totalPenjualan(pabrikanList[j])
			})
			fmt.Println("Pabrikan telah diurutkan berdasarkan total penjualan.")
		case 3:
			for i := range pabrikanList {
				sort.Slice(pabrikanList[i].Mobils, func(x, y int) bool {
					return pabrikanList[i].Mobils[x].Nama < pabrikanList[i].Mobils[y].Nama
				})
			}
			fmt.Println("Mobil telah diurutkan berdasarkan nama.")
		case 4:
			for i := range pabrikanList {
				sort.SliceStable(pabrikanList[i].Mobils, func(x, y int) bool {
					return pabrikanList[i].Mobils[x].Penjualan > pabrikanList[i].Mobils[y].Penjualan
				})
			}
			fmt.Println("Mobil telah diurutkan berdasarkan penjualan.")
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func printMenuBorder() {
	fmt.Println("\n=====================================")
	fmt.Println("          MENU UTAMA               ")
	fmt.Println("=====================================")
}

// Menu Utama
func menu() {
	// Add dummy data on startup
	tambahDataDummy()

	for {
		printMenuBorder()
		fmt.Println("\nMenu:")
		fmt.Println("1. Kelola Pabrikan")
		fmt.Println("2. Kelola Mobil")
		fmt.Println("3. Pencarian Mobil")
		fmt.Println("4. Submenu Sorting")
		fmt.Println("5. Submenu Penjualan")
		fmt.Println("6. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			submenuPabrikan()
		case 2:
			submenuMobil()
		case 3:
			submenuPencarianMobil()
		case 4:
			submenuSorting()
		case 5:
			submenuPenjualan()
		case 6:
			fmt.Println("Terima kasih! Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fungsi main sebagai entry point
func main() {
	menu()
}
