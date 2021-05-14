package imgur

type Album struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	Desc            string    `json:"description"`
	DateTime        Timestamp `json:"datetime"`
	Cover           string    `json:"cover"`
	CoverEdited     int       `json:"cover_edited"`
	CoverWidth      int       `json:"cover_width"`
	CoverHeight     int       `json:"cover_height"`
	AccountURL      string    `json:"account_url"`
	AccountID       int       `json:"account_id"`
	Privacy         string    `json:"privacy"`
	Layout          string    `json:"layout"`
	Views           int       `json:"views"`
	Link            string    `json:"link"`
	Favorite        bool      `json:"favorite"`
	NSFW            bool      `json:"nsfw"`
	Section         string    `json:"section"`
	ImagesCount     int       `json:"images_count"`
	InGallery       bool      `json:"in_gallery"`
	IsAd            bool      `json:"is_ad"`
	IncludeAlbumAds bool      `json:"include_album_ads"`
	IsAlbum         bool      `json:"is_album"`
	Order           int       `json:"order"`
}

// Albums returns a list of the albums for a given user.
func (c *Client) Albums(username string) ([]Album, error) {
	rsp := response{
		Data: new([]Album),
	}
	err := c.get(&rsp, "account", username, "albums")
	return *rsp.Data.(*[]Album), err
}
