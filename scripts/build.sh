echo "Update dependency..."
GOPROXY=direct go get -u github.com/ngdangkietswe/swe-protobuf-shared
GOPROXY=direct go get -u github.com/ngdangkietswe/swe-go-common-shared
go mod tidy
go mod vendor
echo "Update dependency successful!"