// Package millipede provides a framework for creating millipedes.
// millipede is designed to be easy to understand and write, the most simple
// application can be written as follow:
//   func main() {
//     fmt.Println(millipede.New(20))
//   }
package millipede

import (
	"strings"
	"unicode/utf8"

	"github.com/Sirupsen/logrus"
)

type Millipede struct {
	// Size is the amount of feet pairs
	Size uint64

	// Reverse is the flag that indicates the direction (up/down)
	Reverse bool

	// Skin is the current millipede skin (template)
	Skin string

	// Opposite is the flag that indicates the direction (left/right)
	Opposite bool

	// Width is the width of the millipede (depending on its age and the food it consumes)
	Width uint64

	// Curve is the size of the curve
	Curve uint64
}

type Skin struct {
	// Head is used by the millipede to think about its life
	Head string
	// Pede are what make this arthropod so special
	Pede string

	// Reverse is the reverse skin of the millipede
	Reverse *Skin
}

// String returns a string representing a millipede
func (m *Millipede) String() string {
	// --curve support
	paddingOffsets := []string{""}
	if m.Curve > 0 {
		for n := uint64(1); n < m.Curve+1; n++ {
			paddingOffsets = append(paddingOffsets, strings.Repeat(" ", int(n)))
		}
		for n := m.Curve - 1; n > 0; n-- {
			paddingOffsets = append(paddingOffsets, strings.Repeat(" ", int(n)))
		}
	}

	// --opposite support
	if m.Opposite {
		paddingOffsets = append(paddingOffsets[m.Curve:], paddingOffsets[:m.Curve]...)
	}

	skins := map[string]Skin{
		"default": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(███)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(███)═╗",
			},
		},
		"frozen": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(❄❄❄)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(❄❄❄)═╗",
			},
		},
		"corporate": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(©©©)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(©©©)═╗",
			},
		},
		"musician": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(♫♩♬)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(♫♩♬)═╗",
			},
		},
		"bocal": {
			Head: "  ╚⊙ ⊙╝  ",
			Pede: "╚═(🐟🐟🐟)═╝",
			Reverse: &Skin{
				Head: "  ╔⊙ ⊙╗  ",
				Pede: "╔═(🐟🐟🐟)═╗",
			},
		},
		"ascii": {
			Head: "  \\o o/  ",
			Pede: "|=(###)=|",
			Reverse: &Skin{
				Head: "  /o o\\  ",
				Pede: "|=(###)=|",
			},
		},
	}

	// --skin support
	skin := skins[m.Skin]
	if skin.Head == "" {
		logrus.Fatalf("no such skin: '%s'", m.Skin)
	}

	// --reverse support
	if m.Reverse && skin.Reverse != nil && skin.Reverse.Head != "" {
		skin = *skin.Reverse
	}

	// --width support
	if m.Width < 3 {
		logrus.Fatalf("millipede cannot have a witch < 3")
	}
	if m.Width > 3 {
		w := utf8.RuneCountInString(skin.Head)
		head := StringToRuneSlice(skin.Head)
		skin.Head = string(head[:w/2]) + strings.Repeat(string(head[w/2:w/2+1]), int(m.Width-2)) + string(head[w/2+1:])
		pede := StringToRuneSlice(skin.Pede)
		skin.Pede = string(pede[:w/2]) + strings.Repeat(string(pede[w/2:w/2+1]), int(m.Width-2)) + string(pede[w/2+1:])
	}

	// build the millipede body
	body := []string{paddingOffsets[0] + skin.Head}
	var x uint64
	for x = 0; x < m.Size; x++ {
		var line string
		if m.Curve > 0 {
			line = paddingOffsets[x%(m.Curve*2)] + skin.Pede
		} else {
			line = "" + skin.Pede
		}
		body = append(body, line)
	}

	// --reverse support
	if m.Reverse {
		for i, j := 0, len(body)-1; i < j; i, j = i+1, j-1 {
			body[i], body[j] = body[j], body[i]
		}
	}

	return strings.Join(body, "\n")
}

// New returns a millipede
func New(size uint64) *Millipede {
	return &Millipede{
		Size:     size,
		Reverse:  false,
		Skin:     "default",
		Opposite: false,
		Width:    3,
		Curve:    4,
	}
}

// StringToRuneSlice converts a string to a slice of runes
func StringToRuneSlice(input string) []rune {
	output := make([]rune, utf8.RuneCountInString(input))
	n := 0
	for _, r := range input {
		output[n] = r
		n++
	}
	return output
}
