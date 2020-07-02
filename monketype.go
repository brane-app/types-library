package monketype

type MonkeType interface {
	FromMap(map[string]interface{}) (MonkeType, error)
	Map() (map[string]interface{}, error)
	JSON() ([]byte, error)
}
