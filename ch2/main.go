package main


func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}