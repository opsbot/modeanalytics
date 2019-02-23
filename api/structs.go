package mode

// APIResponse struct
type APIResponse struct {
	Status     string
	StatusCode int
	Body       []byte
}

// Configuration struct
type Configuration struct {
	Host      string `json:"mode_host"`
	Org       string `json:"mode_org"`
	APIKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

// SpaceCollection struct
type SpaceCollection struct {
	Spaces []Space `json:"spaces"`
}

// Space struct
type Space struct {
	Token       string `json:"token"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Restricted  bool   `json:"Restricted"`
	SpaceType   string `json:"space_type"`
	State       bool   `json:"state"`
}

// Organization struct
type Organization struct {
	Token                   string   `json:"token"`
	Name                    string   `json:"name"`
	AuthorizedDomains       []string `json:"authorized_domains"`
	DataSourceCount         string   `json:"data_source_count"`
	Membershiptype          string   `json:"membership_type"`
	PaymentMenthodConfirmed bool     `json:"payment_method_confirmed"`
	PlanCode                string   `json:"plan_code"`
	PrivateDefinitionCount  string   `json:"private_definition_count"`
	PrivateDefinitionLimit  string   `json:"private_definition_limit"`
	SpaceCount              string   `json:"space_count"`
	TrialState              string   `json:"trial_state"`
	User                    bool     `json:"user"`
	UserName                string   `json:"union_pos"`
}
