// Copyright (c) 2015 The SkyDNS Authors. All rights reserved.
// Use of this source code is governed by The MIT License (MIT) that can be
// found in the LICENSE file.

package msg

import "testing"

func TestSplit255(t *testing.T) {
	xs := split255("abc")
	if len(xs) != 1 && xs[0] != "abc" {
		t.Logf("Failure to split abc")
		t.Fail()
	}
	s := ""
	for i := 0; i < 255; i++ {
		s += "a"
	}
	xs = split255(s)
	if len(xs) != 1 && xs[0] != s {
		t.Logf("Failure to split 255 char long string")
		t.Logf("%s %v\n", s, xs)
		t.Fail()
	}
	s += "b"
	xs = split255(s)
	if len(xs) != 2 || xs[1] != "b" {
		t.Logf("Failure to split 256 char long string: %d", len(xs))
		t.Logf("%s %v\n", s, xs)
		t.Fail()
	}
	for i := 0; i < 255; i++ {
		s += "a"
	}
	xs = split255(s)
	if len(xs) != 3 || xs[2] != "a" {
		t.Logf("Failure to split 510 char long string: %d", len(xs))
		t.Logf("%s %v\n", s, xs)
		t.Fail()
	}
}

func TestDomainPath(t *testing.T) {
	tests := []string{
		"test.skydns.local.", "/skydns/local/skydns/test",
		"_ldap._tcp.skydns.local.", "/skydns/local/skydns/" + underScore + "tcp/" + underScore + "ldap",
		"l_dap.skydns.local.", "/skydns/local/skydns/l_dap",
	}

	for i := 0; i < len(tests)-1; i += 2 {
		path := Path(tests[i])
		if path != tests[i+1] {
			t.Logf("Path %s: expected %s got %s\n", tests[i], tests[i+1], path)
			t.Fail()
		}
		roundtrip := Domain(path)
		if roundtrip != tests[i] {
			t.Logf("Domain %s: expect %s got %s\n", path, tests[i], roundtrip)
			t.Fail()
		}
	}
}
