package youtubedlwrapper

import "net"

type Options struct {
	Username      string
	Password      string
	Videopassword string
	Ap_mso        string
	Ap_password   string
	Usenetrc      bool
	//verbose
	//quiet
	//no_warnings
	//forceurl
	//forcetitle
	//forceid
	//forcethumbnail
	//forcedescription
	//forcefilename
	//forceduration
	//forcejson
	//Dump_single_json bool
	Simulate                bool
	Format                  string
	Outtmpl                 string
	Restrictfilenames       bool
	Ignoreerrors            bool
	Force_generic_extractor bool
	Nooverwrites            bool
	Playliststart           int
	Playlistend             int
	Playlist_items          []int
	Playlistreverse         bool
	Playlistrandom          bool
	Matchtitle              bool
	Rejecttitle             bool
	//logger
	//logtostderr
	Writedesription        bool
	Writeinfojson          bool
	Writeannotations       bool
	Writethumbnail         bool
	Write_all_thumbnails   bool
	Writesubtitles         bool
	Writeautomaticsub      bool
	Allsubtitles           bool
	Listsubtitles          bool
	Subtitlesformat        string
	Subtitleslangs         []string
	Keepvideo              bool
	Daterange              string
	Skip_download          bool
	Cachedir               string
	Noplaylist             bool
	Age_limit              int
	Min_views              int
	Max_views              int
	Download_archive       string
	Cookiefile             string
	Nocheckcertificate     bool
	Prefer_insecure        bool
	Proxy                  string
	Geo_verification_proxy string
	Socket_timeout         int
	//bidi_workaround
	//debug_printtraffic
	Include_ads    bool
	Default_search string
	Encoding       string
	Extract_flat   bool
	//postprocessors
	//progress_hooks
	Fixup              string
	Source_address     net.Addr
	Call_home          bool
	Sleep_interval     int
	Max_sleep_interval int
	//listformats
	//list_thumbnails
	//match_filter
	//no_color
	Geo_bypass         bool
	Geo_bypass_country string
	//external_downloader
	//hls_prefer_native
	//implement parameters for different downloaders?
	Prefer_ffmpeg                 bool
	Postprocessor_args            []string
	Youtube_include_dash_manifest bool
}
