package kernel

func Run() {
	loadEnv()
	NewMongoConnection()
}
