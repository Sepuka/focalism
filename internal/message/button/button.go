package button

import "github.com/sepuka/vkbotserver/api/button"

const (
	ButtonIdStart     = `start`
	ButtonIdNext      = `next`
	ButtonIdSurrender = `surrender`
	ButtonIdReturn    = `return`
	ButtonIdIrregular = `irregular`
	ButtonIdTopics    = `topics`

	TextButtonType      button.Type = `text`
	NextLabel           button.Text = `ещё`
	ReturnLabel         button.Text = `назад`
	GetAnswerLabel      button.Text = `не знаю`
	IrregularVerbsLabel button.Text = `неправильные глаголы`
	TopicsLabel         button.Text = `темы`
)

func Surrender(taskId string) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: GetAnswerLabel,
					Payload: button.Payload{
						Command: ButtonIdSurrender,
						Id:      taskId,
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
						Command: ButtonIdReturn,
					}.String(),
				},
			},
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
						Command: ButtonIdIrregular,
					}.String(),
				},
			},
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: TopicsLabel,
					Payload: button.Payload{
						Command: ButtonIdTopics,
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
						Command: ButtonIdReturn,
					}.String(),
				},
			},
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: NextLabel,
					Payload: button.Payload{
						Command: ButtonIdNext,
						Id:      topicId,
					}.String(),
				},
			},
		},
	}
}

func Answer(payload button.Payload) [][]button.Button {
	return [][]button.Button{
		{
			{
				Color: button.NegativeColor,
				Action: button.Action{
					Type:    TextButtonType,
					Label:   GetAnswerLabel,
					Payload: payload.String(),
				},
			},
		},
	}
}
