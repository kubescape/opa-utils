package reporthandling

type IReportSummary interface {
	IReportStatus

	// Get
	SetNumberOfResources(n int)
	SetNumberOfWarningResources(n int)
	SetNumberOfFailedResources(n int)

	// Get
	GetNumberOfResources() int
	GetNumberOfWarningResources() int
	GetNumberOfFailedResources() int
}
type IReportStatus interface {
	GetStatus() string
	Passed() bool
	Warning() bool
	Failed() bool
}
