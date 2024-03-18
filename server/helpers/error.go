package helpers

func ThrowIfError(err error) {
	if err != nil {
		panic(err)
	}
}
