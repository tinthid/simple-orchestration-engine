package server

type Server struct {
	RedisConn    string
	RabbitMqConn string
}

func CreateServer(redisConn string, rabbitMqConn string) (s *Server) {
	s = new(Server)
	s.RedisConn = redisConn
	s.RabbitMqConn = rabbitMqConn
	return
}
