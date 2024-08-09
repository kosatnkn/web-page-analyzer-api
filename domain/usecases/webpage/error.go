package webpage

import (
	"fmt"

	err "github.com/kosatnkn/web-page-analyzer-api/domain/errors"
)

func (s *WebPage) errNoWebPage(url string) error {
	return err.NewDomainError("1000", fmt.Sprintf("Web page not found at %s", url), nil)
}
