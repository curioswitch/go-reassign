package dotimport

import . "io"

func direct() {
	EOF = nil // want "reassigning variable"
}

func indirect() {
	p := &EOF // want "taking address"
	*p = nil
}
