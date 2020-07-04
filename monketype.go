package monketype

type MonkeType interface {
	FromMap(map[string]interface{}) (MonkeType, error)
	Map() map[string]interface{}
	JSON() ([]byte, error)
}
