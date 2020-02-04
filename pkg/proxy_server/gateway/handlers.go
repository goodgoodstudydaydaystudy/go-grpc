package gateway

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	log "github.com/sirupsen/logrus"

	"goodgoodstudy.com/go-grpc/pkg/proxy_server/ui/data/swagger"
)

// swaggerServer returns swagger specification files located under "/swagger/"
func swaggerServer(dir string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
			log.Errorf("Not Found: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}

		log.Infof("Serving %s", r.URL.Path)
		p := strings.TrimPrefix(r.URL.Path, "/v1/swagger/") // nginx friendly
		p = path.Join(dir, p)
		http.ServeFile(w, r, p)
	}
}

// swaggerUI Server
func serveSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/v1/swagger-ui/" // nginx friendly
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
