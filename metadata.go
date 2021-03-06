package youtubedlwrapper

type HTTPHeaders struct {
	UserAgent      string `json:"User-Agent"`
	AcceptCharset  string `json:"Accept-Charset"`
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	AcceptLanguage string `json:"Accept-Language"`
}

type DownloaderOptions struct {
	HTTPChunkSize int `json:"http_chunk_size"`
}

type Fragment struct {
	Path     string  `json:"path"`
	Duration float64 `json:"duration,omitempty"`
}

type Format struct {
	FormatID          string            `json:"format_id"`
	ManifestURL       string            `json:"manifest_url,omitempty"`
	Ext               string            `json:"ext"`
	Width             int               `json:"width"`
	Height            int               `json:"height"`
	Tbr               float64           `json:"tbr"`
	Asr               int               `json:"asr"`
	Fps               int               `json:"fps"`
	Language          string            `json:"language,omitempty"`
	FormatNote        string            `json:"format_note"`
	Filesize          int               `json:"filesize"`
	Container         string            `json:"container,omitempty"`
	Vcodec            string            `json:"vcodec"`
	Acodec            string            `json:"acodec"`
	URL               string            `json:"url"`
	FragmentBaseURL   string            `json:"fragment_base_url,omitempty"`
	Fragments         []Fragment        `json:"fragments,omitempty"`
	Protocol          string            `json:"protocol"`
	Format            string            `json:"format"`
	HTTPHeaders       HTTPHeaders       `json:"http_headers"`
	Quality           int               `json:"quality,omitempty"`
	Abr               float64           `json:"abr,omitempty"`
	DownloaderOptions DownloaderOptions `json:"downloader_options,omitempty"`
	Vbr               float64           `json:"vbr,omitempty"`
}

type Thumbnail struct {
	Height     int    `json:"height"`
	URL        string `json:"url"`
	Width      int    `json:"width"`
	Resolution string `json:"resolution"`
	ID         string `json:"id"`
}

type AutomaticCaption struct {
	Ext string `json:"ext"`
	URL string `json:"url"`
}

type Chapter struct {
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	Title     string `json:"title"`
}

type VideoMetadata struct {
	ID                 string                        `json:"id"`
	Title              string                        `json:"title"`
	Formats            []Format                      `json:"formats"`
	Thumbnails         []Thumbnail                   `json:"thumbnails"`
	Description        string                        `json:"description"`
	UploadDate         string                        `json:"upload_date"`
	Uploader           string                        `json:"uploader"`
	UploaderID         string                        `json:"uploader_id"`
	UploaderURL        string                        `json:"uploader_url"`
	ChannelID          string                        `json:"channel_id"`
	ChannelURL         string                        `json:"channel_url"`
	Duration           int                           `json:"duration"`
	ViewCount          int                           `json:"view_count"`
	AverageRating      float64                       `json:"average_rating"`
	AgeLimit           int                           `json:"age_limit"`
	WebpageURL         string                        `json:"webpage_url"`
	Categories         []string                      `json:"categories"`
	Tags               []string                      `json:"tags"`
	IsLive             bool                          `json:"is_live"`
	LikeCount          int                           `json:"like_count"`
	DislikeCount       int                           `json:"dislike_count"`
	Channel            string                        `json:"channel"`
	Track              string                        `json:"track"`
	Artist             string                        `json:"artist,omitempty"`
	Album              string                        `json:"album,omitempty"`
	Creator            string                        `json:"creator,omitempty"`
	AltTitle           string                        `json:"alt_title"`
	Extractor          string                        `json:"extractor"`
	WebpageURLBasename string                        `json:"webpage_url_basename"`
	ExtractorKey       string                        `json:"extractor_key"`
	NEntries           int                           `json:"n_entries"`
	Playlist           string                        `json:"playlist"`
	PlaylistID         string                        `json:"playlist_id"`
	PlaylistTitle      string                        `json:"playlist_title"`
	PlaylistUploader   string                        `json:"playlist_uploader"`
	PlaylistUploaderID string                        `json:"playlist_uploader_id"`
	PlaylistIndex      int                           `json:"playlist_index"`
	Thumbnail          string                        `json:"thumbnail"`
	DisplayID          string                        `json:"display_id"`
	RequestedSubtitles string                        `json:"requested_subtitles"`
	RequestedFormats   []Format                      `json:"requested_formats"`
	Format             string                        `json:"format"`
	FormatID           string                        `json:"format_id"`
	Width              int                           `json:"width"`
	Height             int                           `json:"height"`
	Resolution         string                        `json:"resolution"`
	Fps                int                           `json:"fps"`
	Vcodec             string                        `json:"vcodec"`
	Vbr                float64                       `json:"vbr"`
	StretchedRatio     string                        `json:"stretched_ratio"`
	Acodec             string                        `json:"acodec"`
	Abr                float64                       `json:"abr"`
	Ext                string                        `json:"ext"`
	License            string                        `json:"license"`
	Chapters           []Chapter                     `json:"chapters"`
	AutomaticCaptions  map[string][]AutomaticCaption `json:"automatic_captions,omitempty"`
	Subtitles          map[string]string             `json:"subtitles,omitempty"`
}

type PlaylistMetadata struct {
	Type               string          `json:"_type"`
	Videos             []VideoMetadata `json:"entries"`
	ID                 string          `json:"id"`
	Title              string          `json:"title"`
	Uploader           string          `json:"uploader"`
	UploaderID         string          `json:"uploader_id"`
	UploaderURL        string          `json:"uploader_url"`
	Extractor          string          `json:"extractor"`
	WebpageURL         string          `json:"webpage_url"`
	WebpageURLBasename string          `json:"webpage_url_basename"`
	ExtractorKey       string          `json:"extractor_key"`
}
