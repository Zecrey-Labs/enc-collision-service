package utils

const (
	TypeEncData = iota
	TypeEncDataOmitSpace
)

const (
	// TODO(Gavin): these constraints is not settled yet and should be revised before production
	maxEncDataLength          = 60
	maxEncDataLengthOmitSpace = 50
)
