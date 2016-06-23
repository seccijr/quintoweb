package environment

import "os"

func Root() string {
	root := os.Getenv("QUINTO_PATH")
	if root == "" {
		root = "/etc/root"
	}
	return root
}
