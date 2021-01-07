package reddit

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const AccessTokenUrl = "https://www.reddit.com/api/v1/access_token"
const ReadUrl = "https://oauth.reddit.com"
const UserAgent = "reddit-bot"

type ResponseData struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  interface{} `json:"modhash"`
		Dist     int         `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Saved                      bool          `json:"saved"`
				ModReasonTitle             interface{}   `json:"mod_reason_title"`
				Gilded                     int           `json:"gilded"`
				Clicked                    bool          `json:"clicked"`
				Title                      string        `json:"title"`
				LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
				Hidden                     bool          `json:"hidden"`
				Pwls                       int           `json:"pwls"`
				LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
				Downs                      int           `json:"downs"`
				ThumbnailHeight            int           `json:"thumbnail_height"`
				TopAwardedType             interface{}   `json:"top_awarded_type"`
				HideScore                  bool          `json:"hide_score"`
				Name                       string        `json:"name"`
				Quarantine                 bool          `json:"quarantine"`
				LinkFlairTextColor         string        `json:"link_flair_text_color"`
				UpvoteRatio                float64       `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				SubredditType              string        `json:"subreddit_type"`
				Ups                        int           `json:"ups"`
				TotalAwardsReceived        int           `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        int           `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           struct {
					RedditVideo struct {
						BitrateKbps       int    `json:"bitrate_kbps"`
						FallbackURL       string `json:"fallback_url"`
						Height            int    `json:"height"`
						Width             int    `json:"width"`
						ScrubberMediaURL  string `json:"scrubber_media_url"`
						DashURL           string `json:"dash_url"`
						Duration          int    `json:"duration"`
						HlsURL            string `json:"hls_url"`
						IsGif             bool   `json:"is_gif"`
						TranscodingStatus string `json:"transcoding_status"`
					} `json:"reddit_video"`
				} `json:"secure_media"`
				IsRedditMediaDomain bool        `json:"is_reddit_media_domain"`
				IsMeta              bool        `json:"is_meta"`
				Category            interface{} `json:"category"`
				SecureMediaEmbed    struct {
				} `json:"secure_media_embed"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              bool          `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				PostHint            string      `json:"post_hint"`
				ContentCategories   interface{} `json:"content_categories"`
				IsSelf              bool        `json:"is_self"`
				ModNote             interface{} `json:"mod_note"`
				Created             float64     `json:"created"`
				LinkFlairType       string      `json:"link_flair_type"`
				Wls                 int         `json:"wls"`
				RemovedByCategory   interface{} `json:"removed_by_category"`
				BannedBy            interface{} `json:"banned_by"`
				AuthorFlairType     string      `json:"author_flair_type"`
				Domain              string      `json:"domain"`
				AllowLiveComments   bool        `json:"allow_live_comments"`
				SelftextHTML        interface{} `json:"selftext_html"`
				Likes               interface{} `json:"likes"`
				SuggestedSort       interface{} `json:"suggested_sort"`
				BannedAtUtc         interface{} `json:"banned_at_utc"`
				URLOverriddenByDest string      `json:"url_overridden_by_dest"`
				ViewCount           interface{} `json:"view_count"`
				Archived            bool        `json:"archived"`
				NoFollow            bool        `json:"no_follow"`
				IsCrosspostable     bool        `json:"is_crosspostable"`
				Pinned              bool        `json:"pinned"`
				Over18              bool        `json:"over_18"`
				Preview             struct {
					Images []struct {
						Source struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
						} `json:"variants"`
						ID string `json:"id"`
					} `json:"images"`
					Enabled bool `json:"enabled"`
				} `json:"preview"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    struct {
					RedditVideo struct {
						BitrateKbps       int    `json:"bitrate_kbps"`
						FallbackURL       string `json:"fallback_url"`
						Height            int    `json:"height"`
						Width             int    `json:"width"`
						ScrubberMediaURL  string `json:"scrubber_media_url"`
						DashURL           string `json:"dash_url"`
						Duration          int    `json:"duration"`
						HlsURL            string `json:"hls_url"`
						IsGif             bool   `json:"is_gif"`
						TranscodingStatus string `json:"transcoding_status"`
					} `json:"reddit_video"`
				} `json:"media"`
				IsVideo bool `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func mustBuildReadUrl(subreddit string) string {
	builder := strings.Builder{}
	if len(subreddit) == 0 {
		panic(fmt.Errorf("[buildReadUrl] No subreddit provided"))
	}
	builder.WriteString(ReadUrl)
	builder.WriteString("/")
	builder.WriteString(subreddit)
	builder.WriteString("/new.json")
	return builder.String()
}

