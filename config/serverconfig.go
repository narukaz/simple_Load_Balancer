package config

type Server struct {
	Host string
	Port string
}

var ServerSlice = []Server{
	{Host: "localhost", Port: "8080"},
	{Host: "localhost", Port: "9090"},
	{Host: "localhost", Port: "4090"},
	{Host: "localhost", Port: "9080"},
}
