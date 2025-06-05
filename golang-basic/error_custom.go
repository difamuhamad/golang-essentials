package main

import (
	"errors"
	"fmt"
)

// --------------------------- Custom error with struct
type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (v *ValidationError) Error() string { // defined as error
	return v.Message
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string) error {
	if id == "" {
		return &ValidationError{"ID cannot be empty"} // make custom error
	}
	if id == "wrong-data" {
		return &NotFoundError{"data not found"}
	}

	// if ok
	return nil
}

func main() {
	typeErr := SaveData("wrong-data")

	// check using type assertion
	if typeErr != nil {
		if validationErr, ok := typeErr.(*ValidationError); ok {
			fmt.Println("message:", validationErr.Error())

		} else if notFoundErr, ok := typeErr.(*NotFoundError); ok {
			fmt.Println("message:", notFoundErr.Error())

		} else {
			fmt.Println("unknown error")
		}
	} else {
		// if success
		fmt.Println("Data saved.")
	}

	switchErr := SaveData("")

	// check using switch case
	if switchErr != nil {
		switch errorType := switchErr.(type) {

		case *ValidationError:
			fmt.Println("error type:", errorType.Error())

		case *NotFoundError:
			fmt.Println("error type:", errorType.Error())

		default:
			fmt.Println("unknown error")
		}
	}

	// check using errors.Is
	err := SaveData("wrong-data")

	if err != nil {
		// errors.Is( errorValue , target )
		if errors.Is(err, &ValidationError{}) {
			fmt.Println("Validation error:", err.Error())

		} else if errors.Is(err, &NotFoundError{}) {
			fmt.Println("Not found error:", err.Error())

		} else {
			fmt.Println("Unknown error:", err.Error())
		}
	} else {
		fmt.Println("Data saved.")
	}

}

/*
--------------- Explanation :

Di Go, tipe dianggap sebagai error jika mengimplementasikan method:
Error() string

Di sini, kita membuat ValidationError dan NotFoundError memenuhi interface error.
Error() akan mengembalikan isi dari Message.
Jadi ketika kita lakukan fmt.Println(err), Go akan memanggil err.Error() di balik layar.

----
Memanggil SaveData("") (kondisi error).

Type assertion: cara mengecek tipe asli dari error interface.
- validationErr, ok := err.(*ValidationError)
- Jika ok == true, maka error memang bertipe *ValidationError.
- Menampilkan error berdasarkan tipe error-nya.

----
Selanjutnya adalah bentuk type assertion dengan switch, lebih ringkas dibandingkan if-else.

errorType := switchErr.(type) mengekstrak tipe asli dari interface error.
Lalu, dicek case-nya:
*ValidationError → cetak pesan error-nya.
*NotFoundError → cetak pesan error-nya.

----
func (v *ValidationError) Error()

Mengapa menggunakan pointer?

Efisiensi memori
Dengan pointer, kita tidak menyalin seluruh struct setiap kali method dipanggil. Kita hanya melewatkan alamat memori.
Meskipun struct-nya kecil, ini adalah kebiasaan baik untuk konsistensi dan efisiensi.

Konsistensi
Kalau nanti kita tambahkan method lain yang mengubah nilai struct (mutating method),
maka semua method sebaiknya pakai pointer agar tidak campur aduk antara pointer vs value receiver.
Kesesuaian dengan &ValidationError{...}

Ketika kita buat objek &ValidationError{...}, kita menghasilkan pointer ke struct, jadi naturally cocok dengan method receiver yang juga pointer.


----
if validationErr, ok := err.(*ValidationError); ok {

ok itu dibuat untuk menerima hasil dari pengecekan apakah nilai interface tersebut bisa "di-cast" atau "di-convert" ke tipe tertentu.
value akan berisi hasil konversi ke SomeType jika berhasil.

- ok akan bernilai true jika someInterface memang bertipe SomeType (atau bisa di-convert).
- ok akan bernilai false jika tidak bisa, dan value akan berisi nilai nol dari tipe SomeType.
- Ini memungkinkan kamu untuk melakukan pengecekan tipe dengan aman tanpa risiko program panic.

Ringkasnya:
ok adalah flag boolean hasil pengecekan tipe.
Digunakan untuk memastikan apakah type assertion berhasil atau tidak.

*/
