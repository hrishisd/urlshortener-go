package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			fmt.Println(url)
			http.Redirect(w, r, url, http.StatusSeeOther)
			// No need to do the actual request
			// resp, err := http.Get(url)
			// if err != nil {
			// 	// panic(err)
			// 	fmt.Printf("http get on %s failed\n", url)
			// }
			// defer resp.Body.Close()
			// fmt.Println(resp.Body)
			// respBytes, err := ioutil.ReadAll(resp.Body)
			// if err != nil {
			// 	panic(err)
			// 	fmt.Println("read on resp body failed")
			// }
			// w.Write(respBytes)
			// w.WriteHeader(resp.StatusCode)
		} else {
			fallback.ServeHTTP(w, r)
		}
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
