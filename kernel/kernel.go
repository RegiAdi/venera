package kernel

func Run() {
	loadENV()
	NewMongoConnection()
}
