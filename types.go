package main

type Places struct {
	Find bool `json:"find"`
	In   struct {
		Asis        bool    `json:"asis"`
		Cnt         int     `json:"cnt"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Nolimit     bool    `json:"nolimit"`
		OutCallback string  `json:"outCallback"`
		Q           string  `json:"q"`
		St          string  `json:"st"`
	} `json:"in"`
	Matches []struct {
		_geodist    float64 `json:"@geodist"`
		AddrType    string  `json:"addr_type"`
		AddrTypeID  int     `json:"addr_type_id"`
		CityID      int     `json:"city_id"`
		Class       string  `json:"class"`
		CountryID   int     `json:"country_id"`
		DisplayName string  `json:"display_name"`
		DistrictID  int     `json:"district_id"`
		FullName    string  `json:"full_name"`
		ID          int     `json:"id"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Name        string  `json:"name"`
		Operator    string  `json:"operator"`
		OsmID       string  `json:"osm_id"`
		RegionID    int     `json:"region_id"`
		StreetID    int     `json:"street_id"`
		ThisPoi     int     `json:"this_poi"`
		VillageID   int     `json:"village_id"`
		Weight      int     `json:"weight"`
	} `json:"matches"`
	Search  string `json:"search"`
	UserPos []struct {
		AddrType string `json:"addr_type"`
		FullName string `json:"full_name"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
	} `json:"userPos"`
	Ver string `json:"ver"`
}

type InlineQueryResultVenue struct {
	Type         string  `json:"type"`      // required
	ID           string  `json:"id"`        // required
	Latitude     float64 `json:"latitude"`  // required
	Longitude    float64 `json:"longitude"` // required
	Title        string  `json:"title"`     // required
	Address      string  `json:"address"`   // required
	Foursquareid string  `json:"foursquare_id,omitempty"`
	//ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ThumbURL            string      `json:"thumb_url"`
	ThumbWidth          int         `json:"thumb_width"`
	ThumbHeight         int         `json:"thumb_height"`
}
