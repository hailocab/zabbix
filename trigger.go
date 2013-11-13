package zabbix

type (
	PriorityType int
)

const (
	Default     PriorityType = 0
	Information PriorityType = 1
	Warning     PriorityType = 2
	Average     PriorityType = 3
	High        PriorityType = 4
	Disaster    PriorityType = 5
)

// https://www.zabbix.com/documentation/2.0/manual/appendix/api/trigger/definitions
type Trigger struct {
	TriggerId   string       `json:"triggerid,omitempty"`
	Description string       `json:"description"`
	Expression  string       `json:"expression"`
	Comments    string       `json:"comments,omitempty"`
	Priority    PriorityType `json:"priority,omitempty"`
}

// Wrapper for trigger.create: https://www.zabbix.com/documentation/2.0/manual/appendix/api/trigger/create
func (api *API) TriggerCreate(trigger Trigger) (err error) {
	response, err := api.CallWithError("trigger.create", trigger)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	triggerids := result["triggerids"].([]interface{})
	for _, id := range triggerids {
		trigger.TriggerId = id.(string)
	}
	return
}
