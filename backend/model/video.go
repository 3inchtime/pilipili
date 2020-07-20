package model

type Video struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	PicPath   string `json:"pic_path"`
	VideoPath string `json:"video_path"`
}
