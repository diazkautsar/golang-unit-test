package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWorld2(b *testing.B) {
	b.Run("Djalal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("DJALAL")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("DJALAL")
	}
}

func TestHelloWorldWithTableTest(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "TEST DJALAL",
			request:  "Djalal",
			expected: "Hello World Djalal",
		},
		{
			name:     "TEST DADANG",
			request:  "Dadang",
			expected: "Hello World Dadang",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, test.name)
		})
	}
}

// INI UNTUK SUBTEST. MUNGKIN KALO DI NODE JEST SEPERTI it()
func TestSubTest(t *testing.T) {
	t.Run("PERCOBAAN_1", func(t *testing.T) {
		result := HelloWorld("Djalal")
		assert.Equal(t, "Hello World Djalal", result, "Result Must be Hello World Djalal")
	})

	t.Run("PERCOBAAN_2", func(t *testing.T) {
		result := HelloWorld("Dadang")
		assert.Equal(t, "Hello World Dadang", result, "Result Must be Hello World Djalal")
	})
}

// BEFORE AND AFTER UNIT TEST
func TestMain(m *testing.M) {
	fmt.Println("BEFORE") // Bisa open database
	m.Run()               // untuk menjalankan semua unit test
	fmt.Println("AFTER")  // Bisa close database
}

// for skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on macos")
	}

	result := HelloWorld("Djalal")
	assert.Equal(t, "Hello World Djalal", result, "Result Must be Hello World Djalal")
}

// ini menggunakan assert yang mana ini memanggil fail jika error
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Djalal")
	assert.Equal(t, "Hello World Djalal", result, "Result Must be Hello World Djalal")
	fmt.Println("Test Hello World Assert Done")
}

// ini menggunakan require yang mana ini memanggil failNow jika error
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Djalal")
	require.Equal(t, "Hello World Djalal", result, "Result Must be Hello World Djalal")
	fmt.Println("Test Hello World Require Done")
}

func TestHelloWorldDjalal(t *testing.T) {
	result := HelloWorld("Djalal")
	if result != "Hello World Djalal" {
		//t.Fail()
		t.Fatal("Harusnyaa Hello World Djalal")
	}

	// ketika using fail. code ini akan tetap kepanggil
	fmt.Println("TestHelloWorldDjalal Kepanggil")
}

func TestHelloWorldDadang(t *testing.T) {
	result := HelloWorld("Dadang")
	if result != "Hello World Dadang" {
		//t.FailNow()
		t.Fatal("Harusnya Hello World Dadang")
	}

	// ketika using FailNow() code dibawah ini tidak kepanggil
	fmt.Println("INI CODE DIBAWAHNYA BRO")
}
