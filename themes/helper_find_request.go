package themes

import "github.com/wolacinio/postmanerator/postman"

func helperFindRequest(req postman.Request, name string) *postman.Response {
	for _, res := range req.Responses {
		if res.Name == name {
			return &res
		}
	}
	return nil
}
