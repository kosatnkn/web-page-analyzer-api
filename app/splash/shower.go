package splash

import (
	"fmt"

	"github.com/kosatnkn/web-page-analyzer-api/app/config"
	"github.com/kosatnkn/web-page-analyzer-api/metadata"
)

// Show a splash screen in one of several types.
func Show(style string, cfg *config.Config) {
	fmt.Print(style)
	fmt.Println("---")
	fmt.Print(metadata.BaseInfo())
	fmt.Print(metadata.BuildInfo())
	fmt.Println("---")
}
