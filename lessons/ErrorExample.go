package lessons

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func ErrorExample() {
	err := errors.New("sfd")
	fmt.Println("error1:", err)
	err2 := fmt.Errorf("math: %g cannot be divided by zero", 8)
	fmt.Println("error2:", err2)
	_, err = ioutil.TempDir("", "temp")
	if err != nil {
		fmt.Println(fmt.Errorf("failed to create temp dir: %v", err))
	}
}
