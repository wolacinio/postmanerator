package themes

import "github.com/wolacinio/postmanerator/postman"

func helperFindRequest(req postman.Request, name string) *postman.Response.Requ {
	for _, res := range req.Responses.Requ {
			return &res
	}
	return nil
}
