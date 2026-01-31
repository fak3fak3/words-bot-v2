package repos

type Word struct {
	Word          string
	Transcription string
	Definitions   []WordDefinition
}

type WordDefinition struct {
	PartOfSpeech string
	Meaning      string
	Example      string
}
