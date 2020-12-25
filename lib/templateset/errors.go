package templateset

import (
	"fmt"
)

type TemplateSetError struct {
	Err error
}

func (tse *TemplateSetError) Error() string {
	return fmt.Sprintf("Error in templateset. %s", tse.Err)
}
