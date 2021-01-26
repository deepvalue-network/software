package peers

import "time"

type peer struct {
	content       Content
	createdOn     time.Time
	lastUpdatedOn time.Time
}

func createPeer(
	content Content,
	createdOn time.Time,
	lastUpdatedOn time.Time,
) Peer {
	out := peer{
		content:       content,
		createdOn:     createdOn,
		lastUpdatedOn: lastUpdatedOn,
	}

	return &out
}

// Content returns the content
func (obj *peer) Content() Content {
	return obj.content
}

// CreatedOn returns the creation time
func (obj *peer) CreatedOn() time.Time {
	return obj.createdOn
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *peer) LastUpdatedOn() time.Time {
	return obj.lastUpdatedOn
}
