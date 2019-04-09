// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package socket

import (
	"io"
	"net"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/zerjioang/etherniti/thirdparty/echo"
	"github.com/zerjioang/etherniti/shared/def/listener"
)

func TestUnixSocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		NewSocketListener()
	})
	t.Run("run", func(t *testing.T) {
		s := NewSocketListener()
		s.RunMode("/tmp/go.sock", true)
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		time.Sleep(200000 * time.Second)
	})
	t.Run("request-status", func(t *testing.T) {
		s := NewSocketListener()
		// run the socket server
		s.RunMode("/tmp/go.sock", true)
		err := s.Listen()
		if err != nil {
			t.Error(err)
		}
		// wait one second to bootup
		time.Sleep(1 * time.Second)
		// send GET style request for v1/public for welcome message
		cli := socketHttpClient("/tmp/go.sock")
		resp, err := cli.Get("http://unix/v1/public")
		t.Log("response", resp)
		t.Log("error", err)
	})
}

func TestUnixSocketListener_RunMode(t *testing.T) {
	type fields struct {
		e    *echo.Echo
		path string
		mode bool
	}
	type args struct {
		socketPath string
		background bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := UnixSocketListener{
				e:    tt.fields.e,
				path: tt.fields.path,
				mode: tt.fields.mode,
			}
			l.RunMode(tt.args.socketPath, tt.args.background)
		})
	}
}

func TestUnixSocketListener_Listen(t *testing.T) {
	type fields struct {
		e    *echo.Echo
		path string
		mode bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := UnixSocketListener{
				e:    tt.fields.e,
				path: tt.fields.path,
				mode: tt.fields.mode,
			}
			if err := l.Listen(); (err != nil) != tt.wantErr {
				t.Errorf("UnixSocketListener.Listen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewSocketListener(t *testing.T) {
	tests := []struct {
		name string
		want listener.ListenerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSocketListener(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSocketListener() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader(tt.args.r)
		})
	}
}

func Test_socketClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			socketClient()
		})
	}
}

func Test_socketHttpClient(t *testing.T) {
	type args struct {
		socketPath string
	}
	tests := []struct {
		name string
		args args
		want http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := socketHttpClient(tt.args.socketPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("socketHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnixSocketListener_unixServerListener(t *testing.T) {
	type fields struct {
		e    *echo.Echo
		path string
		mode bool
	}
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := UnixSocketListener{
				e:    tt.fields.e,
				path: tt.fields.path,
				mode: tt.fields.mode,
			}
			l.unixServerListener(tt.args.c)
		})
	}
}

func TestUnixSocketListener_background(t *testing.T) {
	type fields struct {
		e    *echo.Echo
		path string
		mode bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := UnixSocketListener{
				e:    tt.fields.e,
				path: tt.fields.path,
				mode: tt.fields.mode,
			}
			l.background()
		})
	}
}

func TestUnixSocketListener_foreground(t *testing.T) {
	type fields struct {
		e    *echo.Echo
		path string
		mode bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := UnixSocketListener{
				e:    tt.fields.e,
				path: tt.fields.path,
				mode: tt.fields.mode,
			}
			if err := l.foreground(); (err != nil) != tt.wantErr {
				t.Errorf("UnixSocketListener.foreground() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
