package services

import (
	"basic-api/config"
	"basic-api/entity"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

const URL string = "https://screening.moduit.id/"

func QuestionOne() (*entity.BodyResponse, error) {
	// set the URL
	theURL := URL + "backend/question/one"
	getData, errRequest := config.RequestURL(theURL)
	if errRequest != nil {
		fmt.Println(errRequest.Error())
		return &entity.BodyResponse{}, errRequest
	}
	jsonData := []byte(getData)
	// define the data
	var data entity.BodyResponse
	// unmarshal the json data
	var errUnMarshal = json.Unmarshal(jsonData, &data)
	if errUnMarshal != nil {
		fmt.Println(errUnMarshal.Error())
		return &entity.BodyResponse{}, errUnMarshal
	}

	return &entity.BodyResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Footer:      data.Footer,
		Tags:        data.Tags,
		CreatedAt:   data.CreatedAt,
	}, nil
}

func QuestionTwo() ([]*entity.BodyResponse, error) {
	// set the URL
	theURL := URL + "backend/question/two"
	getData, errRequest := config.RequestURL(theURL)
	if errRequest != nil {
		fmt.Println(errRequest.Error())
		return []*entity.BodyResponse{}, errRequest
	}
	jsonData := []byte(getData)
	// define the data
	var data []entity.BodyResponse
	// unmarshal the json data
	var errUnMarshal = json.Unmarshal(jsonData, &data)
	if errUnMarshal != nil {
		fmt.Println(errUnMarshal.Error())
		return []*entity.BodyResponse{}, errUnMarshal
	}
	// filter the data
	var result []entity.BodyResponse
	for _, val := range data {
		if strings.Contains(val.Description, "Ergonomic") || strings.Contains(val.Title, "Ergonomic") {
			if len(val.Tags) > 0 && config.Contains(val.Tags, "Sports") {
				result = append(result, val)
			}
		}
	}
	// sort the data
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID > result[j].ID
	})
	// get only 3 data
	result2 := make([]*entity.BodyResponse, 3)
	for i := 0; i < 3; i++ {
		result2[i] = &result[i]
	}
	return result2, nil
}

func QuestionThree() ([]*entity.BodyResponseForQuestionThree, error) {
	// set the URL
	theURL := URL + "backend/question/three"
	getData, errRequest := config.RequestURL(theURL)
	if errRequest != nil {
		fmt.Println(errRequest.Error())
		return []*entity.BodyResponseForQuestionThree{}, errRequest
	}
	jsonData := []byte(getData)
	// define the data
	var data []entity.BodyResponseForQuestionThree
	// unmarshal the json data
	var errUnMarshal = json.Unmarshal(jsonData, &data)
	if errUnMarshal != nil {
		fmt.Println(errUnMarshal.Error())
		return []*entity.BodyResponseForQuestionThree{}, errUnMarshal
	}
	// filter the data
	var result []*entity.BodyResponseForQuestionThree

	for _, val := range data {
		for _, v := range val.Items {
			result = append(result, &entity.BodyResponseForQuestionThree{
				ID:          val.ID,
				Category:    val.Category,
				Title:       v.Title,
				Description: v.Description,
				Footer:      v.Footer,
				CreatedAt:   val.CreatedAt,
			})
		}
	}

	return result, nil
}
