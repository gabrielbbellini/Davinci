package settings

// Server settings
type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// GetDomain return the server domain URL.
func (server *Server) GetDomain() string {
	return server.Host + ":" + server.Port
}

// Authentication settings.
type Authentication struct {
	JWTSecretKey string `yaml:"JWTSecretKey"`
}

// DataBase settings.
type DataBase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// GetDBSource return the database driver connection source.
func (database *DataBase) GetDBSource() string {
	credentials := database.User + ":" + database.Password
	connection := "@tcp(" + database.Host + ":" + database.Port + ")"
	return credentials + connection + "/" + database.Name + "?parseTime=true"
}

// Settings Davinci project settings.
type Settings struct {
	Server
	DataBase
	Authentication
}
