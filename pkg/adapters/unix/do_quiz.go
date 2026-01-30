package unix

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "math/rand"
	"net/http"
	// "time"
)


func DoQuiz(token string, lessonId int) (error){
  quiz, err := GetQuiz(token ,lessonId)
  if err != nil{
    return err
  }
  if err = startQuiz(token, quiz.Id); err != nil{
    return err
  }
  // time.Sleep(time.Second * time.Duration(rand.Intn(3000)))
  if err = finishQuiz(token, quiz.Id); err != nil{
    return err
  }
  
  
  return nil
}
func finishQuiz(token string, quizId int) (error){

  /* TODO
    https://uni-x.almv.kz/api/quizes/%d/check returns correct answers --> reSend with correct answers
  */
  body, err := json.Marshal(map[string]any{ 
    "answers": []string{},
    "tabSwitch": 0,
  })
  if err != nil{
    return err
  }
  
  req, err := http.NewRequest("POST",
      fmt.Sprintf("https://uni-x.almv.kz/api/quizes/%d/check", quizId), bytes.NewBuffer(body))
  if err != nil{
    return err
  }
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
  req.Header.Add("Content-Type", "application/json")
  resp, err := http.DefaultClient.Do(req)
  if err != nil{
    return err
  }
  if resp.StatusCode != http.StatusCreated{
    return ErrQuizFinish
  }
    
  return nil

}
func startQuiz(token string, quizId int) (error){
  body, err := json.Marshal(map[string]int{
    "quizId": quizId,
  })
  if err != nil{
    return err
  }
  
  req, err := http.NewRequest("POST",
      "https://uni-x.almv.kz/api/quizes-start-time", bytes.NewBuffer(body))
  if err != nil{
    return err
  }
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
  req.Header.Add("Content-Type", "application/json")
  resp, err := http.DefaultClient.Do(req)
  
  if err != nil{
    return err
  }
  if resp.StatusCode != http.StatusCreated{
    return ErrQuizStart
  }
  return nil
}

func GetQuiz(token string, lessonId int) (Quiz, error){
  req, err := http.NewRequest("GET",
        fmt.Sprintf("https://uni-x.almv.kz/api/lessons/%d/quiz", lessonId),nil)
  if err != nil{
    return Quiz{}, err
  }
  jsonResp := Quiz{}
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
  resp, err := http.DefaultClient.Do(req)
  if err != nil{
    return Quiz{}, err
  }
  if resp.StatusCode != http.StatusOK{
    return Quiz{}, ErrInvalidToken
  }
  defer resp.Body.Close()
  if err = json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil{
    return Quiz{}, err
  }

  return jsonResp, nil
  
  
  
  
}
