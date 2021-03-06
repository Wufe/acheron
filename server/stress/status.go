package stress

type StressStatusType int

const (
	STATUS_IDLE StressStatusType = iota
	STATUS_RUNNING
	STATUS_COMPLETED
)

type StressStatus struct {
	Status                 StressStatusType
	Request                *StressSuiteRequest
	Results                *[]*PerformedRequestsResult
	TotalSucceededRequests uint64
	TotalFailedRequests    uint64
	TotalTime              uint64
}

var stressStatusInstance *StressStatus

func GetStressStatus() *StressStatus {
	if stressStatusInstance == nil {
		stressStatusInstance = &StressStatus{
			Status: STATUS_IDLE,
		}
	}
	return stressStatusInstance
}

func SetStatusToIdle() {
	stressStatusInstance = nil
}
