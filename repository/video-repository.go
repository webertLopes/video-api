package repository

import (
	"interface/entity"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	Save(entity.Video) entity.Video
	GetAll() []entity.Video
}

type videoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) VideoRepository {
	return &videoRepository{
		db: db,
	}
}

func (repository *videoRepository) Save(video entity.Video) entity.Video {

	err := repository.db.Debug().Save(video).Error

	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	return video
}

func (repository *videoRepository) GetAll() []entity.Video {

	var videos []entity.Video

	repository.db.Debug().Find(&videos)

	return videos
}
