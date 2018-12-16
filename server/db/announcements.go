package db

import (
	"fmt"
	"server/models"
	"time"
)

func (ctx *DatabaseContext) GetAnnouncements() ([]models.Announcement, error) {
	result := make([]models.Announcement, 0)

	// This loads the entire dataset into memory at once so it will fail if the dataset it too large, but unlikely
	// to happen given the context of this app.
	err := ctx.DB.Select(&result, "SELECT `id`, `datetime`, `message` FROM `announcements` ORDER BY `datetime` DESC")
	if err != nil {
		return nil, fmt.Errorf("db GetAnnouncements failed: %v", err)
	}
	return result, nil
}

func (ctx *DatabaseContext) GetAnnouncementsByDate(date time.Time) ([]models.Announcement, error) {
	year, month, day := date.Date()

	st := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	et := st.Add(24 * time.Hour)

	result := make([]models.Announcement, 0)

	err := ctx.DB.Select(&result, "SELECT `id`, `datetime`, `message` FROM `announcements` WHERE (`datetime` BETWEEN ? AND ?) ORDER BY `datetime` DESC", st, et)
	if err != nil {
		return nil, fmt.Errorf("db GetAnnouncementsByDate failed: %v", err)
	}

	return result, nil
}

func (ctx *DatabaseContext) SaveAnnouncement(announcement models.Announcement) (*models.Announcement, error) {
	result, err := ctx.DB.Exec("INSERT INTO `announcements` (`datetime`, `message`) VALUES (?, ?)", announcement.Time, announcement.Message)
	if err != nil {
		return nil, fmt.Errorf("db SaveAnnouncement failed: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, nil
	}

	return &models.Announcement{ID: id, Time: announcement.Time, Message: announcement.Message}, nil
}
