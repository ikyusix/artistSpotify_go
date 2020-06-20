package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"artistSpotify/artist"
	"artistSpotify/model"
)

type ArtistRepoImpl struct {
	db *sql.DB
}

func CreateArtistRepoMysqlImpl(db *sql.DB) artist.ArtistRepo {
	return &ArtistRepoImpl{db}
}

func (a *ArtistRepoImpl) FindAllArtist() (*[]model.Artist, error) {
	query := "SELECT id,name,debut,category FROM artist"

	rows,err := a.db.Query(query)

	if err != nil{
		if errors.Is(err,sql.ErrNoRows){
			return nil,nil
		}
		return nil, fmt.Errorf("[ArtistRepoImpl.FindAllArtist] Error when FindAllArtist artist : %w", err)
	}

	defer rows.Close()

	var result []model.Artist
	for rows.Next(){
		dataArtist := model.Artist{}
		err = rows.Scan(&dataArtist.ID,&dataArtist.Name,&dataArtist.Debut,&dataArtist.Category)
		if err != nil{
			return nil, fmt.Errorf("[ArtistRepoImpl.FindAll] Error when scan rows artists : %w", err)
		}
		result = append(result,dataArtist)
	}

	return &result, nil
}

func (a *ArtistRepoImpl) FindArtistById(id int) (*model.Artist, error) {
	query := "SELECT id,name,debut,category FROM artist Where id=?"

	dataArtist := model.Artist{}
	err := a.db.QueryRow(query,id).Scan(&dataArtist.ID,&dataArtist.Name,&dataArtist.Debut,&dataArtist.Category)
	if err != nil{
		if errors.Is(err,sql.ErrNoRows){
			return nil,nil
		}
		return nil, fmt.Errorf("[ArtistRepoImpl.FindArtistById] Error when FindArtistById artist : %w", err)
	}

	return &dataArtist,nil
}

func (a *ArtistRepoImpl) InsertArtist(artist *model.Artist) error {
	query := "INSERT INTO artist (name,debut,category) VALUES (?,?,?)"

	_,err := a.db.Exec(query,artist.Name,artist.Debut,artist.Category)
	if err != nil{
		return fmt.Errorf("[ArtistRepoImpl.InsertArtist] Error when Insert artist : %w", err)
	}
	return nil
}

func (a *ArtistRepoImpl) UpdateArtist(id int,artist *model.Artist) error {
	query := "UPDATE artist SET name=?, debut=?, category=? WHERE id=?"

	_,err := a.db.Exec(query,artist.Name,artist.Debut,artist.Category,id)
	if err != nil {
		return fmt.Errorf("[ArtistRepoImpl.UpdateArtist] Error when Update Artist: %w", err)
	}
	return nil
}

func (a *ArtistRepoImpl) DeleteArtist(id int) error {
	query := "DELETE FROM artist WHERE id=?"

	_,err := a.db.Exec(query,id)
	if err != nil {
		return fmt.Errorf("[ArtistRepoImpl.DeleteArtist] Error when Delete Artist: %w", err)
	}
	return nil
}

