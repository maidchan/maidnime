package models

type Anime struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Image         string   `json:"image"`
	Status        string   `json:"status,omitempty"`
	Rating        string   `json:"rating,omitempty"`
	Type          string   `json:"type,omitempty"`
	Genre         []string `json:"genre,omitempty"`
	LastEpisode   string   `json:"last_episode,omitempty"`
	LastEpisodeID string   `json:"last_episode_id,omitempty"`
}

type AnimeDetail struct {
	Anime
	Synopsis string    `json:"synopsis"`
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Date  string `json:"date"`
}

type StreamLink struct {
	Quality string `json:"quality"`
	URL     string `json:"url"`
}

type StreamResponse struct {
	Title     string       `json:"title"`
	Link      string       `json:"link"`
	Qualities []StreamLink `json:"qualities"`
}
