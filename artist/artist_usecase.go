package artist

import "artistSpotify/model"

type ArtistUsecase interface {
	FindAllArtist()(*[]model.Artist,error)
	FindArtistById(id int)(*model.Artist,error)
	InsertArtist(artist *model.Artist)error
	UpdateArtist(id int,artist *model.Artist)error
	DeleteArtist(id int)error
}
