package wallpaper_abyss

type Res struct {
	Success    bool         `json:"success"`
	Wallpapers []Wallpapers `json:"wallpapers"`
}

type Wallpapers struct {
	ID            string `json:"id"`
	Width         string `json:"width"`
	Height        string `json:"height"`
	FileType      string `json:"file_type"`
	FileSize      string `json:"file_size"`
	URLImage      string `json:"url_image"`
	URLThumb      string `json:"url_thumb"`
	URLPage       string `json:"url_page"`
	Category      string `json:"category"`
	CategoryID    string `json:"category_id"`
	SubCategory   string `json:"sub_category"`
	SubCategoryID string `json:"sub_category_id"`
	UserName      string `json:"user_name"`
	UserID        string `json:"user_id"`
}
