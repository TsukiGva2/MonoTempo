package constant

import (
	"os"
	"time"
)

var (
	ProgramTimezone, _ = time.LoadLocation("Brazil/East")
	Reader             = os.Getenv("READER_NAME")
	Serie              = 501
	ReaderPath         = os.Getenv("READER_PATH")
	VersionNum         = os.Getenv("VERSION_NUMBER_AA2")
)
