package monketype

type MonkeType interface {
	Map() map[string]interface{}
	JSON() ([]byte, error)
}
