package button

import (
	"fmt"
	domain2 "github.com/sepuka/focalism/internal/domain"
	"github.com/sepuka/vkbotserver/api/button"
)

const (
	StartIdButton     = `start`
	NextIdButton      = `next`
	SurrenderIdButton = `surrender`
	ReturnIdButton    = `return`
	IrregularIdButton = `irregular`
	TopicsIdButton    = `topics`
	ProgressIdButton  = `progress`

	TextButtonType      button.Type = `text`
	NextLabel           button.Text = `ещё слово`
	ReturnLabel         button.Text = `назад`
	GetAnswerLabel      button.Text = `не знаю`
	IrregularVerbsLabel button.Text = `неправильные глаголы`
	TopicsLabel         button.Text = `темы`
	ProgressLabel       button.Text = `прогресс`
)

func progress(topicId string) button.Button {
	return button.Button{
		Color: button.PrimaryColor,
		Action: button.Action{
			Type:  TextButtonType,
			Label: ProgressLabel,
			Payload: button.Payload{
				Command: ProgressIdButton,
				Id:      topicId,
			}.String(),
		},
	}
}

func SurrenderAndReturn(taskId string) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: GetAnswerLabel,
					Payload: button.Payload{
						Command: SurrenderIdButton,
						Id:      taskId,
					}.String(),
				},
			},
			{
				Color: button.SecondaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ReturnLabel,
					Payload: button.Payload{
						Command: ReturnIdButton,
					}.String(),
				},
			},
		},
	}
}

func Return() [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.SecondaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ReturnLabel,
					Payload: button.Payload{
						Command: ReturnIdButton,
					}.String(),
				},
			},
		},
	}
}

func ReturnWithProgress(topicId string) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.SecondaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ReturnLabel,
					Payload: button.Payload{
						Command: ReturnIdButton,
					}.String(),
				},
			},
			progress(topicId),
		},
	}
}

func ModeChoose() [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: IrregularVerbsLabel,
					Payload: button.Payload{
						Command: IrregularIdButton,
					}.String(),
				},
			},
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: TopicsLabel,
					Payload: button.Payload{
						Command: TopicsIdButton,
					}.String(),
				},
			},
		},
	}
}

func NextWithReturn(topicId string) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.SecondaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ReturnLabel,
					Payload: button.Payload{
						Command: ReturnIdButton,
					}.String(),
				},
			},
			{
				Color: button.PositiveColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: NextLabel,
					Payload: button.Payload{
						Command: NextIdButton,
						Id:      topicId,
					}.String(),
				},
			},
		},
	}
}

func NextWithReturnAndProgress(topicId string) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.SecondaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ReturnLabel,
					Payload: button.Payload{
						Command: ReturnIdButton,
					}.String(),
				},
			},
			{
				Color: button.PositiveColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: NextLabel,
					Payload: button.Payload{
						Command: NextIdButton,
						Id:      topicId,
					}.String(),
				},
			},
			progress(topicId),
		},
	}
}

func ReturnWithTopicCouples(topics []domain2.Topic) [][]button.Button {
	var (
		topic        domain2.Topic
		returnButton = Return()
	)

	switch len(topics) {
	case 0:
		return Return()
	case 1:
		var topicButton = button.Button{
			Color: button.PrimaryColor,
			Action: button.Action{
				Type:  TextButtonType,
				Label: button.Text(topic.Title),
				Payload: button.Payload{
					Command: NextIdButton,
					Id:      fmt.Sprintf(`%d`, topic.TopicId),
				}.String(),
			},
		}
		returnButton[0] = append(returnButton[0], topicButton)

		return returnButton
	default:
		var (
			topicId     int
			buttons     = returnButton
			topicButton button.Button
			rowButton   []button.Button
		)

		for topicId, topic = range topics {
			topicButton = button.Button{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: button.Text(topic.Title),
					Payload: button.Payload{
						Command: NextIdButton,
						Id:      fmt.Sprintf(`%d`, topic.TopicId),
					}.String(),
				},
			}
			rowButton = append(rowButton, topicButton)

			if topicId%2 != 0 {
				buttons = append(buttons, rowButton)
				rowButton = nil
			}
		}

		if len(rowButton) > 0 {
			buttons = append(buttons, rowButton)
		}

		return buttons
	}
}
