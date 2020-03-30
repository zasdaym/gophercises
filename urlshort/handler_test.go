package main

import (
	"reflect"
	"testing"
)

func TestJsonToURLMaps(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []urlMap
		wantErr bool
	}{
		{
			name: "should convert valid json to urlMap",
			args: args{
				data: []byte(`[
					{
						"path": "/go",
						"url": "https://golang.org"
					}
				]`),
			},
			want: []urlMap{
				urlMap{
					Path: "/go",
					URL:  "https://golang.org",
				},
			},
			wantErr: false,
		},
		{
			name: "should return err on invalid json",
			args: args{
				data: []byte(`{"status": false}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jsonToURLMaps(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonToURLMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonToURLMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYamlToURLMaps(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []urlMap
		wantErr bool
	}{
		{
			name: "should convert valid yaml to urlMap",
			args: args{
				data: []byte(`
- path: /go
  url: https://golang.org`),
			},
			want: []urlMap{
				urlMap{
					Path: "/go",
					URL:  "https://golang.org",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error on invalid yaml",
			args: args{
				data: []byte(`{notvalidyaml}`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlToURLMaps(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlToURLMaps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlToURLMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildURLMap(t *testing.T) {
	type args struct {
		urlMaps []urlMap
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "convert to map",
			args: args{
				urlMaps: []urlMap{
					urlMap{
						Path: "/go",
						URL:  "https://golang.org",
					},
				},
			},
			want: map[string]string{
				"/go": "https://golang.org",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildURLMap(tt.args.urlMaps); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildURLMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
