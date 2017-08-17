// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"fmt"
	"strings"
)

const (
	// ColorBlack is a Color of type Black
	ColorBlack Color = iota
	// ColorWhite is a Color of type White
	ColorWhite
	// ColorRed is a Color of type Red
	ColorRed
	// ColorGreen is a Color of type Green
	ColorGreen Color = iota + 30
	// ColorBlue is a Color of type Blue
	ColorBlue Color = iota + 30
	// ColorGrey is a Color of type Grey
	ColorGrey Color = iota + 30
	// ColorYellow is a Color of type Yellow
	ColorYellow Color = iota + 30
)

const _ColorName = "BlackWhiteRedGreenBlueGreyYellow"

var _ColorMap = map[Color]string{
	0:  _ColorName[0:5],
	1:  _ColorName[5:10],
	2:  _ColorName[10:13],
	33: _ColorName[13:18],
	34: _ColorName[18:22],
	35: _ColorName[22:26],
	36: _ColorName[26:32],
}

func (i Color) String() string {
	if str, ok := _ColorMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Color(%d)", i)
}

var _ColorValue = map[string]Color{
	_ColorName[0:5]:                    0,
	strings.ToLower(_ColorName[0:5]):   0,
	_ColorName[5:10]:                   1,
	strings.ToLower(_ColorName[5:10]):  1,
	_ColorName[10:13]:                  2,
	strings.ToLower(_ColorName[10:13]): 2,
	_ColorName[13:18]:                  33,
	strings.ToLower(_ColorName[13:18]): 33,
	_ColorName[18:22]:                  34,
	strings.ToLower(_ColorName[18:22]): 34,
	_ColorName[22:26]:                  35,
	strings.ToLower(_ColorName[22:26]): 35,
	_ColorName[26:32]:                  36,
	strings.ToLower(_ColorName[26:32]): 36,
}

// ParseColor attempts to convert a string to a Color
func ParseColor(name string) (Color, error) {
	if x, ok := _ColorValue[name]; ok {
		return Color(x), nil
	}
	return Color(0), fmt.Errorf("%s is not a valid Color", name)
}

func (x *Color) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Color) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseColor(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
