// Copyright gaethway
// SPDX-License-Identifier: Apache License 2.0

// +build !prod
// +build !dev

package config

const (
	DevelopmentAddress = "localhost"
	HttpsPort          = ":4430"
	HttpPort           = ":8080"
	HttpAddress        = DevelopmentAddress + HttpPort
	HttpsAddress       = DevelopmentAddress + HttpsPort
)

const (
	CertPem = `-----BEGIN CERTIFICATE-----
MIIC+jCCAeKgAwIBAgIRAI4ga6WaCWzhnIgevZi02qgwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAeFw0xOTAxMjgxNjA0NDNaFw0yMDAxMjgxNjA0
NDNaMBIxEDAOBgNVBAoTB0FjbWUgQ28wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQCTSjnQuVl/vC++FnPFcREbwIsrPLY8G4CQHgKqxYkR0J/NCyCSwNfC
/dGDRSE0Lun9lvUjjXBi94Ftd+r+f3okgYPrmgn7K/R5N/K+3kGvgfxUgZFXEYtK
z5wojb+pUFwTdgfT3BHp2naBFLMKI838A3Jt5MHEXJWENHs8ovchMWivlVoBjEJD
B+SaJUGD7+PC1vvGKda/P52X+sYKPrwnlze0sNdtYD1OUX4W+YntJZdr9CgznPMg
QYSZsqRr4oGiS7ONJCfxFnGvHL/WwyBfin+QXLUTbqkSa4aXPYVD5om8tk7eGL3F
eKmwaGkC5xybq7oEUCa/DlkpJgxyDTCjAgMBAAGjSzBJMA4GA1UdDwEB/wQEAwIF
oDATBgNVHSUEDDAKBggrBgEFBQcDATAMBgNVHRMBAf8EAjAAMBQGA1UdEQQNMAuC
CWxvY2FsaG9zdDANBgkqhkiG9w0BAQsFAAOCAQEAeCNiTCXCwKNkXvXZaP+xfYcs
fSB3S/UAnUxfmOBCdfyK4JCM47BA3Hz9SzLMMwrR4IP53a9hXfQxYiMffZi8R3XF
YWJTnS3giuOhe8aFH91PhPDF+sC5dlm1cd2B3i1ylv0ogbrO9ZGtO47zA41bTPiy
E8IccKiKru2bL0llj4aqg0sdHmdLMBtsjWbT/yQaveBG/bNNDk0u5IqgJWSVePwk
jFPtgDvxFkDoDAhzrJcenMSt6LtTAoBLKkWPSRC3u+iwVLacIv0pmxj+1nGW+H18
mklI/9mByeejncVBGPp5vHastJpTFyRJ4V8CRZOQ4j9fRx7sEmQ7N+9pqDNtsw==
-----END CERTIFICATE-----`
	KeyPem = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAk0o50LlZf7wvvhZzxXERG8CLKzy2PBuAkB4CqsWJEdCfzQsg
ksDXwv3Rg0UhNC7p/Zb1I41wYveBbXfq/n96JIGD65oJ+yv0eTfyvt5Br4H8VIGR
VxGLSs+cKI2/qVBcE3YH09wR6dp2gRSzCiPN/ANybeTBxFyVhDR7PKL3ITFor5Va
AYxCQwfkmiVBg+/jwtb7xinWvz+dl/rGCj68J5c3tLDXbWA9TlF+FvmJ7SWXa/Qo
M5zzIEGEmbKka+KBokuzjSQn8RZxrxy/1sMgX4p/kFy1E26pEmuGlz2FQ+aJvLZO
3hi9xXipsGhpAuccm6u6BFAmvw5ZKSYMcg0wowIDAQABAoIBAHHvkRmsx1bQM/5P
T+8Dr8BQCVfA9xc4DxNso5OGiqmFQJhUazYahs0HmvJ4n17Gi6rnA2olFzL3Ut9j
TBzib5GdvnaaCe6J6et7JAQR2a3yV0bnk45Ou/l678lPHVvUFeXX/+Ya7qB/pfvk
Dztgxw6TfAkWU+2Z0O8bydj2F0VMzskrxwwF384mZgB5ysjiVcMmKI2kCJa/ovUd
B6kFMd6Y77ohlyI0jf9YWIOKeMnRcsKorFfQtzqBpTpa3purCieUMY4jh/5iYa0n
UTxmQOAoAYerNVu/d5Qayy9Y1VAW+Zl8pBOknvK+Zo49O0Dx1/vRoIEUSolczexY
CdqUVkECgYEAw1JHLGq3Cf0YjEJeSMi1IMPpx2PdGGmtznKnnxXCsYHkue/mXl8K
l58f0QmDtlwkfNRRfSLj4XsZLQk1UvzJ1aCbehwfGqEAdtCOb2A1GWlKipvIR4pr
b5NpMToJ+3jcP1cJJTh18bOUG75y5axKz2MkiG8I62dFcG0b/puLRh0CgYEAwQwN
gBUM/VpinnqZ3xC5vQeXn6k3BLu773CgSaBFcyMJedWrpKA7d+87kvn5N9onN1Ww
aJF8MckREwefKp6D0UfJYusJD5DAInYfeKPX2PT+OKncT3G4G8kq0MSjsvQks8B/
zvPsTszJKIYJqfe7KShSDIZY6GDBpCIsw8aZlb8CgYBNWrGLWrwg/ZaSPdqfUrXB
QzW73MX8XCYUg/30mCaiLEJMjUEcEOHeCIwOOolqWHWu5ltbhszfSORAnMv8kNbS
fyf0JV0AK9FGPPScEWsWJEf8OxQHmT9RUf0wHL9FU6lOgIbDseesEKXQkw1n/mMm
XSpjyi2rJRwwGVYj8LAo1QKBgQCbbbbY7xn8Sm+opZGJ9g910M0VccqodvbDu+xy
GyaPoyAYBh8idxgqYmWW2sj7XRvCA637I1fZRcgHiFVwnRwIvkG48P/klmj71htU
qKY7OlYNDUYieK8BQCDG4evjQ4rhZxYAbIhQkbVMeU8CmEEKzDnzd5/RyUVff1yH
bDlwRQKBgGxz/v8dIP5xRXXXQYre+KxXtohY7QsDJxuC3R2NMCh9lovrsqWof5OQ
xC1++6t6BnPJnMe4vdpMeuW8QTAKhHvm+XvPiPqnNeVSj7SLbOZDlivUiNZrr87t
DagBWzI58Ymmo2EJHbe48ChjOf5aeZpH7l8ZtSDbdHRFOKcUPDUJ
-----END RSA PRIVATE KEY-----`

	TokenSecret = "t0k3n-s3cr3t-h3r3"
)

//simply converts http requests into https
func GetRedirectUrl(host string, path string) string {
	return "https://" + host + path
}