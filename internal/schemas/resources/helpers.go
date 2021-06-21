package resources

import (
	"crypto/sha256"
	"fmt"
)

func HashSum(contents interface{}) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(contents.(string))))
}
