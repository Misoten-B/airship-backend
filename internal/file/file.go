package file

import "mime/multipart"

type File struct {
	file       multipart.File
	fileHeader *multipart.FileHeader
}

func NewMyFile(file multipart.File, fileHeader *multipart.FileHeader) *File {
	return &File{
		file:       file,
		fileHeader: fileHeader,
	}
}

func (f *File) File() multipart.File {
	return f.file
}

func (f *File) FileHeader() *multipart.FileHeader {
	return f.fileHeader
}
