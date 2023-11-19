package main

import "fmt"

// Pertanyaan struct untuk menyimpan informasi pertanyaan dan jawabannya
type Question struct {
	Pertanyaan   string
	Pilihan      []string
	JawabanBenar int
}

func main() {
	fmt.Print("Masukkan Nama Anda : ")

	var userInput string
	fmt.Scanln(&userInput)

	fmt.Printf("Salam Sejahtera %s \n", userInput)
	fmt.Println()
	fmt.Println("Berikut ada beberapa soal yang perlu kamu jawab :")

	// Slice untuk pertanyaan-pertanyaan
	pertanyaan := []Question{
		{
			"Bahasa pemrograman apa yang digunakan untuk perancangan design web?",
			[]string{"C++", "PHP", "JavaScript", "C#"},
			2,
		},
		{
			"Siapa yang pernah menyandang gelar Raja Bajak Laut di One Piece?",
			[]string{"Marshall D Teach", "Gol D Roger", "Monkey D Luffy", "Akagami no Shanks"},
			1,
		},
	}

	// Variabel untuk menyimpan skor dan statistik jawaban
	var scoreUser int
	var jawabanBenar int
	var jawabanSalah int

	// Loop melalui setiap pertanyaan
	for _, q := range pertanyaan {
		var userAnswer int

		fmt.Printf("%s\n", q.Pertanyaan)

		// Menampilkan pilihan jawaban
		for i, pilihan := range q.Pilihan {
			fmt.Printf("%d. %s\n", i, pilihan)
		}

		fmt.Print("Tentukan pilihan jawaban mu dengan mengisikan nomor (0, 1, 2, 3) : ")
		fmt.Scanln(&userAnswer)

		fmt.Println()

		if userAnswer == q.JawabanBenar {
			fmt.Println("Tepat Sekali")
			scoreUser += 50
			jawabanBenar++
		} else {
			fmt.Println("Kurang Tepat")
			jawabanSalah++
		}
	}

	// Menampilkan statistik kuis
	fmt.Println("Statistik Kuis")
	fmt.Printf("Nama          : %s \n", userInput)
	fmt.Printf("Score         : %d \n", scoreUser)
	fmt.Printf("Jawaban Benar : %d \n", jawabanBenar)
	fmt.Printf("Jawaban Salah : %d \n", jawabanSalah)
}