package database

type Config struct {
	Driver         string `yaml:"driver"`
	Addr           string `yaml:"addr"`
	Port           string `yaml:"port"`
	DB             string `yaml:"db"`
	UserEnvKey     string `yaml:"user_env_key"`
	PasswordEnvKey string `yaml:"password_env_key"`
}
