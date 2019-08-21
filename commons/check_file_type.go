package commons

/**
check if the file type of given source path matches given file type
 */
func FileTypeCheck(link string, fileType string) bool {
	length := len(link)
	return length > 4 && link[length-len(fileType):length] == fileType
}
