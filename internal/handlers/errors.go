package handlers

import "errors"
var (
  ErrUnreachableTopic = errors.New("can't reach the given topic") // if topic.prev isnt watched
) 
