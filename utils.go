package zhanst

import "path"

func joinPath(absolutePath, relativePath string) string {
	if relativePath == "" {
		return absolutePath
	}
	finalPath := path.Join(absolutePath, relativePath)
	appendSlash := relativePath[len(relativePath)-1] == '/' && finalPath[len(finalPath)-1] != '/'
	if appendSlash {
		return finalPath + "/"
	}
	return finalPath
}
