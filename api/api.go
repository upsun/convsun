package api

import "github.com/upsun/convsun/internal/logic"

func Convert(projectWorkspace string, outputFilePath string) {
	logic.Convert(projectWorkspace, outputFilePath)
}
