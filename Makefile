.PHONY: all test clean

install:
	pip install --upgrade pip
	pip install mkdocs mkdocs-include-markdown-plugin

docs: install codegen
	mkdocs build

serve: install codegen
	mkdocs serve

gh-deploy: install
	mkdocs gh-deploy

PROTOBUF_DIR?=codegen
PROTOBUF_PACKAGE?=gtproto

codegen: go-protobuf
	cd codegen
	go run codegen/main.go

version:
	cat docs/CHANGELOG.md  | grep '##' | head -n 1 | sed -e 's/## //' > $(PROTOBUF_DIR)/gtproto/version.txt

go-protobuf: version
	mkdir -p $(PROTOBUF_DIR)
	rm -rf $(PROTOBUF_DIR)/$(PROTOBUF_PACKAGE)/*.pb.go
	protoc --fatal_warnings -I protobuf \
		--go_opt=Mactions.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mauth.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mbadges.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mchat_events.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mcommands.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mcommons.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mdevices.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mgroup_chats.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mgroups.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mmessages.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mreactions.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mstate.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Msystem_messages.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mtokens.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Muploads.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Musers.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mviewed_messages.proto=./$(PROTOBUF_PACKAGE) \
		--go_out=$(PROTOBUF_DIR) \
		protobuf/*.proto
