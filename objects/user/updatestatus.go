package user

type UpdateStatus struct {
	Since      *int             `json:"since"` // unix time in ms when client went idle, null if not idle
	Activities []Activity       `json:"activities"`
	Status     ClientStatusType `json:"status"`
	Afk        bool             `json:"afk"`
}

func BuildStatus(activityType ActivityType, status string) UpdateStatus {
	return UpdateStatus{
		Activities: []Activity{
			{
				Name: status,
				Type: activityType,
			},
		},
		Status: ClientStatusTypeOnline,
	}
}
