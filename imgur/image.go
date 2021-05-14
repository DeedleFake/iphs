package imgur

type Image struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Desc        string    `json:"description"`
	DateTime    Timestamp `json:"datetime"`
	Type        string    `json:"type"`
	Animated    bool      `json:"animated"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	Size        int       `json:"size"`
	Views       int       `json:"views"`
	Bandwidth   int       `json:"bandwidth"`
	Vote        int       `json:"vote"`
	Favorite    bool      `json:"favorite"`
	NSFW        bool      `json:"nsfw"`
	Section     string    `json:"section"`
	AccountURL  string    `json:"account_url"`
	AccountID   int       `json:"account_id"`
	IsAd        bool      `json:"is_ad"`
	InMostViral bool      `json:"in_most_viral"`
	HasSound    bool      `json:"has_sound"`
	Tags        []string  `json:"tags"`
	AdType      int       `json:"ad_type"`
	AdURL       string    `json:"ad_url"`
	Edited      string    `json:"edited"`
	InGallery   bool      `json:"in_gallery"`
	Link        string    `json:"link"`
	// AdConfig ommitted because it's complex and unecessary.
}

func (c *Client) AlbumImages(album string) ([]Image, error) {
	rsp := response{
		Data: new([]Image),
	}
	err := c.get(&rsp, "album", album, "images")
	return *rsp.Data.(*[]Image), err
}
