package filesmanager

type FileInfo struct {
	Name string `json:"Name"`
	Dir  string `json:"Dir"`
}

type ByteFile struct {
	Name    string `json:"Name"`
	Content []byte `json:"Content"`
}
