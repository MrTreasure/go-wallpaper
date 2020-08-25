package downloader

type payload struct {
	index int
	data  []byte
	err   error
}

type contentMeta struct {
	size        int
	contentType string
}
