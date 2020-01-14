package commons

import "testing"

import "strings"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestBase64Encode(t *testing.T) {
	want := "SG9sYSBNdW5kbw"
	got := Base64Encode("Hola Mundo")
	if got != want {
		t.Errorf("Base64Encode() = %q, want %q", got, want)
	}
}

func TestBase64Decode(t *testing.T) {
	want := "Decodificado"
	got, _ := Base64Decode("RGVjb2RpZmljYWRv")
	if got != want {
		t.Errorf("Base64Decode() = %q, want %q", got, want)
	}

	want = "Decoding Error illegal base64 data at input byte 16"
	_, err := Base64Decode("RGVjb2RpZmljYWRvñ")
	if err.Error() != want {
		t.Errorf("Base64Decode() = %q, want %q", err, want)
	}
}

func TestGetDataJWT(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyLCJpZFVzZXIiOiIxMjMxMiJ9.5cqpNmlEcroWY7w-xj-zkO9p7Vz2-VaOFUOPvOblWzc"
	want := "12312"

	value, _ := GetDataJWT(token, "idUser")
	if value != want {
		t.Errorf("TestGetDataJWT() = %q, want %q", value, want)
	}

	value, err := GetDataJWT("wrong-token", "idUser")
	want = "Invalid JWT"
	if err.Error() != want {
		t.Errorf("TestGetDataJWT() = %q, want %q", err, want)
	}

	value, err = GetDataJWT("wrong.token.jwt", "idUser")
	want = "Invalid JWT"
	if !strings.Contains(err.Error(), want) {
		t.Errorf("TestGetDataJWT() = %q, want %q", err, want)
	}
}
