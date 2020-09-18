protoc example command:
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto

ubuntu be like:
protoc -I=/workspaces/Golang/proto/4-golang-proto/1-example/proto/ --go_out=/workspaces/Golang/proto/4-golang-proto/1-example/tutorial /workspaces/Golang/proto/4-golang-proto/1-example/proto/example.proto