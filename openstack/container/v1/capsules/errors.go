package capsules

import (
	"fmt"

	"github.com/chjlangzi/gophercloud"
)

type ErrInvalidDataFormat struct {
	gophercloud.BaseError
}

func (e ErrInvalidDataFormat) Error() string {
	return fmt.Sprintf("Data in neither json nor yaml format.")
}
