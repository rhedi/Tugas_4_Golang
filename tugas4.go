package main

import "fmt"

func main()  {
  var tampil_mahasiswa = map[string]string{"Aldo":"182 cm","Yosep":"178 cm"}
  fmt.Println ("Aldo :",tampil_mahasiswa["Aldo"])
  fmt.Println("Yosep : ",tampil_mahasiswa["Yosep"])
}

func tampil_mahasiswa(x string, y string)(string,string)  {
var m = x ;
var m1 = y ;
return m, m1
}
