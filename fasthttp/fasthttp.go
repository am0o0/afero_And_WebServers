package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("error in net.Listen: %v", err)
	}
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		filePath := ctx.QueryArgs().Peek("filePath")
		_ = ctx.Response.SendFile(string(filePath))
		ctx.SendFile(string(filePath))
		ctx.SendFileBytes(filePath)
		fileHeader, _ := ctx.FormFile("file")
		_ = fasthttp.SaveMultipartFile(fileHeader, string(filePath))
		fasthttp.ServeFile(ctx, string(filePath))
		fasthttp.ServeFileUncompressed(ctx, string(filePath))
		fasthttp.ServeFileBytes(ctx, filePath)
		fasthttp.ServeFileBytesUncompressed(ctx, filePath)
	}
	if err := fasthttp.Serve(ln, requestHandler); err != nil {
		log.Fatalf("error in Serve: %v", err)
	}
}
