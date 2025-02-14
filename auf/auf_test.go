package auf

import "testing"

func Test_hash_password(t *testing.T) {
	str, err := HashPassword("password")
	if err != nil {
		t.Fatal("aaaa")
	}
	if str == "password" {
		t.Fatal("Хэш совпадает с паролем")
	}
}
