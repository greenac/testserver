package models

type Member struct {
	PpId string `json:"id"`
	ApiUri string `json:"api_uri"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Party string `json:"party"`
	LeadershipRole string `json:"leadership_role"`
	TwitterAccount string `json:"twitter_account"`
	FacebookAccount string `json:"facebook_account"`
	Url string `json:"url"`
	GovTrackId string `json:"govtrack_id"`
	InOffice bool `json:"in_office"`
	NextElection int `json:"next_election"`
	TotalVotes int `json:"total_votes"`
	MissedVotes int `json:"missed_votes"`
	Office string `json:"office"`
	Phone string `json:"phone"`
	State string `json:"state"`
}

func (mem *Member)ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = mem.PpId
	m["api_uri"] = mem.ApiUri
	m["first_name"] = mem.FirstName
	m["last_name"] = mem.LastName
	m["party"] = mem.Party
	m["leadership_role"] = mem.LeadershipRole
	m["twitter_account"] = mem.TwitterAccount
	m["facebook_account"] = mem.FacebookAccount
	m["url"] = mem.Url
	m["govtrack_id"] = mem.GovTrackId
	m["in_office"] = mem.InOffice
	m["next_election"] = mem.NextElection
	m["total_votes"] = mem.TotalVotes
	m["missed_votes"] = mem.MissedVotes
	m["office"] = mem.Office
	m["phone"] = mem.Phone
	m["state"] = mem.State

	return m
}
