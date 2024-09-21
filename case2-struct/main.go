// Case 3: Book Store

package main

import "fmt"

// Sebuah transaksi di toko buku akan disimpan ke dalam database.
// Untuk itu, perlu dibuat model dari transaksi menggunakan struct.
// Setiap transaksi memiliki beberapa field, yaitu:
// - Total: total harga yang perlu dibayar
// - Date: tanggal terjadinya transaksi berupa string
// - Books: buku-buku yang dibeli oleh pembeli

// Setiap buku juga memiliki beberapa field, yaitu:
// - judul buku
// - nama penulis buku
// - harga buku
// - banyak buku yang dibeli oleh pembeli

// Pada tanggal 23-01-2021, terjadi transaksi dengan total harga sebesar
// 32000. Pembeli membeli 2 buah buku. Buku pertama berjudul "Laut Bercerita"
// karangan Leila S. Chudori seharga 108000 sebanyak 1 buah. Buku berikutnya
// berjudul "Bumi" karangan Tere Liye seharga 97000 sebanyak 1 buah.

type Books struct{
	Title string
	Writer string
	Price int
	Count int
}
type Transaction struct{
	Total int
	Date string
	Books []Books
}

func main() {
	// Modifikasi pendefinisian variable tx1 sehingga sesuai dengan kasus
	// yang diceritakan.
	tx1 := Transaction{
		2, 
		"23-01-2021", 
		[]Books{
			{
				"Laut Bercerita", 
				"Leila S. Chudori", 
				108000, 
				1,
			},
			{
				"Bumi", 
				"Tere Liye", 
				97000, 
				1,
			},
		},
	}

	fmt.Println(tx1)
}
