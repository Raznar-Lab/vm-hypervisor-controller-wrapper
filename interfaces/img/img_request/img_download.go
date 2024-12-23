package img_request

type IMGDownloadImageRequest struct {
	DownloadURL string `json:"download_url"`
	Filename    string `json:"filename"`
}
