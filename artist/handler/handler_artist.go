package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"artistSpotify/artist"
	"artistSpotify/model"
	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	artistUsecase artist.ArtistUsecase
}

func CreateArtistHandler(r *mux.Router, artistUsecase artist.ArtistUsecase)  {
	artistHandler := ArtistHandler{artistUsecase}

	r.HandleFunc("/artist",artistHandler.FindAllArtist).Methods(http.MethodGet)
	s:= r.PathPrefix("/artist").Subrouter()
	s.HandleFunc("/{id}",artistHandler.FindArtistById).Methods(http.MethodGet)
	s.HandleFunc("/add",artistHandler.InsertArtist).Methods(http.MethodPost)
	s.HandleFunc("/{id}",artistHandler.UpdateArtist).Methods(http.MethodPut)
	s.HandleFunc("/{id}",artistHandler.DeleteArtist).Methods(http.MethodDelete)
}

func handlerSuccess(resp http.ResponseWriter, data interface{}) {
	returnData := model.ResponseWrapper{
		Success : true,
		Message:"SUCCESS",
		Data:data,
	}


	jsonData,err := json.Marshal(returnData)
	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops something went wrong"))
		fmt.Printf("[ArtistHandler.Find] Error when do json Marshalling for error handling : %v \n", err)
	}

	resp.Header().Set("Content-Type","application/json")
	resp.Write(jsonData)
}
func handlerError(resp http.ResponseWriter, message string) {
	data := model.ResponseWrapper{
		Success : true,
		Message:message,
	}
	jsonData,err := json.Marshal(data)
	if err != nil{
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops something went wrong"))
		fmt.Printf("[ArtistHandler.Find] Error when do json Marshalling for error handling : %v \n", err)
	}
	resp.Header().Set("Content-Type","application/json")
	resp.Write(jsonData)
}

func (a *ArtistHandler) FindAllArtist(resp http.ResponseWriter, req *http.Request)  {
	artist,err :=a.artistUsecase.FindAllArtist()
	if err != nil {
		handlerError(resp,err.Error())
		return
	}
	handlerSuccess(resp,artist)
}
func (a *ArtistHandler) FindArtistById(resp http.ResponseWriter, req *http.Request)  {
	val := mux.Vars(req)
	valId,err := strconv.Atoi(val["id"])
	if err != nil {
		handlerError(resp,"ID harus angka")
		return
	}
	artist,err := a.artistUsecase.FindArtistById(valId)
	if err != nil{
		handlerError(resp,err.Error())
		return
	}
	handlerSuccess(resp,artist)
}

func (a *ArtistHandler) InsertArtist(resp http.ResponseWriter, req *http.Request)  {
	data := model.Artist{}
	reqBody,err := ioutil.ReadAll(req.Body)
	if err != nil {
		handlerError(resp,"Ooops, Something went wrong")
		fmt.Println("[ShipsHandler.InsertArtist] error when reading request body : " + err.Error())
		return
	}
	err = json.Unmarshal(reqBody,&data)
	if err != nil{
		handlerError(resp,"Ooops, Something went wrong")
		fmt.Println("[ShipsHanlder.InsertArtist] error when unmarshall artist json : " + err.Error())
		return
	}
	err = a.artistUsecase.InsertArtist(&data)
	if err != nil {
		handlerError(resp,err.Error())
		fmt.Println("[ShipsHandler.InsertArtist] error when call insert service : " + err.Error())
		return
	}
	handlerSuccess(resp,nil)
}

func (a *ArtistHandler) UpdateArtist(resp http.ResponseWriter, req *http.Request)  {
	val := mux.Vars(req)
	valId,err := strconv.Atoi(val["id"])
	if err != nil {
		handlerError(resp,"ID harus angka")
		return
	}
	data := model.Artist{}
	reqBody,err :=ioutil.ReadAll(req.Body)
	if err != nil {
		handlerError(resp,"Ooops, Something went wrong")
		fmt.Println("[ShipsHandler.UpdateArtist] error when reading request body : " + err.Error())
		return
	}

	err = json.Unmarshal(reqBody,&data)
	if err != nil{
		handlerError(resp,"Ooops, Something went wrong")
		fmt.Println("[ShipsHanlder.Update] error when unmarshall ships json : " + err.Error())
		return
	}

	artist:= a.artistUsecase.UpdateArtist(valId,&data)
	handlerSuccess(resp,artist)
}

func (a *ArtistHandler) DeleteArtist(resp http.ResponseWriter, req *http.Request)  {
	val := mux.Vars(req)
	valID,err := strconv.Atoi(val["id"])
	if err != nil{
		handlerError(resp,"ID harus angka")
		return
	}

	artist := a.artistUsecase.DeleteArtist(valID)
	handlerSuccess(resp,artist)
}


