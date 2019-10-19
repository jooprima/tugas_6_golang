package library

import "fmt"

type mahasiswa struct {
	nama string
	umur int
}

func Public() {
	var m1, m2 mahasiswa
	m1.nama = "yohanes"
	m2.nama = "prima"
	m1.umur = 14
	m2.umur = 15

	m1.namasaya()
	m2.namasaya()
}

func (m mahasiswa) namasaya() {
	fmt.Println("nama saya adalah", m.nama)
	fmt.Println("umur saya adalah", m.umur)
	fmt.Println("")
}
