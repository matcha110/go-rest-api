package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
		// validation.Field(
		// 	&task.Deadline,
		// 	validation.NilOrNotEmpty.Error("deadline must be either set or empty"),
		// 	validation.When(task.Deadline != nil, validation.By(func(value interface{}) error {
		// 		if (*value.(*time.Time)).Before(time.Now()) {
		// 			return validation.NewError("validation_deadline_past", "deadline cannot be in the past")
		// 		}
		// 		return nil
		// 	})),
		// ),
	)
}
