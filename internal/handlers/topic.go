package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unibot/internal/services"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)

type topic struct {
	u *services.UserService
}

func (t topic) Handle(c telebot.Context) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	user, err := t.u.GetUser(ctx, c.Chat().ID)
	if err != nil {
		return err
	}

	token, err := unix.Login(user.Login, user.Password)
	if err != nil {
		return err
	}

	topicID, err := strconv.Atoi(strings.Split(c.Callback().Data, ":")[0])
	if err != nil {
		return err
	}
	courseID, err := strconv.Atoi(strings.Split(c.Callback().Data, ":")[1])
	if err != nil {
		return err
	}

	topic, err := getTopic(token, courseID, topicID)
	if err != nil {
		return err
	}
	go t.watchTopic(token, topic, c)
	return c.Send("OK")
}

func (t topic) watchTopic(token string, topic unix.Topic, c telebot.Context) {
	for _, lesson := range topic.Lessons {
		if len(lesson.LessonProgress) > 0 && lesson.LessonProgress[0].IsPassed {
			continue
		}
		if len(lesson.LessonProgress) > 0 && !lesson.LessonProgress[0].IsPassed {
			unix.DoQuiz(token, lesson.Id)
			continue
		}
		c.Send(fmt.Sprintf("started %s", lesson.Title))
		err := unix.WatchLesson(token, lesson.Id)
		if err != nil {
			return
		}
		err = unix.DoQuiz(token, lesson.Id)
		if err != nil {
			return
		}
		c.Send(fmt.Sprintf("finished %s", lesson.Title))

	}
}

func NewTopic(u *services.UserService) topic {
	return topic{
		u: u,
	}
}

func getTopic(token string, courseID, topicID int) (unix.Topic, error) {
	course, err := unix.GetCourse(token, courseID)
	if err != nil {
		return unix.Topic{}, err
	}
	for _, topic := range course.Topics {
		if topic.Id == topicID {
			return topic, nil
		}
	}
	return unix.Topic{}, nil
}
