package auth

import (
	"testing"
)

func TestCompareHash(t *testing.T) {
	testCases := []struct {
		desc       string
		p          *params
		password   string
		RePassword string
		match      bool
	}{
		{
			desc: "Comparing Hash should Fail",
			p: &params{
				memory:      64 * 1024,
				iterations:  3,
				parallelism: 2,
				saltLength:  16,
				keyLength:   32,
			},
			password:   "test_password",
			RePassword: "test^password",
			match:      true,
		},
		{
			desc: "Comparing Hash should Pass",
			p: &params{
				memory:      64 * 1024,
				iterations:  3,
				parallelism: 2,
				saltLength:  16,
				keyLength:   32,
			},
			password:   "test^password",
			RePassword: "test^password",
			match:      false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			encodedHash, err := generateFromPassword(tC.password, tC.p)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(encodedHash)
			match, err := comparePasswordAndHash(tC.RePassword, encodedHash)
			if err != nil {
				t.Fatal(err)
			}
			if match == tC.match {
				t.Fail()
			}
		})
	}
}
