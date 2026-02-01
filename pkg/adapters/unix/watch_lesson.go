package unix

import (
	"bytes"
	"encoding/json"
	"strings"
	"time"
	"fmt"
	"net/http"
)


const (
	csrfURL  string = "https://uni-x.almv.kz/api/validates/csrf"
	eventURL string = "https://uni-x.almv.kz/api/validates/watched"
)


func WatchLesson(token string, lessonId int) error {
	preBody := map[string]any{
		"event":           "video-start", 
		"lessonId":        lessonId,
		"currentSpeed":    1,
		"clientTimestamp": time.Now().UTC().Format(time.RFC3339Nano),
		"lang":            "EN",
	}
	lesson, err := getLessonInfo(token, lessonId)
	if err != nil{
		return err
	}
	
	if preBody["token"], err = doEventReq(preBody, token, "video-start"); err != nil {
		return err
	}
	time.Sleep(time.Second * time.Duration(lesson.VideoDurationEn))
	
	if preBody["token"], err = doEventReq(preBody, token, "reached"); err != nil {
		return err
	}
	if preBody["token"], err = doEventReq(preBody, token, "video-end"); err != nil{
		return err
	}
	cookie, err := doCookieReq(token)
	if err != nil{
		return err
	}
	_, err = doFinishReq(token, preBody["token"].(string), cookie, &lesson)
	if err != nil{
		return err
	}
	return nil

}


func doFinishReq(authToken, token, cookie string, lesson *Lesson) (map[string]any, error){
	b := map[string]any{
		"videoDuration": lesson.VideoDurationEn,
		"videoWatched": lesson.VideoDurationEn + 2,
		"token": token,
	}
	body, err := json.Marshal(b)
	if err != nil{
		return map[string]any{}, nil
	}
	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://uni-x.almv.kz/api/lessons/%d/watched", lesson.Id), bytes.NewBuffer(body),)
	if err != nil{
		return map[string]any{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return map[string]any{}, err
	}
	if resp.StatusCode != http.StatusCreated{
		return map[string]any{}, ErrInvalidCSRF
	}
	defer resp.Body.Close()
	jsonResponse := make(map[string]any)
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil{
		return map[string]any{}, err 
	}
	return jsonResponse, nil
}

func doCookieReq(token string) (string, error){
	req, err := http.NewRequest("POST", csrfURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return "", err
	}
	if resp.StatusCode != http.StatusCreated{

		return "", ErrInvalidCSRF 
	}
	defer resp.Body.Close()
	return strings.Split(resp.Header["Set-Cookie"][0], ";")[0], nil


	
} 
func doEventReq(preBody map[string]any, token, event string) (string, error){
	preBody["clientTimestamp"] = time.Now().UTC().Format(time.RFC3339Nano)
	preBody["event"] = event
	startBody, err := json.Marshal(preBody)
	if err != nil {
		return "", err
	}
	jsonResp := make(map[string]any)

	req, err := http.NewRequest("POST", eventURL, bytes.NewBuffer(startBody))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", ErrInvalidToken
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		return "", err
	}
	return jsonResp["token"].(string), nil
	
}
func getLessonInfo(token string, lessonId int) (Lesson, error){
	lesson := Lesson{}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://uni-x.almv.kz/api/lessons/%d", lessonId), nil,
	)
	if err != nil{
		return Lesson{}, err
	}
	req.Header.Add(
		"Authorization", fmt.Sprintf("Bearer %s", token),
	)
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return Lesson{}, err
	}
	if resp.StatusCode != http.StatusOK{
		return Lesson{}, ErrInvalidToken 
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&lesson)
	if err != nil{
		return Lesson{}, err
	}
	
	return lesson, nil
}
