package dao

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"pilipili/backend/model"

	_ "github.com/go-sql-driver/mysql"
)

func (d *Dao) GetAllVideo() []*model.Video {
	querySQL, err := d.DB.Prepare("SELECT id, title, note, pic_path, video_path FROM video")
	if err != nil {
		logrus.Errorf("Prepare Select All SQL Error: %s", err.Error())
	}
	rows, err := querySQL.Query()
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query All Video Error: %s", err.Error())
	}

	var videoList []*model.Video
	for rows.Next() {
		v := new(model.Video)
		err = rows.Scan(v.Title, v.Note, v.PicPath, v.VideoPath)
		if err != nil {
			logrus.Errorf("Scan Query Error: %s", err.Error())
		}
		videoList = append(videoList, v)
	}
	return videoList
}

func (d *Dao) GetVideoInfo(id string) *model.Video {
	querySQL, err := d.DB.Prepare("SELECT title, note, pic_path, video_path FROM video WHERE id = ?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
	}

	v := new(model.Video)
	rows, err := querySQL.Query(id)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query Video Error: %s", err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&v.Title, &v.Note, &v.PicPath, &v.VideoPath)
		if err != nil {
			logrus.Errorf("Scan Query Error: %s", err.Error())
		}
		return v
	}
	return nil
}

func (d *Dao) CreateNewVideo(v *model.Video) {
	id := UUID()
	title := v.Title
	note := v.Note
	picPath := v.PicPath
	videoPath := v.VideoPath

	insertSql, err := d.DB.Prepare("INSERT INTO video (id, title, note, pic_path, video_path) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Prepare Insert SQL Error: %s", err.Error())
	}

	_, err = insertSql.Exec(id, title, note, picPath, videoPath)
	if err != nil {
		logrus.Errorf("Insert Video Info SQL Error: %s", err.Error())
	}
	defer insertSql.Close()
	return
}

func UUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Infof("UUID Error: %s", err.Error())
	}
	return id.String()
}
