package main

import "fmt"

type Absen struct {
	Nama,Alamat,Pekerjaan,Alasan	string
}

func main(){
	
	//var absen1 = Absen{Nama: "Bima", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Sertifikat"}
	//var absen2 = Absen{Nama: "Ary", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Mencari Ilmu"}

	var absen = []Absen{
		{Nama: "Bima", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Sertifikat"},
		{Nama: "Ary", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Mencari Ilmu"},
		{Nama: "Bima", Alamat: "Pati", Pekerjaan: "Mahasiswa", Alasan: "Sertifikat"},
	}

	//fmt.Println(absen1)
	//fmt.Println(absen2)
	
	for  _, v := range absen {
		fmt.Printf("%v \n", v)
	}



}