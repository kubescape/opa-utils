package reporthandling

type ResourcesIDs struct {
	passedResources  []string
	failedResources  []string
	warningResources []string
}

func (r *ResourcesIDs) append(a *ResourcesIDs) {
	r.setFailedResources(append(r.failedResources, a.GetFailedResources()...))
	r.setWarningResources(append(r.warningResources, a.GetWarningResources()...)) // initialize after failed
	r.setPassedResources(append(r.passedResources, a.GetPassedResources()...))    // initialize after warning
}

func (r *ResourcesIDs) GetAllResources() []string {
	return append(append(r.GetFailedResources(), r.GetWarningResources()...), r.GetPassedResources()...)
}
func (r *ResourcesIDs) GetPassedResources() []string {
	return r.passedResources
}
func (r *ResourcesIDs) GetFailedResources() []string {
	return r.failedResources
}
func (r *ResourcesIDs) GetWarningResources() []string {
	return r.warningResources
}

func (r *ResourcesIDs) setFailedResources(a []string) {
	r.failedResources = GetUniqueResourcesIDs(a)
}

// setWarningResources - initialized after failed resources are set
func (r *ResourcesIDs) setWarningResources(a []string) {
	r.warningResources = TrimUniqueIDs(GetUniqueResourcesIDs(a), r.failedResources)
}

// setPassedResources - initialized after warning resources are set
func (r *ResourcesIDs) setPassedResources(a []string) {
	r.passedResources = TrimUniqueIDs(GetUniqueResourcesIDs(a), append(r.failedResources, r.warningResources...))
}

func deleteFromMap(m map[string]interface{}, keepFields []string) {
	for k := range m {
		if StringInSlice(keepFields, k) {
			continue
		}
		delete(m, k)
	}
}

func StringInSlice(strSlice []string, str string) bool {
	for i := range strSlice {
		if strSlice[i] == str {
			return true
		}
	}
	return false
}

func RemoveResponse(slice []RuleResponse, index int) []RuleResponse {
	return append(slice[:index], slice[index+1:]...)
}

func percentage(big, small int) int {
	if big == 0 {
		if small == 0 {
			return 100
		}
		return 0
	}
	return int(float64(float64(big-small)/float64(big)) * 100)
}
