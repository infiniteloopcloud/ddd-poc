package settings

type Settings struct {
	HttpAddress string
}

func Get() Settings {
	return Settings{
		HttpAddress: ":8080",
	}
}
