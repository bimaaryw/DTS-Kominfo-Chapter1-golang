package main

import "fmt"
import "os"
import "strconv"

type Absen struct {
	Nama,Alamat,Pekerjaan,Alasan	string
}

func main(){
	
	var absen = []Absen{
		{Nama: "Bima", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Sertifikat"},
		{Nama: "Ary", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Mencari Ilmu"},
		{Nama: "Widiatmaja", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Mencari teman baru"},
	}

	//fmt.Println(absen[0])
	//fmt.Println(absen2)
	
	//for  _, v := range absen {
	//	fmt.Printf("%v \n", v)
	//}

	arg, _ := strconv.Atoi(os.Args[1])

	fmt.Println("Nama : ", absen[arg].Nama)
	fmt.Println("Alamat : ", absen[arg].Alamat)
	fmt.Println("Pekerjaan : ", absen[arg].Pekerjaan)
	fmt.Println("Alasan : ", absen[arg].Alasan)



}