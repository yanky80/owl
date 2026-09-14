package sip

import "testing"

func TestNewResponseFromRequestIncludesGBVersion(t *testing.T) {
	uri, err := ParseSipURI("sip:34020000001320000001@127.0.0.1")
	if err != nil {
		t.Fatal(err)
	}

	req := NewRequest("", MethodRegister, &uri, DefaultSipVersion, nil, nil)
	res := NewResponseFromRequest("", req, 200, "OK", nil)

	headers := res.GetHeaders("X-GB-Ver")
	if len(headers) != 1 || headers[0].String() != "X-GB-Ver: 3.0" {
		t.Fatalf("X-GB-Ver = %v, want X-GB-Ver: 3.0", headers)
	}
}
