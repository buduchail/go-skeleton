package app

type (
	Config struct {
		Prefix     string
		Router     string
		Port       int
		CorrID     string
		Profile    string
		LogLevel   string
		LogFile    string
		LogFormat  string
		Repository string
		Dsn        string
	}
)
