package button

import "github.com/sepuka/vkbotserver/api/button"

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

func Surrender(taskId string) [][]button.Button {
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
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  TextButtonType,
					Label: ProgressLabel,
					Payload: button.Payload{
						Command: ProgressIdButton,
						Id:      topicId,
					}.String(),
				},
			},
		},
	}
}
