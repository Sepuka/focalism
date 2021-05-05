package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/focalism/internal/repository/mocks"
	api2 "github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	mocks2 "github.com/sepuka/vkbotserver/api/mocks"
	"github.com/sepuka/vkbotserver/config"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
)

func TestNewTipHandler_Handle(t *testing.T) {
	var (
		taskRepository = mocks.TaskRepository{}
		logger         = zap.NewNop().Sugar()
		client         mocks2.HTTPClient
		rnd            = mocks2.Rnder{}
		rndId          = int64(777)
		cfg            = config.Config{
			Api: config.Api{
				Token: `some_token`,
			},
		}
		handler MessageHandler
		api     *api2.Api

		errMsg      string
		err         error
		keyboard    []byte
		httpRequest *http.Request
		response    *http.Response
		endpoint    = fmt.Sprintf(`%s/%s`, api2.Endpoint, `messages.send`) //TODO fix const vendor/github.com/sepuka/vkbotserver/api/api.go:19
		answer, _   = json.Marshal(&api2.Response{})

		payload = api2.OutcomeMessage{
			AccessToken: cfg.Api.Token,
			ApiVersion:  api2.Version,
			RandomId:    rndId,
			PeerId:      1,
		}
		testCases = map[string]struct {
			taskPayload   *button.Payload
			clientRequest *domain.Request
			answer        string
			tip           string
			example       string
			taskId        int64
		}{
			`answer only`: {
				taskPayload: &button.Payload{
					Command: "",
					Id:      `0`,
				},
				clientRequest: &domain.Request{
					Object: domain.Object{
						Message: domain.Message{
							FromId: 1,
						},
					},
				},
				answer: `Answer1`,
				tip:    `"A*****1"`,
				taskId: 0,
			},
			`answer with example`: {
				taskPayload: &button.Payload{
					Command: "",
					Id:      `1`,
				},
				clientRequest: &domain.Request{
					Object: domain.Object{
						Message: domain.Message{
							FromId: 1,
						},
					},
				},
				answer:  `Answer2`,
				tip:     `"A*****2"`,
				taskId:  1,
				example: `There is some example`,
			},
		}
	)

	rnd.On(`Rnd`).Return(rndId)

	for caseName, caseValue := range testCases {
		keyboard, _ = json.Marshal(button.Keyboard{
			OneTime: true,
			Buttons: button2.SurrenderAndReturn(strconv.Itoa(int(caseValue.taskId))),
		})

		payload.Keyboard = string(keyboard)
		payload.Message = caseValue.tip
		if caseValue.example != `` {
			payload.Message = fmt.Sprintf("%s\n\n%s", payload.Message, caseValue.example)
		}
		params, _ := query.Values(payload)
		httpRequest, _ = http.NewRequest(`POST`, fmt.Sprintf(`%s?%s`, endpoint, params.Encode()), nil)
		response = &http.Response{
			Body: ioutil.NopCloser(bytes.NewReader(answer)),
		}

		client = mocks2.HTTPClient{}
		client.On(`Do`, httpRequest).Return(response, nil)

		taskRepository.On(`GetById`, caseValue.taskId).Return(domain2.Task{
			Id: caseValue.taskId,
			Vocabulary: &domain2.Vocabulary{
				Answer:  caseValue.answer,
				Example: caseValue.example,
			},
		}, nil)

		api = api2.NewApi(logger, cfg, client, rnd)
		handler = NewTipHandler(api, taskRepository)

		err = handler.Handle(caseValue.clientRequest, caseValue.taskPayload)
		if err != nil {
			errMsg = fmt.Sprintf(`test "%s" fall`, caseName)
			t.Errorf(`%s. error result is not expected`, errMsg)
		}
	}
}
