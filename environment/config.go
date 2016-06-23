package environment

import "os"

func Root() string {
	root := os.Getenv("QUINTO_PATH")
	if root == "" {
		root = "."
	}
	return root
}
