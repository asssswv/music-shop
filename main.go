package main

func main() {
	router := getRouter()
	_ = router.Run("localhost:8080")
}
