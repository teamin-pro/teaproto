package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"os"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/teamin-pro/teanode/generated/teaproto"
)

//go:embed teaproto/version.txt
var apiVersion string

func main() {
	state()
	chats()
}

func chats() {
	mustSave("group-members.request.yml", &teaproto.ClientRequest{
		Actions: []*teaproto.Action{
			{
				Request: &teaproto.Action_GroupMembersListRequest{
					GroupMembersListRequest: &teaproto.GroupMembersListRequest{
						ChatId: sampleGroupId(),
						Limit:  2,
					},
				},
			},
			{
				Request: &teaproto.Action_GroupDetailsRequest{
					GroupDetailsRequest: &teaproto.GroupDetailsRequest{
						ChatId: sampleGroupId(),
					},
				},
			},
		},
	})

	mustSave("group-members.response.yml", &teaproto.ServerResponse{
		Results: []*teaproto.Result{
			{
				Response: &teaproto.Result_GroupMembersListResponse{
					GroupMembersListResponse: &teaproto.GroupMembersListResponse{
						Members: []*teaproto.GroupMember{
							{
								User:                 alice(),
								GroupRoleId:          1,
								CanIRemoveThisMember: true,
							},
							{
								User:                 bob(),
								GroupRoleId:          2,
								CanIRemoveThisMember: true,
							},
						},
					},
				},
			},
			{
				Response: &teaproto.Result_GroupDetailsResponse{
					GroupDetailsResponse: &teaproto.GroupDetailsResponse{
						GroupRoles: []*teaproto.GroupRole{
							{
								Id:                   1,
								Title:                "Admin",
								CanAddNewMembers:     true,
								CanChangeInformation: true,
								CanSetRoleIds:        []uint32{1, 2},
							},
							{
								Id:    2,
								Title: "User",
							},
						},
					},
				},
			},
		},
	})
}

func state() {
	mustSave("state.request.yml", &teaproto.ClientRequest{
		Actions: []*teaproto.Action{
			{
				Request: &teaproto.Action_StateRequest{
					StateRequest: &teaproto.StateRequest{},
				},
			},
		},
	})

	mustSave("state.response.yml", &teaproto.ServerResponse{
		Results: []*teaproto.Result{
			{
				Response: &teaproto.Result_StateResponse{
					StateResponse: &teaproto.StateResponse{
						Now:                      sampleDate(),
						ApiVersion:               strings.TrimSpace(apiVersion),
						MaxResultsOnPage:         10,
						MaxActionsInRequest:      10,
						MaxFileSize:              1024 * 1024 * 10,
						MinIconSize:              100,
						MaxUsernameLength:        20,
						MinPasswordLength:        5,
						MaxPasswordLength:        100,
						MaxMessageLength:         1024,
						MaxAudioLengthInSeconds:  120,
						MaxUploadsForMessage:     5,
						MaxGroupTitleLength:      10,
						PushNotificationsEnabled: false,
						AuthMethods: []teaproto.AuthMethod{
							teaproto.AuthMethod_EMAIL_CODE_AUTH_METHOD,
						},
					},
				},
			},
		},
	})

	mustSave("fake-error.request.yml", &teaproto.ClientRequest{
		Actions: []*teaproto.Action{
			{
				Request: &teaproto.Action_FakeErrorRequest{
					FakeErrorRequest: &teaproto.FakeErrorRequest{
						Mode: teaproto.FakeErrorRequest_RETURN_SERVER_ERROR,
					},
				},
			},
		},
	})

	mustSave("fake-error.response.yml", &teaproto.ServerResponse{
		Errors: []teaproto.ServerResponse_Error{
			teaproto.ServerResponse_INTERNAL_SERVER_ERROR,
		},
	})

	mustSave("badge.request.yml", &teaproto.ClientRequest{
		Token: "<token>",
		Actions: []*teaproto.Action{
			{
				Request: &teaproto.Action_BadgeRequest{
					BadgeRequest: &teaproto.BadgeRequest{},
				},
			},
		},
	})

	mustSave("badge.response.yml", &teaproto.ServerResponse{
		Results: []*teaproto.Result{
			{
				Response: &teaproto.Result_BadgeResponse{
					BadgeResponse: &teaproto.BadgeResponse{
						GlobalBadges: []*teaproto.GlobalBadge{
							{
								Type:      teaproto.BadgeType_BADGE_TYPE_MAIN,
								Counter:   42,
								UpdatedAt: sampleDate(),
							},
						},
					},
				},
			},
		},
	})
}

func mustSave(dest string, m proto.Message) {
	s, err := toJSON(m)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("codegen/samples/"+dest, []byte("```json\n"+s+"\n```\n"), 0644)
	if err != nil {
		panic(err)
	}
}

func toJSON(m proto.Message) (string, error) {
	rawJSON, err := protojson.Marshal(m)
	if err != nil {
		return "", err
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, rawJSON, "", "  ")
	if err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func sampleDate() uint64 {
	return unixTime(time.Date(2021, 1, 1, 2, 3, 4, 6, time.UTC))
}

func sampleGroupId() string {
	return "f47ac10b-58cc-4372-a567-0e02b2c3d479"
}

func alice() *teaproto.User {
	userId := "a2f1f0d6-5b9e-4e0e-8d4a-8b5b7a0e0b0e"
	return &teaproto.User{
		Id:   userId,
		Name: "Alice Smith",
		IsMe: true,
		Icon: &teaproto.Icon{
			Letters: "AS",
			Color:   "#ff0000",
		},
		OnlineStatus: &teaproto.OnlineStatus{
			UserId: userId,
			Status: teaproto.OnlineStatus_ONLINE,
		},
	}
}

func bob() *teaproto.User {
	userId := "b2f1f0d6-5b9e-4e0e-8d4a-8b5b7a0e0b0e"
	return &teaproto.User{
		Id:   userId,
		Name: "Bob Doe",
		IsMe: false,
		Icon: &teaproto.Icon{
			Letters: "BD",
			Color:   "#00ff00",
		},
		OnlineStatus: &teaproto.OnlineStatus{
			UserId:   userId,
			Status:   teaproto.OnlineStatus_OFFLINE,
			LastSeen: sampleDate(),
		},
	}
}

func unixTime(dt time.Time) uint64 {
	return uint64(dt.UTC().UnixMilli())
}
