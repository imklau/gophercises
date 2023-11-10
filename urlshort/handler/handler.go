package handler

import (
	"net/http"

	"gopkg.in/yaml.v2"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentPath := r.URL.Path

		if dest, ok := pathsToUrls[currentPath]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

type item struct {
	Path string `yaml:"path"`
	Url string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var items []item
	pathsToUrls := map[string]string{}

	err := yaml.Unmarshal(yml, &items)

	if err != nil {
		return nil, err
	}

	for _, item := range items {
		pathsToUrls[item.Path] = item.Url
	}


	return MapHandler(pathsToUrls, fallback), nil
}