package response

type Meta struct {
	TotalRecord uint64 `json:"total_record"`
	Page        uint64 `json:"page"`
	Limit       uint64 `json:"limit"`
}
