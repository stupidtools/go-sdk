package connectionmodes

type ConnectionModes string

var (
	SERIAL  ConnectionModes = "SERIAL"
	TCP_IP  ConnectionModes = "TCP_IP"
	SSL_TCP ConnectionModes = "SSL_TCP"
	HTTP    ConnectionModes = "HTTP"
)
