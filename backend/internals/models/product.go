package models

type ProductSummary struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	ThumbnailURI string `json:"thumbnail"`
	Rating       int    `json:"rating"`
	Category     string `json:"category"`
}

type ProductDetail struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Price          int            `json:"price"`
	ImageURIs      []string       `json:"images"`
	Specifications Specifications `json:"specifications"`
	Rating         int            `json:"rating"`
	Reviews        []Review       `json:"reviews"`
	Category       string         `json:"category"`
}

type Specifications map[string]string

type Review struct {
	User      string   `json:"user"`
	Comment   string   `json:"comment"`
	Rating    int      `json:"rating"`
	ImageURIs []string `json:"images"`
}
