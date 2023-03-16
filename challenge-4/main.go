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

absen := map[int]map[string]string{
    1: {
        "nama":      "Mochamad",
        "alamat":    "Jl.Jombang-Mojokerto No.1",
        "pekerjaan": "Programmer",
        "alasan":    "Ingin belajar lebih banyak tentang Go",
    },
    2: {
        "nama":      "Syahrul",
        "alamat":    "Jl.Jombang-Mojokerto No.2",
        "pekerjaan": "Freelancer",
        "alasan":    "Ingin tau tentang microservice",
    },
    3: {
        "nama":      "Samsudin",
        "alamat":    "Jl.Jombang-Mojokerto No.3",
        "pekerjaan": "Karyawan",
        "alasan":    "Mencari nafkah",
    },
    4: {
        "nama":      "Hayolo",
        "alamat":    "Jl.Jombang-Mojokerto No.4",
        "pekerjaan": "Mahasiswa",
        "alasan":    "Mencari ijazah",
    },
} 

fmt.Println(absen.Nama)
fmt.Println(absen.Alamat)
fmt.Println(absen.Pekerjaan)
fmt.Println(absen.Alasan)

}

