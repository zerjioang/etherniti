package string

import (
	"strings"
	"testing"
)

func BenchmarkString(b *testing.B) {
	b.Run("create-empty", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = New()
		}
	})
	b.Run("create-with-data", func(b *testing.B) {
		example := []byte("foo-bar-hello-world")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewWith(example)
		}
	})
	b.Run("last-index", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = strings.LastIndex(example, "h")
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			search := []byte("h")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.LastIndex(search)
			}
		})
	})
	b.Run("to-bytes", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = []byte(example)
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				_ = exampleString.Bytes()
			}
		})
	})
	b.Run("chart-at", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = string(example[1])
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.CharAt(1)
			}
		})
	})

	b.Run("length", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = len(example)
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				_ = exampleString.Length()
			}
		})
	})

	b.Run("is-empty", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = len(example) == 0
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				_ = exampleString.IsEmpty()
			}
		})
	})

	b.Run("to-lowercase", func(b *testing.B) {
		b.Run("standard", func(b *testing.B) {
			example := "foo-bar-hello-world"
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				strings.ToLower(example)
			}
		})
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.LowerCase()
			}
		})
	})

	b.Run("reverse", func(b *testing.B) {
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.Reverse()
			}
		})
	})

	b.Run("title-case", func(b *testing.B) {
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.TitleCase()
			}
		})
	})

	b.Run("count-byte-match", func(b *testing.B) {
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				exampleString.CountByte([]byte("o")[0])
			}
		})
	})
	b.Run("contains", func(b *testing.B) {
		b.Run("custom", func(b *testing.B) {
			example := []byte("foo-bar-hello-world")
			item := []byte("world")
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			exampleString := NewWith(example)
			for n := 0; n < b.N; n++ {
				_ = exampleString.Contains(item)
			}
		})
	})
	b.Run("generate-uintptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = uintptr(n)
		}
	})

}
