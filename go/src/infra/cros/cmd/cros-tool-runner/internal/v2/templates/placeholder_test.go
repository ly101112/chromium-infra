// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"errors"
	"testing"

	"go.chromium.org/chromiumos/config/go/test/api"
	labApi "go.chromium.org/chromiumos/config/go/test/lab/api"
)

type mockLookuper struct {
	ContainerLookuper
	portLookupFunc func(string) ([]*api.Container_PortBinding, error)
	ipLookupFunc   func(string) (string, error)
	hostLookupFunc func() (string, error)
}

func (m *mockLookuper) LookupContainerPortBindings(name string) ([]*api.Container_PortBinding, error) {
	return m.portLookupFunc(name)
}

func (m *mockLookuper) LookupContainerIpAddress(name string) (string, error) {
	return m.ipLookupFunc(name)
}

func (m *mockLookuper) LookupHostIpAddress() (string, error) {
	return m.hostLookupFunc()
}

func TestExtract_invalid(t *testing.T) {
	router := populatorRouter{containerLookuper: &mockLookuper{}}
	endpoint := labApi.IpEndpoint{Address: "test"}
	scheme, _, err := router.extract(endpoint)
	if scheme != "" {
		t.Fatalf("scheme should be empty for invalid endpoint")
	}
	if err == nil {
		t.Fatalf("error should be thrown")
	}
}

func TestExtract_containerPort(t *testing.T) {
	router := populatorRouter{containerLookuper: &mockLookuper{}}
	endpoint := labApi.IpEndpoint{Address: "ctr-container-port://container-name", Port: 0}
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 0}
	scheme, returnedEndpoint, err := router.extract(endpoint)
	if scheme != "ctr-container-port" {
		t.Fatalf("scheme does not match")
	}
	if err != nil {
		t.Fatalf("unexpectedError")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerPort(t *testing.T) {
	expectedAddress := "container-name"
	expectedPort := 4222
	expectedContainerName := "container-name"
	expectedEndpoint := labApi.IpEndpoint{
		Address: expectedAddress,
		Port:    int32(expectedPort),
	}
	router := populatorRouter{containerLookuper: &mockLookuper{
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			if s != expectedContainerName {
				t.Fatalf("container name does not match\nexpect: %s\nactual: %s",
					expectedContainerName, s)
			}
			return []*api.Container_PortBinding{{ContainerPort: int32(expectedPort)}}, nil
		}}}
	endpoint := labApi.IpEndpoint{Address: "ctr-container-port://container-name", Port: 0}

	returnedEndpoint, err := router.populate(endpoint)

	if err != nil {
		t.Fatalf("unexpectedError %v", err)
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerPort_error(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 0}
	router := populatorRouter{containerLookuper: &mockLookuper{
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			return nil, errors.New("command throw error")
		}}}
	endpoint := labApi.IpEndpoint{Address: "ctr-container-port://container-name", Port: 0}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerPort_multiplePorts(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 0}
	router := populatorRouter{containerLookuper: &mockLookuper{
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			return []*api.Container_PortBinding{{ContainerPort: 42}, {ContainerPort: 43}}, nil
		}}}
	endpoint := labApi.IpEndpoint{Address: "ctr-container-port://container-name", Port: 0}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerPort_nonZeroPortInput(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 1}
	router := populatorRouter{containerLookuper: &mockLookuper{
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			return []*api.Container_PortBinding{{ContainerPort: 42}}, nil
		}}}
	endpoint := labApi.IpEndpoint{Address: "ctr-container-port://container-name", Port: 1}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerIp(t *testing.T) {
	expectedAddress := "192.168.10.2"
	expectedPort := int32(80)
	expectedEndpoint := labApi.IpEndpoint{
		Address: expectedAddress,
		Port:    expectedPort,
	}
	router := populatorRouter{containerLookuper: &mockLookuper{
		ipLookupFunc: func(s string) (string, error) {
			return expectedAddress, nil
		},
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			return []*api.Container_PortBinding{{ContainerPort: expectedPort}}, nil
		},
	}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-container-ip://container-name",
		Port:    0,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err != nil {
		t.Fatalf("unexpectedError: %v", err)
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerIp_nonZeroPort(t *testing.T) {
	expectedAddress := "192.168.10.2"
	expectedPort := int32(123)
	expectedEndpoint := labApi.IpEndpoint{
		Address: expectedAddress,
		Port:    expectedPort,
	}
	router := populatorRouter{containerLookuper: &mockLookuper{
		ipLookupFunc: func(s string) (string, error) {
			return expectedAddress, nil
		},
	}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-container-ip://container-name",
		Port:    expectedPort,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err != nil {
		t.Fatalf("unexpectedError: %v", err)
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerIp_error(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 1}
	router := populatorRouter{containerLookuper: &mockLookuper{
		ipLookupFunc: func(s string) (string, error) {
			return "", errors.New("command throw error")
		}}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-container-ip://container-name",
		Port:    1,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_containerIp_empty(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 1}
	router := populatorRouter{containerLookuper: &mockLookuper{
		ipLookupFunc: func(s string) (string, error) {
			return "", nil
		}}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-container-ip://container-name",
		Port:    1,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_hostIp(t *testing.T) {
	expectedAddress := "192.168.10.2"
	expectedPort := int32(40080)
	expectedEndpoint := labApi.IpEndpoint{
		Address: expectedAddress,
		Port:    expectedPort,
	}
	router := populatorRouter{containerLookuper: &mockLookuper{
		hostLookupFunc: func() (string, error) {
			return expectedAddress, nil
		},
		portLookupFunc: func(s string) ([]*api.Container_PortBinding, error) {
			return []*api.Container_PortBinding{{HostPort: expectedPort}}, nil
		},
	}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-host-ip://container-name",
		Port:    0,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err != nil {
		t.Fatalf("unexpectedError: %v", err)
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_hostIp_nonZeroPort(t *testing.T) {
	expectedAddress := "192.168.10.2"
	expectedPort := int32(8082)
	expectedEndpoint := labApi.IpEndpoint{
		Address: expectedAddress,
		Port:    expectedPort,
	}
	router := populatorRouter{containerLookuper: &mockLookuper{
		hostLookupFunc: func() (string, error) {
			return expectedAddress, nil
		},
	}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-host-ip://container-name",
		Port:    expectedPort,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err != nil {
		t.Fatalf("unexpectedError: %v", err)
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func TestPopulate_hostIp_error(t *testing.T) {
	expectedEndpoint := labApi.IpEndpoint{Address: "container-name", Port: 1}
	router := populatorRouter{containerLookuper: &mockLookuper{
		hostLookupFunc: func() (string, error) {
			return "", errors.New("command throw error")
		}}}
	endpoint := labApi.IpEndpoint{
		Address: "ctr-host-ip://container-name",
		Port:    1,
	}

	returnedEndpoint, err := router.populate(endpoint)

	if err == nil {
		t.Fatalf("expect error to be returned")
	}
	checkEndpoint(t, expectedEndpoint, returnedEndpoint)
}

func checkEndpoint(t *testing.T, expect labApi.IpEndpoint, actual labApi.IpEndpoint) {
	if actual.Address != expect.Address || actual.Port != expect.Port {
		t.Fatalf("returned endpoint doesn't match\nexpect: %v\nactual: %v",
			expect, actual)
	}
}
