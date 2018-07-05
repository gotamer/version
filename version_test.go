package version

import (
	"log"
	"testing"
)

func TestRunMain(t *testing.T) {
	log.Printf("Version Long: %v\n", Version.Long)
	log.Printf("Git: %s\n", Version.Git)
	log.Printf("Commits: %d\n", Version.Commits)
	log.Printf("Tags: %s\n", Version.Tags)
	log.Printf("Dirty: %v\n", Version.Dirty)

}
