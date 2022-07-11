package rsa_test

import (
	"bytes"
	"crypto/rsa"
	"encoding/pem"
	"io"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/IgnorantSapient/gokey"
)

func pemEqual(pem1, pem2 string, t *testing.T) bool {
	b1, err := pem.Decode([]byte(pem1))
	if b1 == nil {
		t.Fatalf("failed to decode %v: %v", pem1, err)
	}

	b2, err := pem.Decode([]byte(pem2))
	if b2 == nil {
		t.Fatalf("failed to decode %v: %v", pem2, err)
	}

	return b1.Type == b2.Type && bytes.Equal(b1.Bytes, b2.Bytes)
}

func getKey(kt gokey.KeyType, t *testing.T) string {
	priv, err := gokey.GetKey("pass", "example.com", nil, kt, true)
	if err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	err = gokey.EncodeToPem(priv, &b)
	if err != nil {
		t.Fatal(err)
	}

	return b.String()
}

// generated by "$ gokey -p pass -r example.com -u -t rsa2048" with Go < 1.11
var knownRsa2048 = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAqIT/jlQnJAIxS2h/AEz6RSiN9/GWY8lh9V81vfpzjZ069NOH
fVL0rQiRj11GqReHTUOjuOiCwYn6d6tPSCVC8fDyBjQ6cf9tm0G+lMf8LsMCPFTL
PDivqH8KwXLgC2+Zmm906nGkGUZ79AdOa6A2k07PbQ791/7s6ghWU6Dx6lPqiHC+
jFfQctQWNUdIchYt1X9eARNTcJQAYBAikH0X+dZwcm9sfTRtJCo7WvjOBwxIk/XA
k9GcoR2bbGrN8wnVGO2Luw9exg/GavfpkWI1jRx82Hfqfd+Ux3478yOS2Z1kbve1
g3DI8qscaN8oZugw4R/EBg0jOOinDwE9oISZ3wIDAQABAoIBAGJooE/+RokZmq29
jQSg7zl5sEYNV1RYYpMGkXyqh9Y37hjQefuueOGe8lm1D7Fo4wM0r6Qoa0sYByLg
8EBiOhDNMph64XJ2xgv3PZLmohawnFqc7b3yIGoWHjLPoZQsDJgJ5E2QJVL5PSNJ
LPteqOAnEqxOJ+B9pt4YFklp4DuE6H21pgM8tn0azH5Hf4w6LDrkjk4+WVAFrQqj
UJFSVVQ6Vq54BxnBUJzGCmlKCqDbUrOLQ2TrDodDVPHzmzxX69WVtwqtkiJ6X7HH
pnodlRV9LwLMkZwgQBEcoZnQkWLimGBRBaO/uuZewmXJ9u6HgXIL9+VVHFnkcdZs
S6QwLxECgYEA3ghowXxlR7GsC8+07YJUJV+KvKNK4qEZVp2MMscxOPWjcO1tWO88
8K/E54Gp8ktsqmejQjnUD7749h2pEUZHfy2sCpWOHjXzLouCkZlSZQpsWVcFkFLL
qZvwv+pCV9osxpf8o2n/ZuKGP3VgKeSlql6eHLdRV6edEUEJogwFgykCgYEAwkzR
d5uqToUlbz3fKTs9c4bTt+xnhlOyXPdnbRph/N96Yp0eIi2v7F0FBB5qzKZKy72M
4+DHvWrpAruh14nryXolQ/OesT27R00auxjOAWM2v2BFeNjBadwnlYfbzCtM+1Wd
6U+Ahd6Nkdk5tAkMvP03RwPxOTPXuXOtJNFgHccCgYEApbq1JqdJmegeuXpCXH5J
fNQB5KgmP19sYGCcw1I3hYKkiqhOVHHOlQE+AmegiMCPzeopzEcJ6O7tOhgNmF1T
BLT8k7HqMNaoO/fab/93pv/OvCjeeEm0x1ckrruW5ahuf5X296so/ozbFAbyzpJi
rfaOInUa/EiaTsSzAgfjHXECgYA13opsuPHc1zlrwCGxEsWU1Bq68YY6TdYzxDwe
maP1MhiCYsFKBJSz2Y5cd/pwRuKR3jnDrDx0ncGAinjyg0CmGYFfd5nV1iGoQQ5a
NSRYaiNxp8VbHe1x5iHraUFdN9weCry/RNWDSBLEDw/ahG/Nrf63Z7Znf6Idvp6Q
iKQ3XQKBgGpKBezvB82Jk0suZZdqZIs3OVDy7i6SIwS3MgABRE5pYgDWhlnEzGFS
Tk4zAHR2tIeN4RjOcanspxKbN+Qfm1TkeFwRq5553f0ypMz7xkCdUCKIAeCfTkbs
qkciBuCm/Oa4RlB/LwH+paBV3RvQOYLXw50/uzAr5S2+2wPx4wMT
-----END RSA PRIVATE KEY-----`

// generated by "$ gokey -p pass -r example.com -u -t rsa4096" with Go < 1.11
var knownRsa4096 = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

func TestKnownKey(t *testing.T) {
	rsa2048 := getKey(gokey.RSA2048, t)
	if !pemEqual(rsa2048, knownRsa2048, t) {
		t.Fatal("generated RSA 2048 does not match the expected result")
	}

	rsa4096 := getKey(gokey.RSA4096, t)
	if !pemEqual(rsa4096, knownRsa4096, t) {
		t.Fatal("generated RSA 4096 does not match the expected result")
	}
}

// reader which tries to negate the effect of crypto/internal/randutil/randutil.MaybeReadByte
// https://github.com/golang/go/blob/6269dcdc24d74379d8a609ce886149811020b2cc/src/crypto/internal/randutil/randutil.go#L25
// useful to regain deterministic output of some crypto routines with respect to a fixed pseudo-random stream
// such as https://github.com/golang/go/blob/6269dcdc24d74379d8a609ce886149811020b2cc/src/crypto/rsa/rsa.go#L222
type maybenotReader struct {
	io.Reader
	*testing.T
}

func (r maybenotReader) Read(p []byte) (n int, err error) {
	if strings.Contains(string(debug.Stack()), "crypto/internal/randutil.MaybeReadByte") {
		r.T.Log("mitigating crypto/internal/randutil.MaybeReadByte...")
		// feed MaybeReadByte with dummy zeroes
		for i := range p {
			p[i] = 0
		}
		return len(p), nil
	}

	return r.Reader.Read(p)
}

// generate an RSA key using stdlib crypto/rsa package
// but with crypto/internal/randutil.MaybeReadByte removed
func getStdRsaKey(bits int, t *testing.T) string {
	realm := "example.com"

	// mimic gokey.GetKey behaviour
	switch bits {
	case 2048:
		realm += "-key(RSA2048)"
	case 4096:
		realm += "-key(RSA4096)"
	}

	rng := &maybenotReader{gokey.NewDRNG("pass", realm), t}
	priv, err := rsa.GenerateKey(rng, bits)
	if err != nil {
		t.Fatal(err)
	}

	var b strings.Builder
	err = gokey.EncodeToPem(priv, &b)
	if err != nil {
		t.Fatal(err)
	}

	return b.String()
}

// this test is here to see if the RSA key generation algorithm in the stdlib is the same as in our package
// for now we want to track any stdlib changes to the algorithm
// however, this test might be removed in the future, if the changes in the stdlib will not be backwards compatible
func TestStdKey(t *testing.T) {
	rsa2048 := getStdRsaKey(2048, t)
	if !pemEqual(rsa2048, knownRsa2048, t) {
		t.Fatal("RSA key generation algorithm from stdlib deviates from the one in gokey")
	}

	rsa4096 := getStdRsaKey(4096, t)
	if !pemEqual(rsa4096, knownRsa4096, t) {
		t.Fatal("RSA key generation algorithm from stdlib deviates from the one in gokey")
	}
}
