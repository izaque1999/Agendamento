package main

func main() {

	fmt.Println("rodando a api")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
