package structures

import "time"

type structure struct {
	content    Content
	isDeleted  bool
	executesOn *time.Time
	expiresOn  *time.Time
}

func createStructure(
	content Content,
	isDeleted bool,
) Structure {
	return createStructureInternally(content, isDeleted, nil, nil)
}

func createStructureWithExecutesOn(
	content Content,
	isDeleted bool,
	executesOn *time.Time,
) Structure {
	return createStructureInternally(content, isDeleted, executesOn, nil)
}

func createStructureWithExpiresOn(
	content Content,
	isDeleted bool,
	expiresOn *time.Time,
) Structure {
	return createStructureInternally(content, isDeleted, nil, expiresOn)
}

func createStructureWithExecutesOnAndExpiresOn(
	content Content,
	isDeleted bool,
	executesOn *time.Time,
	expiresOn *time.Time,
) Structure {
	return createStructureInternally(content, isDeleted, executesOn, expiresOn)
}

func createStructureInternally(
	content Content,
	isDeleted bool,
	executesOn *time.Time,
	expiresOn *time.Time,
) Structure {
	out := structure{
		content:    content,
		isDeleted:  isDeleted,
		executesOn: executesOn,
		expiresOn:  expiresOn,
	}

	return &out
}

// Content returns the content
func (obj *structure) Content() Content {
	return obj.content
}

// IsActive returns true if the structure is active, false otherwise
func (obj *structure) IsActive() bool {
	return true
}

// IsDeleted returns true if the structure is deleted
func (obj *structure) IsDeleted() bool {
	return obj.isDeleted
}

// HasExecutesOn returns true if the structure has an execution time, false otherwise
func (obj *structure) HasExecutesOn() bool {
	return obj.executesOn != nil
}

// ExecutesOn returns the execution time, if any
func (obj *structure) ExecutesOn() *time.Time {
	return obj.executesOn
}

// HasExpiresOn returns true if the structure has an expiration time, false otherwise
func (obj *structure) HasExpiresOn() bool {
	return obj.expiresOn != nil
}

// ExpiresOn returns the expiration time, if any
func (obj *structure) ExpiresOn() *time.Time {
	return obj.expiresOn
}
