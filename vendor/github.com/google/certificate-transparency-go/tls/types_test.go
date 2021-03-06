package tls

import "testing"

func TestHashAlgorithmString(t *testing.T) {
	var tests = []struct {
		algo HashAlgorithm
		want string
	}{
		{None, "None"},
		{MD5, "MD5"},
		{SHA1, "SHA1"},
		{SHA224, "SHA224"},
		{SHA256, "SHA256"},
		{SHA384, "SHA384"},
		{SHA512, "SHA512"},
		{99, "UNKNOWN(99)"},
	}
	for _, test := range tests {
		if got := test.algo.String(); got != test.want {
			t.Errorf("%v.String()=%q; want %q", test.algo, got, test.want)
		}
	}
}

func TestSignatureAlgorithmString(t *testing.T) {
	var tests = []struct {
		algo SignatureAlgorithm
		want string
	}{
		{Anonymous, "Anonymous"},
		{RSA, "RSA"},
		{DSA, "DSA"},
		{ECDSA, "ECDSA"},
		{99, "UNKNOWN(99)"},
	}
	for _, test := range tests {
		if got := test.algo.String(); got != test.want {
			t.Errorf("%v.String()=%q; want %q", test.algo, got, test.want)
		}
	}
}

func TestDigitallySignedString(t *testing.T) {
	var tests = []struct {
		ds   DigitallySigned
		want string
	}{
		{
			ds:   DigitallySigned{Algorithm: SignatureAndHashAlgorithm{Hash: SHA1, Signature: RSA}, Signature: []byte{0x01, 0x02}},
			want: "Signature: HashAlgo=SHA1 SignAlgo=RSA Value=0102",
		},
		{
			ds:   DigitallySigned{Algorithm: SignatureAndHashAlgorithm{Hash: 99, Signature: 99}, Signature: []byte{0x03, 0x04}},
			want: "Signature: HashAlgo=UNKNOWN(99) SignAlgo=UNKNOWN(99) Value=0304",
		},
	}
	for _, test := range tests {
		if got := test.ds.String(); got != test.want {
			t.Errorf("%v.String()=%q; want %q", test.ds, got, test.want)
		}
	}
}
