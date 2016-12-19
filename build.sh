GOOS=linux   GOARCH=amd64 bash -c 'go build -o bin/cvlin-$GOOS-$GOARCH'
GOOS=darwin  GOARCH=amd64 bash -c 'go build -o bin/cvlin-$GOOS-$GOARCH'
GOOS=windows GOARCH=amd64 bash -c 'go build -o bin/cvlin-$GOOS-$GOARCH.exe'

# rm $GOPATH/src/github.com/megane42/cvlin/bin/*.zip
# ls $GOPATH/src/github.com/megane42/cvlin/bin/cvlin-* | xargs -I{} zip -m {}.zip {}
