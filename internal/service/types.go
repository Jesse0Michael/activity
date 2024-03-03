package service

type Activity struct {

	// Unique identifier for the activity
	ID string `json:"id"`

	// The source platform the activity is from
	Source string `json:"source"`

	// Unix timestamp (seconds) for when the item was published
	TS int64 `json:"ts"`

	// Permalink to the activity its source
	URL string `json:"url,omitempty"`

	// Text content for the item (may contain HTML)
	Content string `json:"content,omitempty"`

	// Array of media items (images, videos, etc...)
	Media []Media `json:"media,omitempty"`
}

type Media struct {

	// The URL to the media content
	URL string `json:"url"`

	// The kind of media
	Kind string `json:"kind"`

	// The URL to a thumbnail image
	Thumbnail string `json:"thumbnail,omitempty"`
}
