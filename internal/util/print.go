package util

import "github.com/fatih/color"

const (
	formatSepLine = "------------------------ %s ------------------------"
)

func ModelGenFinished() {
	color.Green(formatSepLine, "Model Generate Finished")
}