func parseContent(body io.Reader) (*ResponseData, error) {
	var responseData ResponseData
	decodeError := json.NewDecoder(body).Decode(&responseData)
	if decodeError != nil {
		return nil, decodeError
	}
	return &responseData, nil
}

type Credentials struct {
	Username     string
	Password     string
	ClientId     string
	ClientSecret string
}

// TODO: redo this
func (resp ResponseData) String() string {
	builder := strings.Builder{}
	//builder.WriteString(fmt.Sprintf("{Id: %s}", resp.Id))
	builder.WriteString(",{data: { children: [")
	//for _, child := range resp.Data.Children {
	//	builder.WriteString(child.String())
	//}
	return builder.String()
}

type Error struct {
	id uuid.UUID
	error error
	message string
}

type Config struct {
	Limit int
}

type Client struct {
	Initialized bool
	accessToken string
	credentials *Credentials
	config *Config
	httpClient  *http.Client
}

func InitClient(credentials *Credentials) (*Client, error) {
	client := &Client{credentials: credentials, httpClient: &http.Client{}}
	// TODO add limit to config
	client.config = &Config{5}
	form := url.Values{
		"grant_type": {"password"},
		"username":   {credentials.Username},
		"password":   {credentials.Password},
	}
	log.Printf("Initializing Reddit Client 2")

	req, reqError := http.NewRequest("POST", AccessTokenUrl, strings.NewReader(form.Encode()))
	if reqError != nil {
		return nil, fmt.Errorf("error during Access token request creation: %w", reqError)
	}
	req.Header.Add("User-agent", UserAgent)
	req.SetBasicAuth(credentials.ClientId, credentials.ClientSecret)
	resp, respError := client.httpClient.Do(req)
	if respError != nil {
		return nil, fmt.Errorf("error during fetch access token:  %w", respError)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	var responseData map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseData)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error response code from auth api: %d", resp.StatusCode)
	}
	for k, v := range responseData {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}
	client.Initialized = true
	client.accessToken = fmt.Sprintf("%v", responseData["access_token"])
	return client, nil
}

// Request the client to fill in the buffer with Content
func (client Client) Read(subreddit string, respChannel chan *ResponseData, errorChannel chan error) {
	req, reqError := http.NewRequest("GET", mustBuildReadUrl(subreddit), nil)
	query := url.Values{}
	query.Add("limit", strconv.Itoa(client.config.Limit))
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Authorization", "Bearer " + client.accessToken)

	if reqError != nil {
		errorChannel <- fmt.Errorf("error during request creation: %w", reqError)
	}
	req.Header.Add("User-agent", UserAgent)
	client.httpClient.Do(req)
	resp, respError := client.httpClient.Do(req)
	if respError != nil {
		errorChannel <- fmt.Errorf("error during fetch content:  %w", respError)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}
	responseData, parseError := parseContent(resp.Body)
	if parseError != nil {
		errorChannel <- fmt.Errorf("error during fetch content:  %w", respError)
	}
	//json.NewDecoder(resp.Body).Decode(&responseData)
	log.Printf("Read and Parsed Content %s", responseData.String())
	respChannel <- responseData
}