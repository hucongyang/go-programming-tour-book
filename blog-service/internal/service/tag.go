package service

type CountTagRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binging:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binging:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binging:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binging:"required,gte=1"`
	Name       string `form:"name" binging:"min=3,max=100"`
	State      uint8  `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binging:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binging:"required,gte=1"`
}
