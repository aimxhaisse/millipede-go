package millipede

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	millipede := New(20)
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
}

func ExampleMillipede() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "default",
		Opposite: false,
		Width:    3,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
}

func ExampleMillipede_reverse() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  true,
		Skin:     "default",
		Opposite: false,
		Width:    3,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//  ╔═(███)═╗
	//   ╔═(███)═╗
	//    ╔═(███)═╗
	//     ╔═(███)═╗
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//  ╔═(███)═╗
	//   ╔═(███)═╗
	//    ╔═(███)═╗
	//     ╔═(███)═╗
	//    ╔═(███)═╗
	//   ╔═(███)═╗
	//  ╔═(███)═╗
	// ╔═(███)═╗
	//   ╔⊙ ⊙╗
}

func ExampleMillipede_opposite() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "default",
		Opposite: true,
		Width:    3,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//       ╚⊙ ⊙╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
}

func ExampleMillipede_skin() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "bocal",
		Opposite: false,
		Width:    3,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//     ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//     ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	// ╚═(🐟🐟🐟)═╝
	//  ╚═(🐟🐟🐟)═╝
	//   ╚═(🐟🐟🐟)═╝
	//    ╚═(🐟🐟🐟)═╝
}

func ExampleMillipede_width() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "default",
		Opposite: false,
		Width:    6,
		Curve:    4,
	}
	fmt.Println(millipede)
	// Output:
	//   ╚⊙    ⊙╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
	//     ╚═(██████)═╝
	//    ╚═(██████)═╝
	//   ╚═(██████)═╝
	//  ╚═(██████)═╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
	//     ╚═(██████)═╝
	//    ╚═(██████)═╝
	//   ╚═(██████)═╝
	//  ╚═(██████)═╝
	// ╚═(██████)═╝
	//  ╚═(██████)═╝
	//   ╚═(██████)═╝
	//    ╚═(██████)═╝
}

func ExampleMillipede_curve() {
	millipede := &Millipede{
		Size:     20,
		Reverse:  false,
		Skin:     "default",
		Opposite: false,
		Width:    3,
		Curve:    6,
	}
	fmt.Println(millipede)
	// Output:
	//   ╚⊙ ⊙╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//      ╚═(███)═╝
	//       ╚═(███)═╝
	//      ╚═(███)═╝
	//     ╚═(███)═╝
	//    ╚═(███)═╝
	//   ╚═(███)═╝
	//  ╚═(███)═╝
	// ╚═(███)═╝
	//  ╚═(███)═╝
	//   ╚═(███)═╝
	//    ╚═(███)═╝
	//     ╚═(███)═╝
	//      ╚═(███)═╝
	//       ╚═(███)═╝
	//      ╚═(███)═╝
}

func TestMillipede_zalgo(t *testing.T) {
	// FIXME: find a better test
	millipede := New(20)
	millipede.Zalgo = true
	millipede.String()
}

func TestMillipede_small_width(t *testing.T) {
	millipede := New(20)
	millipede.Width = 2
	// FIXME: check if it exits
	//millipede.String()
}

func ExampleNew_complex() {
	millipede := &Millipede{
		Size:     42,
		Reverse:  true,
		Skin:     "bocal",
		Opposite: true,
		Width:    6,
		Curve:    10,
	}
	fmt.Println(millipede)
	// Output:
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	// ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	// ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//  ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//   ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//    ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//     ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//      ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//       ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//        ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//         ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//          ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//           ╔═(🐟🐟🐟🐟🐟🐟)═╗
	//             ╔⊙    ⊙╗
}

func ExampleStringToRuneSlice() {
	input := "╔═(🐟🐟🐟🐟🐟🐟)═╗"
	output := StringToRuneSlice(input)
	fmt.Println(output)
	// Output: [9556 9552 40 128031 128031 128031 128031 128031 128031 41 9552 9559]
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(20)
	}
}

func BenchmarkRendering(b *testing.B) {
	millipede := New(20)
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", millipede)
	}
}
