package filemanager

import "errors"

type ElementType int

const (
	Unknown ElementType = -1
	File ElementType = iota
	Directory
)

func StringToElementType(elementType string) (ElementType, error) {
	switch elementType {
	case "file", "f":
		return File, nil
	case "directory", "dir", "d":
		return Directory, nil
	default:
		return Unknown, errors.New("Invlaid type")
	}
}