package structs

import (
	"fmt"
	"reflect"
)

// Phone  it descirbe smartPhone contains some properties.
type Phone struct {
	Brand        string
	Price        float32
	Version      string
	SystemVesion string
}

// IosPhone  the function is the information of object Phone
func IosPhone(ver string) Phone {
	apple := Phone{
		Brand:        "apple",
		Price:        8675,
		Version:      ver,
		SystemVesion: "ios14.4",
	}
	return apple
}

func main() {
	fmt.Printf("phone:%+v\n", IosPhone("iPhone 12"))
	fmt.Println("================================")
	iphoneX := Phone{
		Version:      "iPhone X",
		SystemVesion: "ios14.4",
	}

	iphone12 := Phone{
		Version:      "iPhone 12",
		SystemVesion: "ios14.4",
	}
	fmt.Println(reflect.TypeOf(iphoneX))
	fmt.Println(reflect.TypeOf(iphone12))
}
