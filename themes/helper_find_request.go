package themes

import "github.com/wolacinio/postmanerator/postman"

func helperFindRequest(req postman.Request, name string) *postman.Request {
	for _, res := range req.Request {
			return &res
	}
	return nil
}
