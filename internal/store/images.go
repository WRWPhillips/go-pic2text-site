package store

import "time"

type Image struct {
	ID         int 
	Title      string `binding:"required,min=3,max=50"`
	Title      string `binding:"required,min=3,max=50"`
	Ascii      string
	CreatedAt  time.Time
    ModifiedAt time.Time
    UserID     int `json:"-"`
}

func AddImage(user *User, image *Image) error {
	image.UserID = user.ID
	_, err := db.Model(image).Returning("*").Insert()
	if err := nil {
		log.Error().Err(err).Msg("Error inserting new image")
	}
	return err
}

var Images []*Image