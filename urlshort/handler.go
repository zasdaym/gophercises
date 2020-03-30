package main

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type urlMap struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if destination, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destination, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlMaps, err := yamlToURLMaps(yamlData)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildURLMap(urlMaps)
	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     {
//       "path": "/some-path",
//       "url": "https://www.some-url.com/demo"
//     }
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	urlMaps, err := jsonToURLMaps(jsonData)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildURLMap(urlMaps)
	return MapHandler(pathsToUrls, fallback), nil
}

func buildURLMap(urlMaps []urlMap) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, item := range urlMaps {
		pathsToUrls[item.Path] = item.URL
	}
	return pathsToUrls
}

func yamlToURLMaps(data []byte) ([]urlMap, error) {
	var urlMaps []urlMap
	err := yaml.Unmarshal(data, &urlMaps)
	if err != nil {
		return nil, err
	}
	return urlMaps, nil
}

func jsonToURLMaps(data []byte) ([]urlMap, error) {
	var urlMaps []urlMap
	err := json.Unmarshal(data, &urlMaps)
	if err != nil {
		return nil, err
	}
	return urlMaps, nil
}
