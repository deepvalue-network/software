package disks

type hydratedUser struct {
	Name     string            `json:"name" hydro:"0"`
	Seed     string            `json:"seed" hydro:"1"`
	Accesses *hydratedAccesses `json:"accesses" hydro:"2"`
}

type hydratedAccesses struct {
	List map[string]*hydratedAccess `json:"list" hydro:"0"`
}

type hydratedAccess struct {
	ID    string `json:"id" hydro:"0"`
	SigPK string `json:"sigpk" hydro:"1"`
	EncPK string `json:"encpk" hydro:"2"`
}
