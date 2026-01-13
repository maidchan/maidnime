package scraper

import (
	"backend/internal/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const SankaBaseURL = "https://www.sankavollerei.com/anime"

type SankaResponse struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type SankaAnime struct {
	Title           string `json:"title"`
	Poster          string `json:"poster"`
	Episodes        any    `json:"episodes"` // Can be int or string
	Score           string `json:"score"`
	LastReleaseDate string `json:"lastReleaseDate"`
	AnimeId         string `json:"animeId"`
	Status          string `json:"status"`
	Type            string `json:"type"`
}

type SankaAnimeList struct {
	AnimeList []SankaAnime `json:"animeList"`
}

type SankaHome struct {
	Ongoing   SankaAnimeList `json:"ongoing"`
	Completed SankaAnimeList `json:"completed"`
}

type SankaDetail struct {
	Title       string         `json:"title"`
	Poster      string         `json:"poster"`
	Score       string         `json:"score"`
	Status      string         `json:"status"`
	Episodes    any            `json:"episodes"`
	Type        string         `json:"type"`
	Synopsis    SankaSynopsis  `json:"synopsis"`
	GenreList   []SankaGenre   `json:"genreList"`
	EpisodeList []SankaEpisode `json:"episodeList"`
}

type SankaSynopsis struct {
	Paragraphs []string `json:"paragraphs"`
}

type SankaGenre struct {
	Title   string `json:"title"`
	GenreId string `json:"genreId"`
}

type SankaEpisode struct {
	Title     string `json:"title"`
	EpisodeId string `json:"episodeId"`
	Date      string `json:"date"`
}

type SankaStream struct {
	Title               string      `json:"title"`
	DefaultStreamingUrl string      `json:"defaultStreamingUrl"`
	Server              SankaServer `json:"server"`
}

type SankaServer struct {
	Qualities []SankaQuality `json:"qualities"`
}

type SankaQuality struct {
	Title      string        `json:"title"`
	ServerList []SankaMirror `json:"serverList"`
}

type SankaMirror struct {
	Title    string `json:"title"`
	ServerId string `json:"serverId"`
}

func getJSON(url string, target interface{}) error {
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var result struct {
		Status string          `json:"status"`
		Data   json.RawMessage `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return err
	}

	if result.Status != "success" {
		return fmt.Errorf("api error: %s", result.Status)
	}

	return json.Unmarshal(result.Data, target)
}

func UnwrapSafeLink(link string) string {
	if strings.Contains(link, "safe.php?url=") {
		parts := strings.Split(link, "safe.php?url=")
		if len(parts) > 1 {
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err == nil {
				return string(decoded)
			}
		}
	}
	// Also handle other common safe link patterns if needed
	u, err := url.Parse(link)
	if err == nil {
		q := u.Query()
		if urlParam := q.Get("url"); urlParam != "" {
			// Check if it looks like base64
			if decoded, err := base64.StdEncoding.DecodeString(urlParam); err == nil {
				// Only use if it starts with http
				if strings.HasPrefix(string(decoded), "http") {
					return string(decoded)
				}
			}
		}
	}
	return link
}

func GetServerURL(id string) (string, error) {
	var data struct {
		URL string `json:"url"`
	}
	err := getJSON(SankaBaseURL+"/server/"+id, &data)
	if err != nil {
		return "", err
	}
	return UnwrapSafeLink(data.URL), nil
}

func GetLatestAnime() ([]models.Anime, error) {
	var sankaHome SankaHome
	err := getJSON(SankaBaseURL+"/home", &sankaHome)
	if err != nil {
		return nil, err
	}

	var animeList []models.Anime
	for _, a := range sankaHome.Ongoing.AnimeList {
		animeList = append(animeList, models.Anime{
			ID:            a.AnimeId,
			Title:         a.Title,
			Image:         a.Poster,
			Rating:        a.Score,
			LastEpisode:   fmt.Sprintf("%v", a.Episodes),
			LastEpisodeID: a.AnimeId,
		})
	}
	return animeList, nil
}

func GetOngoingAnime() ([]models.Anime, error) {
	var data SankaAnimeList
	err := getJSON(SankaBaseURL+"/ongoing-anime", &data)
	if err != nil {
		return nil, err
	}

	var animeList []models.Anime
	for _, a := range data.AnimeList {
		animeList = append(animeList, models.Anime{
			ID:            a.AnimeId,
			Title:         a.Title,
			Image:         a.Poster,
			Rating:        a.Score,
			LastEpisode:   fmt.Sprintf("%v", a.Episodes),
			LastEpisodeID: a.AnimeId,
		})
	}
	return animeList, nil
}

func GetPopularAnime() ([]models.Anime, error) {
	var sankaHome SankaHome
	err := getJSON(SankaBaseURL+"/home", &sankaHome)
	if err != nil {
		return nil, err
	}

	var animeList []models.Anime
	for _, a := range sankaHome.Completed.AnimeList {
		animeList = append(animeList, models.Anime{
			ID:            a.AnimeId,
			Title:         a.Title,
			Image:         a.Poster,
			Rating:        a.Score,
			LastEpisode:   fmt.Sprintf("%v", a.Episodes),
			LastEpisodeID: a.AnimeId,
		})
	}
	return animeList, nil
}

func SearchAnime(query string) ([]models.Anime, error) {
	var data SankaAnimeList
	err := getJSON(SankaBaseURL+"/search/"+query, &data)
	if err != nil {
		return nil, err
	}

	var animeList []models.Anime
	for _, a := range data.AnimeList {
		animeList = append(animeList, models.Anime{
			ID:     a.AnimeId,
			Title:  a.Title,
			Image:  a.Poster,
			Status: a.Status,
			Rating: a.Score,
		})
	}
	return animeList, nil
}

func GetAnimeDetail(id string) (*models.AnimeDetail, error) {
	var data SankaDetail
	err := getJSON(SankaBaseURL+"/anime/"+id, &data)
	if err != nil {
		return nil, err
	}

	detail := &models.AnimeDetail{
		Anime: models.Anime{
			ID:     id,
			Title:  data.Title,
			Image:  data.Poster,
			Status: data.Status,
			Rating: data.Score,
			Type:   data.Type,
		},
		Synopsis: strings.Join(data.Synopsis.Paragraphs, "\n"),
	}

	for _, g := range data.GenreList {
		detail.Genre = append(detail.Genre, g.Title)
	}

	for _, e := range data.EpisodeList {
		detail.Episodes = append(detail.Episodes, models.Episode{
			ID:    e.EpisodeId,
			Title: e.Title,
			Date:  e.Date,
		})
	}

	return detail, nil
}

func GetEpisodeDetail(id string) (*models.StreamResponse, error) {
	var data SankaStream
	err := getJSON(SankaBaseURL+"/episode/"+id, &data)
	if err != nil {
		return nil, err
	}

	res := &models.StreamResponse{
		Title: data.Title,
		Link:  UnwrapSafeLink(data.DefaultStreamingUrl),
	}

	for _, q := range data.Server.Qualities {
		for _, s := range q.ServerList {
			// We point to our own server endpoint which redirects
			res.Qualities = append(res.Qualities, models.StreamLink{
				Quality: fmt.Sprintf("%s (%s)", q.Title, s.Title),
				URL:     "/api/server/" + s.ServerId,
			})
		}
	}

	return res, nil
}

func GetGenres() ([]string, error) {
	var data struct {
		GenreList []SankaGenre `json:"genreList"`
	}
	err := getJSON(SankaBaseURL+"/genre", &data)
	if err != nil {
		return nil, err
	}

	var genres []string
	for _, g := range data.GenreList {
		genres = append(genres, g.Title)
	}
	return genres, nil
}
