package unix

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCourse(token string, courseId int) (Course, error) {
	course := Course{}

	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://uni-x.almv.kz/api/courses/%d", courseId),
		nil,
	)
	if err != nil {
		return Course{}, err
	}

	req.Header.Add(
		"Authorization", fmt.Sprintf("Bearer %s", token),
	)
	req.Header.Add(
		"Content-Type","application/json",
	)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return Course{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return Course{}, ErrInvalidToken
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&course)
	
	return course, nil
}

func GetCourses(token string, moduleID int) ([]ModuleCourse, error){
	modules, err := GetModules(token)
	if err != nil{
		return nil, err
	}
	for _, module := range modules{
		if module.Id == moduleID{
			return module.Courses, nil
		}
	}
	return nil, nil
}

