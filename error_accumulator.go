package ML

import (
	"fmt"
	"io"
)

type errorAccumulator struct {
	totalCount int
	errorCount int
}

func (ea *errorAccumulator) Add(error float64) {
	ea.totalCount += 1
	if error != 0.0 {
		ea.errorCount += 1
	}
}

func (ea *errorAccumulator) Clone() ErrorAccumulator {
	return &errorAccumulator{
		totalCount: ea.totalCount,
		errorCount: ea.errorCount}
}

func (ea *errorAccumulator) Count() int {
	return ea.totalCount
}

func (ea *errorAccumulator) Estimate() float64 {
	if ea.totalCount > 0 {
		return float64(ea.errorCount)/float64(ea.totalCount)
	}
	return 0.0
}

func (ea *errorAccumulator) Clear() {
	ea.errorCount = 0
	ea.totalCount = 0
}

func (ea *errorAccumulator) Dump(w io.Writer, indent int) {
	fmt.Fprintf (w, "%*scount: %d, errorCount: %d\n", indent, "", ea.totalCount, ea.errorCount)
}










