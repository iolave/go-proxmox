package node

type GetFWRuleResponse struct {
	ID          string
	Action      string `json:"action"`
	Comment     string `json:"comment"`
	Destination string `json:"dest"`
	DPort       string `json:"destPort"`
	Enable      int    `json:"enable"`
	ICMPType    string `json:"icmpType"`
	Interface   string `json:"iface"`
	IPVersion   int    `json:"ipVersion"`
	LogLevel    string `json:"log"`
	Macro       string `json:"macro"`
	Pos         int    `json:"pos"`
	Proto       string `json:"proto"`
	Source      string `json:"source"`
	SPort       string `json:"srcPort"`
	Type        string `json:"type"`
}

func (s Service) GetFWRules(node string)
