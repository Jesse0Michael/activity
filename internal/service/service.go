package service

type ActivitiesRequest struct {
	Source string `query:"source"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}
