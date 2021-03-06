package remote

type Server struct {
	Site     string `yaml:"site"`
	Speed    int    `yaml:"speed"`
	Filling  int    `yaml:"filling"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Use23    bool   `yaml:"use23"`
}

type Config struct {
	Databases []Server `yaml:"database"`
}
