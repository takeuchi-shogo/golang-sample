package main

type Person struct {
	Name        string
	Age         int
	Postcode    *int
	Discription *string
}

func main() {
	p := Person{}

	p.Name = "太郎"
	p.Age = 30
	*p.Postcode = 0002222
	*p.Discription = "私は日本人です"

	message := "私は社会人です"

	*p.Discription = message

}
