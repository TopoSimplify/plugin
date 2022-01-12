package geometry

//checkError panics on error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
