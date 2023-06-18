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

	"github.com/gamarjoba-team/gtnode/gtproto/codegen/gtproto"
)

//go:embed gtproto/version.txt
var gtprotoVersion string

func main() {
	state()
	chats()
}

func chats() {
	mustSave("group-members.request.yml", &gtproto.ClientRequest{
		Actions: []*gtproto.Action{
			{
				Request: &gtproto.Action_GroupMembersListRequest{
					GroupMembersListRequest: &gtproto.GroupMembersListRequest{
						ChatId: sampleGroupId(),
						Limit:  2,
					},
				},
			},
			{
				Request: &gtproto.Action_GroupDetailsRequest{
					GroupDetailsRequest: &gtproto.GroupDetailsRequest{
						ChatId: sampleGroupId(),
					},
				},
			},
		},
	})

	mustSave("group-members.response.yml", &gtproto.ServerResponse{
		Results: []*gtproto.Result{
			{
				Response: &gtproto.Result_GroupMembersListResponse{
					GroupMembersListResponse: &gtproto.GroupMembersListResponse{
						Members: []*gtproto.GroupMember{
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
				Response: &gtproto.Result_GroupDetailsResponse{
					GroupDetailsResponse: &gtproto.GroupDetailsResponse{
						GroupRoles: []*gtproto.GroupRole{
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
	mustSave("state.request.yml", &gtproto.ClientRequest{
		Actions: []*gtproto.Action{
			{
				Request: &gtproto.Action_StateRequest{
					StateRequest: &gtproto.StateRequest{},
				},
			},
		},
	})

	mustSave("state.response.yml", &gtproto.ServerResponse{
		Results: []*gtproto.Result{
			{
				Response: &gtproto.Result_StateResponse{
					StateResponse: &gtproto.StateResponse{
						Now:                      sampleDate(),
						GtprotoVersion:           strings.TrimSpace(gtprotoVersion),
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
						AuthMethods: []gtproto.AuthMethod{
							gtproto.AuthMethod_EMAIL_CODE_AUTH_METHOD,
						},
					},
				},
			},
		},
	})

	mustSave("fake-error.request.yml", &gtproto.ClientRequest{
		Actions: []*gtproto.Action{
			{
				Request: &gtproto.Action_FakeErrorRequest{
					FakeErrorRequest: &gtproto.FakeErrorRequest{
						Mode: gtproto.FakeErrorRequest_RETURN_SERVER_ERROR,
					},
				},
			},
		},
	})

	mustSave("fake-error.response.yml", &gtproto.ServerResponse{
		Errors: []gtproto.ServerResponse_Error{
			gtproto.ServerResponse_INTERNAL_SERVER_ERROR,
		},
	})

	mustSave("badge.request.yml", &gtproto.ClientRequest{
		Token: "<token>",
		Actions: []*gtproto.Action{
			{
				Request: &gtproto.Action_BadgeRequest{
					BadgeRequest: &gtproto.BadgeRequest{},
				},
			},
		},
	})

	mustSave("badge.response.yml", &gtproto.ServerResponse{
		Results: []*gtproto.Result{
			{
				Response: &gtproto.Result_BadgeResponse{
					BadgeResponse: &gtproto.BadgeResponse{
						GlobalBadges: []*gtproto.GlobalBadge{
							{
								Type:      gtproto.BadgeType_BADGE_TYPE_MAIN,
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

func alice() *gtproto.User {
	userId := "a2f1f0d6-5b9e-4e0e-8d4a-8b5b7a0e0b0e"
	return &gtproto.User{
		Id:   userId,
		Name: "Alice Smith",
		IsMe: true,
		Icon: &gtproto.Icon{
			Letters: "AS",
			Color:   "#ff0000",
		},
		OnlineStatus: &gtproto.OnlineStatus{
			UserId: userId,
			Status: gtproto.OnlineStatus_ONLINE,
		},
	}
}

func bob() *gtproto.User {
	userId := "b2f1f0d6-5b9e-4e0e-8d4a-8b5b7a0e0b0e"
	return &gtproto.User{
		Id:   userId,
		Name: "Bob Doe",
		IsMe: false,
		Icon: &gtproto.Icon{
			Letters: "BD",
			Color:   "#00ff00",
		},
		OnlineStatus: &gtproto.OnlineStatus{
			UserId:   userId,
			Status:   gtproto.OnlineStatus_OFFLINE,
			LastSeen: sampleDate(),
		},
	}
}

func unixTime(dt time.Time) uint64 {
	return uint64(dt.UTC().UnixMilli())
}
