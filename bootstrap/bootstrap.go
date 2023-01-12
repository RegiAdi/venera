package bootstrap

func Run() {
	loadENV()
	connectDB()
}