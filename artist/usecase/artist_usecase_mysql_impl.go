package usecase

import (
	"artistSpotify/artist"
	"artistSpotify/model"
)

type ArtistUsecaseImpl struct {
	artistRepo artist.ArtistRepo
}

func CreateArtistUsecaseImpl(artistRepo artist.ArtistRepo) artist.ArtistUsecase {
	return &ArtistUsecaseImpl{artistRepo}
}

func (a *ArtistUsecaseImpl) FindAllArtist() (*[]model.Artist, error) {
	return a.artistRepo.FindAllArtist()
}

func (a *ArtistUsecaseImpl) FindArtistById(id int) (*model.Artist, error) {
	return a.artistRepo.FindArtistById(id)
}

func (a *ArtistUsecaseImpl) InsertArtist(artist *model.Artist) error {
	return a.artistRepo.InsertArtist(artist)
}

func (a *ArtistUsecaseImpl) UpdateArtist(id int,artist *model.Artist) error {
	return a.artistRepo.UpdateArtist(id,artist)
}

func (a ArtistUsecaseImpl) DeleteArtist(id int) error {
	return a.artistRepo.DeleteArtist(id)
}
