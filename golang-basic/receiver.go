package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// value receiver
func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

// pointer receiver
func (p *Person) HaveBirthday() {
	p.Age += 1
}

func main() {
	p := Person{Name: "John", Age: 30}
	p.SayHello() // Output: Hello, my name is John
	p.HaveBirthday()
	fmt.Println(p.Age) // Output: 31

}

/*
---------------------- Explanation :

receiver adalah bagian dari deklarasi metode yang mengikat sebuah fungsi ke tipe tertentu (biasanya struct).
Ini memungkinkan fungsi tersebut dipanggil sebagai metode dari objek bertipe tersebut
mirip seperti method pada objek dalam OOP (Object-Oriented Programming).


- (r ReceiverType): bagian ini disebut receiver. Ini mirip dengan this dalam bahasa lain, tetapi eksplisit di Go.
- r: adalah nama variabel receiver (bisa kamu namai bebas, seperti t, p, s, dst.).
- ReceiverType: adalah tipe dari receiver, biasanya sebuah struct.
- MethodName: nama dari metode.
- args...: parameter fungsi.
- returnType: tipe nilai yang dikembalikan (opsional, tergantung fungsinya).

---- Summary :

Receiver dalam Go adalah cara untuk mengaitkan fungsi dengan tipe tertentu sehingga fungsi itu menjadi method.
Ini adalah dasar dari bagaimana Go mendukung prinsip object-oriented programming tanpa pewarisan eksplisit.
*/
