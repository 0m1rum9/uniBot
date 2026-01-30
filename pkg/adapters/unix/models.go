package unix

type ModulesResponse struct {
	Modules []Module `json:"modules"`
}

type Module struct {
	Id             int            `json:"id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	Type           string         `json:"type"`
	IsPublic       bool           `json:"isPublic"`
	StartDate      string         `json:"startDate"`
	OrganizationId int            `json:"organizationId"`
	ImageLink      string         `json:"imageLink"`
	Order          int            `json:"order"`
	CreatedAt      string         `json:"createdAt"`
	UpdatedAt      string         `json:"UpdatedAt"`
	Courses        []ModuleCourse `json:"courses"`
}

type ModuleCourse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`
	Type        string `json:"type"`
	IsPublic    bool   `json:"isPublic"`
	ModuleId    int    `json:"moduleId"`
	ImageLink   string `json:"imageLink"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"UpdatedAt"`
}

type Course struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Order       int     `json:"order"`
	Type        string  `json:"type"`
	IsPublic    bool    `json:"isPublic"`
	ModuleId    int     `json:"moduleId"`
	ImageLink   string  `json:"imageLink"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"UpdatedAt"`
	Topics      []Topic `json:"topics"`
}

type Topic struct {
	Id                  int      `json:"id"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	Order               int      `json:"order"`
	Deadline            string   `json:"deadline"`
	CourseId            int      `json:"courseId"`
	IsPublic            bool     `json:"isPublic"`
	Type                string   `json:"type"`
	Lessons             []Lesson `json:"lessons"`
	LessonCount         int      `json:"lessonCount"`
	QuizesTotalCount    int      `json:"quizesTotalCount"`
	LessonProgressesMap int      `json:"lessonProgressesMap"`
	QuizesPassedCount   int      `json:"quizesPassedCount"`
	IsPass              bool     `json:"isPass"`
	IsFinished          bool     `json:"isFinished"`
	ProgressCreatedAt   string   `json:"progressCreatedAt"`
	ProgressUpdateAt    string   `json:"progressUpdatedAt"`
	StatusSession       string   `json:"statusSession"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"UpdatedAt"`
}

type Lesson struct {
	Id              int            `json:"id"`
	Title           string         `json:"title"`
	TopicId         int            `json:"topicId"`
	VideoLink       string         `json:"videoLink"`
	VideoDurationEn int            `json:"videoDurationEn"`
	Order           int            `json:"order"`
	IsPublic        bool           `json:"isPublic"`
	CreatedAt       string         `json:"createdAt"`
	UpdatedAt       string         `json:"UpdatedAt"`
	Quizes          []Quiz         `json:"quizes"`
	LessonProgress  []lessonProgress `json:"lessonsProgressesMap"`
	// TODO
}

type Quiz struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LessonId    int    `json:"lessonId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"UpdatedAt"`
	Questions []Question `json:"questions"`
	IsWatched bool `json:"isWatched"`
	CurrentAttempt int `json:"currentAttempt"`
	IsStart bool `json:"isStart"`
	IsPassedQuiz bool `json:"isPassedQuiz"`
}
type Question struct{
	Id int `json:"id"`
	QuestionText string `json:"questionText"`
	QuizId int `json:"quizId"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"UpdatedAt"`
	Answers []Answer `json:"answers"`
	IsMultiple bool `json:"isMultiple"`
}
type Answer struct{
	AnswerText string `json:"answerText"`
	Id string `json:"id"`
}
type lessonProgress struct { // aka lessonsProgressesMap
	Id             int    `json:"id"`
	LessonId       int    `json:"lessonId"`
	UserId         int    `json:"userId"`
	IsWatched      bool   `json:"isWatched"`
	IsAdded        bool   `json:"isAdded"`
	VideoDuration  int    `json:"videoDuration"`
	VideoWatched   int    `json:"videoWatched"`
	IsPassed       bool   `json:"isPassed"`
	IsPassedOnTime bool   `json:"isPassedOnTime"`
	IsWrong        bool   `json:"isWrong"`
	IsWrongOnTime  bool   `json:"isWrongOnTime"`
	IsCheated      bool   `json:"isCheated"`
	IsValidWatched bool   `json:"isValidWatched"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"UpdatedAt"`
}
