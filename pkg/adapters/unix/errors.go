package unix

import "errors"


var(
  ErrInvalidCredentials = errors.New("invalid credentials")
  ErrInvalidToken = errors.New("invalid token")
  ErrQuizFinish = errors.New("quiz finish err")
  ErrQuizStart = errors.New("quiz start err")
  ErrInvalidCSRF = errors.New("invalid csrf")
)
