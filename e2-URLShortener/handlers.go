package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data, err := parseJson(jsonData)
	jsonPathMapping := buildJSONPathMap(data)
	return MapHandler(jsonPathMapping, fallback), err
}

func parseJson(data []byte) ([]map[string]string, error) {
	var jsonData []map[string]string
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		fmt.Println("error:", err)
	}
	return jsonData, err
}

func buildJSONPathMap(jsonData []map[string]string) map[string]string {
	mapping := make(map[string]string)
	for idx := range jsonData {
		entry := jsonData[idx]
		mapping[entry["path"]] = entry["url"]
	}
	return mapping
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
		if _, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, pathsToUrls[path], http.StatusPermanentRedirect)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYaml(yml)
	ymlPathMapping := buildYAMLPathMap(parsedYaml)
	return MapHandler(ymlPathMapping, fallback), err
}

type YAMLPathMapping struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(ymlBytes []byte) ([]YAMLPathMapping, error) {
	var pathURLs []YAMLPathMapping
	err := yaml.Unmarshal(ymlBytes, &pathURLs)
	return pathURLs, err
}

func buildYAMLPathMap(paths []YAMLPathMapping) map[string]string {
	mapping := map[string]string{}
	for _, ymlEntry := range paths {
		mapping[ymlEntry.Path] = ymlEntry.URL
	}
	return mapping
}
