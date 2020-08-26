package model

type Names struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type FacebookAudienceSegment struct {
	ID             string `json:"id"`
	InstanceID     string `json:"instance_id"`
	AdvertiserID   string `json:"advertiser_id"`
	FbAccountID    string `json:"fb_account_id"`
	Typed          string `json:"type"`
	CreatedTime    uint64 `json:"created_time"`
	ModifiedTime   uint64 `json:"modified_time"`
	Archived       bool   `json:"archived"`
	DmpAudience    string `json:"dmp_audience_segment_ref"`
	ServiceSegment string `json:"service_segment"`
	EventSource    string `json:"event_source_group"`
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	GroupID  int    `json:"groupID"`
	Status   int    `json:"status"`
}

type Task struct {
	TaskID      int    `json:"taskID"`
	Title       string `json:"title"`
	Desc        string `json:"description"`
	CreatedBy   string `json:"created_by"`
	AssignedTo  string `json:"assigned_to"`
	Approved    bool   `json:"approved"`
	Status      string `json:"status"`
	GroupID     int    `json:"groupID"`
	SourceID    int    `json:"sourceID"`
	Comment     string `json:"comment"`
	Mark        int    `json:"mark"`
	CommentTime string `json:"commentTime"`
	Start       string `json:"start"`
	Finish      string `json:"finish"`
	CreatedAt   string `json:"createdAt"`
	TimeSpent   string `json:"timeSpent"`
	ApprovedAt  string `json:"approvedAt"`
	// ConfirmationIMG string `json:"confirmationIMG"`
}

type Group struct {
	GroupID   int    `json:"groupID"`
	GroupName string `json:"groupName"`
}
