package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
	"net/http"
	"regexp"
)

func DownloadFile(c *gin.Context) {
	// https://github.com/advisories/GHSA-85cf-gj29-f555
	// https://github.com/1Panel-dev/1Panel/commit/f6b84d384e41c1c708d41160b2af9761d984f558#diff-3daefeffef5edafe17bc47b7e75c1332b684aa728ee67efffc6fe795d95d3627L556-L568
	c.File("req.Path")
}
func test() {

	//osFS := afero.NewMemMapFs()
	// OR
	osFS := afero.NewOsFs()
	fmt.Println(osFS.MkdirAll("tmp/b", 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/a", []byte("this is me a !"), 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/b/c", []byte("this is me c !"), 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/d", []byte("this is me d !"), 0755))
	// HttpFS
	fmt.Println("HttpFS:")
	httpFs := afero.NewHttpFs(osFS)
	httpFile, _ := httpFs.Open("tmp/b/dd/../../d")
	tmpbytes := make([]byte, 30)
	fmt.Println(httpFile.Read(tmpbytes))
	fmt.Println(string(tmpbytes))

}
func main() {
	//test()
	http.HandleFunc("/afero", hello)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		return
	}
}

func hello(writer http.ResponseWriter, request *http.Request) {
	// https://github.com/advisories/GHSA-hf7j-xj3w-87g4
	// https://github.com/spf13/afero
	filepath := request.URL.Query()["filepath"][0]
	//osFS := afero.NewMemMapFs()
	// OR
	osFS := afero.NewOsFs()
	fmt.Println(osFS.MkdirAll("tmp/b", 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/a", []byte("this is me a !"), 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/b/c", []byte("this is me c !"), 0755))
	fmt.Println(afero.WriteFile(osFS, "tmp/d", []byte("this is me d !"), 0755))
	content, _ := afero.ReadFile(osFS, filepath)
	fmt.Println(string(content))
	fmt.Println(osFS.Open(filepath))
	//fmt.Println(afero.SafeWriteReader(osFS, filepath, os.Stdout))
	//fmt.Println(afero.WriteReader(osFS, filepath, os.Stdout))

	// BasePathFs
	fmt.Println("BasePathFs:")
	basePathFs := afero.NewBasePathFs(osFS, "tmp")
	fmt.Println(afero.ReadFile(basePathFs, filepath))

	// RegexpFs
	fmt.Println("RegexpFs:")
	regex, _ := regexp.Compile(".*")
	regexpFs := afero.NewRegexpFs(osFS, regex)
	fmt.Println(afero.ReadFile(regexpFs, filepath))

	// ReadOnlyFS
	fmt.Println("ReadOnlyFS:")
	readOnlyFS := afero.NewReadOnlyFs(osFS)
	fmt.Println(afero.ReadFile(readOnlyFS, filepath))

	// CacheOnReadFs
	fmt.Println("CacheOnReadFs:")
	cacheOnReadFs := afero.NewCacheOnReadFs(osFS, osFS, 10)
	fmt.Println(afero.ReadFile(cacheOnReadFs, filepath))

	// HttpFS
	fmt.Println("HttpFS:")
	httpFs := afero.NewHttpFs(osFS)
	httpFile, _ := httpFs.Open(filepath)
	tmpbytes := make([]byte, 30)
	fmt.Println(httpFile.Read(tmpbytes))
	fmt.Println(string(tmpbytes))

	// Afero
	fmt.Println("Afero:")
	afs := &afero.Afero{Fs: osFS}
	fmt.Println(afs.ReadFile(filepath))

	// IOFS ==>safe
	fmt.Println("IOFS:")
	ioFS := afero.NewIOFS(osFS)
	fmt.Println(ioFS.ReadFile(filepath))
	fmt.Println(ioFS.Open(filepath))
}
