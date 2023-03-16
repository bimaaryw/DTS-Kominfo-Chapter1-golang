package main

import "fmt"

type Absen struct {
	Nama	string
	Alamat	string
	Pekerjaan	string
	Alasan	string
}

func main(){
	var absen Absen

//	absen.Nama = "Bima","Ary"
//	absen.Alamat = "Pati","Kayen"
//	absen.Pekerjaan= "mahasiswa","pelajar"
//	absen.Alasan ="sertifikat","ujian"

//	fmt.Println(absen.Nama)
//	fmt.Println(absen.Alamat)
//	fmt.Println(absen.Pekerjaan)
//	fmt.Println(absen.Alasan)

fmt.Println(absen.Nama)
fmt.Println(absen.Alamat)
fmt.Println(absen.Pekerjaan)
fmt.Println(absen.Alasan)

}

