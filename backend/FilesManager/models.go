package filesmanager

type FileInfo struct {
	Name string
	Dir  string
	Meta map[interface{}]interface{}
}
