package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unibot/internal/services"
	"unibot/internal/repository"
	"unibot/pkg/adapters/unix"

	"gopkg.in/telebot.v4"
)


type topic struct {
	u repository.UserRepository
	n *services.NotificationService
}

type lessonProgressInfo struct {
	timeStatus string
	title string
	bar string
	isFinished bool
}

func makeMsg(l []lessonProgressInfo) string {
	msg := ""
	for i, lp := range l {
		if lp.isFinished {
			msg += fmt.Sprintf("%d. %s ✅\n", i + 1, lp.title)
		} else {
			msg += fmt.Sprintf("%d. %s\n%s %s\n", i + 1, lp.title, lp.bar, lp.timeStatus)
		}
	}
	return msg
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
		if errors.Is(err, ErrUnreachableTopic){
			return c.Send("Can't watch that topic\nNote that you can only watch topic if the previous ones were watched")
		}
		return err
	}
	if topic.IsPass {
		return c.Send("This topic has already been done")
	}
	go t.watchTopic(token, topic, c.Message())
	return nil
}

func (t topic) watchTopic(token string, topic unix.Topic, msgInstance *telebot.Message) {
	
	progress := make([]lessonProgressInfo, topic.LessonCount)
	
	for i, lesson := range topic.Lessons {
		progress[i].title = lesson.Title
		progress[i].bar = makeProgressBar(time.Second * 0, time.Duration(lesson.VideoDurationEn + 2) * time.Second)

		progress[i].timeStatus = fmt.Sprintf("00:00/%d:%d", lesson.VideoDurationEn / 60, lesson.VideoDurationEn - (lesson.VideoDurationEn / 60) * 60)
	}
	
	for i, lesson := range topic.Lessons {
		if len(lesson.LessonProgress) > 0 && lesson.LessonProgress[0].IsPassed {
			progress[i].isFinished = true
			continue
		}
		if len(lesson.LessonProgress) > 0 && !lesson.LessonProgress[0].IsPassed {
			unix.DoQuiz(token, lesson.Id)
			progress[i].isFinished = true
			continue
		}
		ctx, cancel := context.WithTimeout(context.Background(),time.Duration(lesson.VideoDurationEn + 2) * time.Second)
		defer cancel()
		
		
		go t.notificateWithInterval(ctx,
			 msgInstance,
			 time.Second * time.Duration(lesson.VideoDurationEn + 2),
			 progress, i,
		) 
		
		err := unix.WatchLesson(token, lesson.Id)
		if err != nil {
			cancel()
			t.n.Update("internal error", msgInstance)
			log.Print(msgInstance, err)
			return 
		}
		err = unix.DoQuiz(token, lesson.Id)
		if err != nil {
			cancel()
			t.n.Update("internal error", msgInstance)
			log.Print(msgInstance, err)
			return 
		}

	}
	t.n.Notificate(fmt.Sprintf("Finished %s", topic.Title), msgInstance.Chat.ID)
}



func (t topic) notificateWithInterval(ctx context.Context, editable *telebot.Message, duration time.Duration, msg []lessonProgressInfo, cur int) {
	t.n.Update(makeMsg(msg), editable)

	ticker := time.NewTicker(duration / 70)
	start := time.Now()
	for {
		select {
			case <-ctx.Done():
				msg[cur].isFinished = true
				t.n.Update(makeMsg(msg), editable)
				return
			case <-ticker.C:
				elapsed := time.Since(start)
				
				msg[cur].bar = makeProgressBar(elapsed, duration)
				msg[cur].timeStatus = fmt.Sprintf("%s/%s", convertToMMSS(elapsed), convertToMMSS(duration))
				t.n.Update(makeMsg(msg), editable)
		}
	}
	
}
func convertToMMSS(t time.Duration) string {
	minutes := int(math.Floor(t.Seconds() / 60))
	second := int(math.Floor(t.Seconds() - float64(minutes) * 60))
	return fmt.Sprintf("%02d:%02d", minutes, second)
}

func makeProgressBar(elapsed, duration time.Duration) string {
	
	percents := (elapsed.Seconds() / duration.Seconds()) * 30
	cur := "▓"
	left := "░"

		
	bar := strings.Repeat(cur, int(percents))
	bar += strings.Repeat(left, 30 - int(percents))
	return bar 
	
}

func NewTopic(u repository.UserRepository, n *services.NotificationService) topic {
	return topic{
		u: u,
		n: n,
	}
}

func getTopic(token string, courseID, topicID int) (unix.Topic, error) {
	course, err := unix.GetCourse(token, courseID)
	if err != nil {
		return unix.Topic{}, err
	}
	for _, topic := range course.Topics {
		if topic.Id == topicID {
			
			err := isReachable(course.Topics, topic)
			if err != nil {
				return unix.Topic{}, err
			}
			return topic, nil
		}
	}
	return unix.Topic{}, nil
}
func isReachable(topics []unix.Topic, topic unix.Topic) error {
	for _, t := range topics {
		if !t.IsPass && t.Order < topic.Order {
			return ErrUnreachableTopic
		}
	}
	return nil
}
