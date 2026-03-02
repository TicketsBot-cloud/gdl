package user

type ClientStatus struct {
	Desktop ClientStatusType `json:"desktop,omitempty"`
	Mobile  ClientStatusType `json:"mobile,omitempty"`
	Web     ClientStatusType `json:"web,omitempty"`
}
