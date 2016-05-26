package conf

type Configuration struct {
	Timeout struct {
		Read  uint8
		Write uint8
	}
	Upstreams []string
	Loggers   struct {
		Console struct {
			Enable bool
		}
		File struct {
			Enable bool
			Level  string
			Path   string
		}
	}
}
