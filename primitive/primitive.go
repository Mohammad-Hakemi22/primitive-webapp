package primitive

import (
	"fmt"
	"os/exec"
	"strings"
)

type Mode int
type Options func() []string 

const (
	Combo Mode = iota
	Triangle
	Rect
	Ellipse
	Circle
	Rotatedrect
	Beziers
	Rotatedellipse
	Polygon
)

func WithMode(mode Mode) Options {
	return func() []string {
		return strings.Fields(fmt.Sprintf("-m %d", mode))
	}
}


func Primitive(inputfile, outputfile string, numShapes int, opts ...Options) (string, error) {
	cmdstr := fmt.Sprintf("-i %s -o %s -n %d", inputfile, outputfile, numShapes)
	command := strings.Fields(cmdstr)
	for _, opt := range opts {
		command = append(command, opt()...)
	}
	cmd := exec.Command("primitive", command...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(b), nil
}
