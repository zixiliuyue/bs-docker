

package service

import (
	"reflect"
	"testing"
)

func TestLvsProxy_IsVSAvailable(t *testing.T) {
	lvs := NewLvsProxy("sh")

	vsList := []struct {
		vs string
		ok bool
	}{
		{
			vs: "127.0.0.1:6443",
			ok: true,
		},
		{
			vs: "127.0.0.2:6443",
			ok: false,
		},
	}

	for _, server := range vsList {
		ok := lvs.IsVirtualServerAvailable(server.vs)
		if server.ok != ok {
			t.Logf("IsVirtualServerAvailable failed")
		}
	}

	t.Logf("IsVirtualServerAvailable successful")
}

func TestLvsProxy_CreateVirtualServer(t *testing.T) {
	lvs := NewLvsProxy("sh")

	vsList := []string{"127.0.0.1:6443", "127.0.0.2:6443"}

	for _, server := range vsList {
		err := lvs.CreateVirtualServer(server)
		if err != nil {
			t.Logf("CreateVirtualServer failed")
		}
	}

	t.Logf("CreateVirtualServer successful")
}

func TestLvsProxy_DeleteVirtualServer(t *testing.T) {
	lvs := NewLvsProxy("sh")
	vsList := []string{"127.0.0.1:6443", "127.0.0.2:6443"}

	for _, server := range vsList {
		err := lvs.DeleteVirtualServer(server)
		if err != nil {
			t.Logf("DeleteVirtualServer failed")
		}
	}

	t.Logf("DeleteVirtualServer successful")
}

func TestLvsProxy_CreateRealServer(t *testing.T) {
	lvs := NewLvsProxy("sh")

	vs := "127.0.0.1:6443"
	ok := lvs.IsVirtualServerAvailable(vs)
	if !ok {
		err := lvs.CreateVirtualServer(vs)
		if err != nil {
			t.Fatalf("CreateVirtualServer failed: %v", err)
			return
		}
	}

	rss := []string{"192.168.0.1:8081", "192.168.0.2:8082", "192.168.0.3:8083"}
	for _, rs := range rss {
		err := lvs.CreateRealServer(rs)
		if err != nil {
			t.Fatalf("CreateRealServer rs[%s] failed: %v", rs, err)
			return
		}
	}

	t.Logf("CreateRealServer successful")
}

func TestLvsProxy_DeleteRealServer(t *testing.T) {
	lvs := NewLvsProxy("sh")

	vs := "127.0.0.1:6443"
	ok := lvs.IsVirtualServerAvailable(vs)
	if !ok {
		t.Fatalf("IsVirtualServerAvailable failed: %s", vs)
		return
	}

	rss := []string{"192.168.0.1:8081", "192.168.0.2:8082", "192.168.0.3:8083"}
	for _, rs := range rss {
		err := lvs.DeleteRealServer(rs)
		if err != nil {
			t.Fatalf("CreateRealServer rs[%s] failed: %v", rs, err)
			return
		}
	}

	t.Logf("DeleteRealServer successful")
}

func TestLvsProxy_ListRealServer(t *testing.T) {
	lvs := NewLvsProxy("sh")

	vs := "127.0.0.1:6443"
	ok := lvs.IsVirtualServerAvailable(vs)
	if !ok {
		t.Fatalf("IsVirtualServerAvailable failed: %s", vs)
		return
	}

	expectRss := []string{"192.168.0.1:8081", "192.168.0.2:8082", "192.168.0.3:8083"}
	rss, err := lvs.ListRealServer()
	if err != nil {
		t.Fatalf("ListRealServer vs[%s] failed", vs)
		return
	}

	if !reflect.DeepEqual(rss, expectRss) {
		t.Logf("ListRealServer failed")
	}

	t.Logf("ListRealServer successful")
}
