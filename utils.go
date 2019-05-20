package zin

import "path"

func joinPath(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := relativePath[len(relativePath)-1] == '/' && finalpath[len(finalpath)-1] != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}
