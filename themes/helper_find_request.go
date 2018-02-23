package themes

import "github.com/wolacinio/postmanerator/postman"

func helperFindResponse(req postman.Request, name string) *postman.Response {
	for _, res := range req.Responses {
		if res.Name == name {
			return &res
		}
	}
	return nil
}
