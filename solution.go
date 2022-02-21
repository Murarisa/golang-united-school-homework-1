package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	helloworld := emoji.Sprint("Hello :world:!")
	return helloworld
}
