package monketype

type MonkeType interface {
	FromMap(map[string]interface{}) (MonkeType, error)
	JSON() ([]byte, error)
}
