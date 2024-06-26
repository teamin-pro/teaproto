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
PROTOBUF_PACKAGE?=teaproto

codegen: go-protobuf
	cd codegen
	go run codegen/main.go

version:
	cat docs/CHANGELOG.md  | grep '##' | head -n 1 | sed -e 's/## //' > version.txt

go-protobuf: version
	mkdir -p $(PROTOBUF_DIR)/$(PROTOBUF_PACKAGE)
	cp version.txt $(PROTOBUF_DIR)/$(PROTOBUF_PACKAGE)/
	rm -rf $(PROTOBUF_DIR)/$(PROTOBUF_PACKAGE)/*.pb.go
	protoc --version
	protoc --fatal_warnings -I protobuf \
		--go_opt=Mactions.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mauth.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mbadges.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mboard_members.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mboards.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mchats.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mdevices.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mgroup_members.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mgroups.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mmessages.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mreactions.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mstate.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mtask_members.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mtasks.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tbadges.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tboards.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tchats.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tcommands.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tgroups.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/timages.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/ttasks.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tuploads.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mteatypes/tusers.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Mtokens.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Muploads.proto=./$(PROTOBUF_PACKAGE) \
		--go_opt=Musers.proto=./$(PROTOBUF_PACKAGE) \
		--go_out=$(PROTOBUF_DIR) \
		protobuf/actions.proto \
		protobuf/auth.proto \
		protobuf/badges.proto \
		protobuf/board_members.proto \
		protobuf/boards.proto \
		protobuf/chats.proto \
		protobuf/devices.proto \
		protobuf/group_members.proto \
		protobuf/groups.proto \
		protobuf/messages.proto \
		protobuf/reactions.proto \
		protobuf/state.proto \
		protobuf/task_members.proto \
		protobuf/tasks.proto \
		protobuf/teatypes/tbadges.proto \
		protobuf/teatypes/tboards.proto \
		protobuf/teatypes/tchats.proto \
		protobuf/teatypes/tcommands.proto \
		protobuf/teatypes/tgroups.proto \
		protobuf/teatypes/timages.proto \
		protobuf/teatypes/ttasks.proto \
		protobuf/teatypes/tuploads.proto \
		protobuf/teatypes/tusers.proto \
		protobuf/tokens.proto \
		protobuf/uploads.proto \
		protobuf/users.proto
