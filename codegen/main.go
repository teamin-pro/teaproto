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

	"github.com/gamarjoba-team/gtnode/generated/gtproto"
)

//go:embed gtproto/version.txt
var gtprotoVersion string

func main() {
	state()
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
						Now:                      UnixTime(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
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

	mustSave("state-server-error.request.yml", &gtproto.ClientRequest{
		Actions: []*gtproto.Action{
			{
				Request: &gtproto.Action_StateRequest{
					StateRequest: &gtproto.StateRequest{
						Mode: gtproto.StateRequest_RETURN_SERVER_ERROR,
					},
				},
			},
		},
	})

	mustSave("state-server-error.response.yml", &gtproto.ServerResponse{
		Errors: []gtproto.ServerResponse_Error{
			gtproto.ServerResponse_INTERNAL_SERVER_ERROR,
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

func UnixTime(dt time.Time) uint64 {
	return uint64(dt.UTC().UnixMilli())
}
