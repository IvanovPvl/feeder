package feeder

import (
	"log"
	"os"

	_ "github.com/ivanovpvl/feeder/matchers"
	"github.com/ivanovpvl/feeder/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
