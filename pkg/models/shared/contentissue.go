// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Service string

const (
	ServiceCommonMistakes     Service = "common-mistakes"
	ServiceBannedWords        Service = "banned-words"
	ServiceDictionary         Service = "dictionary"
	ServiceGec                Service = "gec"
	ServiceInfinitive         Service = "infinitive"
	ServiceSpelling           Service = "spelling"
	ServiceWritingStyle       Service = "writing-style"
	ServiceCustomRules        Service = "custom-rules"
	ServiceSentenceCase       Service = "sentence-case"
	ServiceAcronym            Service = "acronym"
	ServiceOxfordComma        Service = "oxford-comma"
	ServiceMlPunctuation      Service = "ml-punctuation"
	ServiceEmojis             Service = "emojis"
	ServiceGenderPronouns     Service = "gender-pronouns"
	ServiceSensitivity        Service = "sensitivity"
	ServicePlagiarism         Service = "plagiarism"
	ServiceReadability        Service = "readability"
	ServiceSentenceComplexity Service = "sentence-complexity"
	ServiceVocabulary         Service = "vocabulary"
	ServiceParagraphLength    Service = "paragraph-length"
	ServicePlainLanguage      Service = "plain-language"
	ServiceHealthyCommn       Service = "healthy-commn"
	ServiceConfidence         Service = "confidence"
	ServiceDataLossPrevention Service = "data-loss-prevention"
	ServiceHateSpeech         Service = "hate-speech"
	ServiceContentSafeguards  Service = "content-safeguards"
	ServiceFeedback           Service = "feedback"
	ServiceClaim              Service = "claim"
	ServiceQuote              Service = "quote"
	ServiceGenderNouns        Service = "gender-nouns"
	ServiceGenderTone         Service = "gender-tone"
	ServiceGrammar            Service = "grammar"
	ServicePunctuationDark    Service = "punctuation-dark"
	ServiceFormatting         Service = "formatting"
	ServiceTwitter            Service = "twitter"
	ServiceGecDark            Service = "gec-dark"
	ServiceGecGpt3            Service = "gec-gpt3"
)

func (e Service) ToPointer() *Service {
	return &e
}

func (e *Service) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "common-mistakes":
		fallthrough
	case "banned-words":
		fallthrough
	case "dictionary":
		fallthrough
	case "gec":
		fallthrough
	case "infinitive":
		fallthrough
	case "spelling":
		fallthrough
	case "writing-style":
		fallthrough
	case "custom-rules":
		fallthrough
	case "sentence-case":
		fallthrough
	case "acronym":
		fallthrough
	case "oxford-comma":
		fallthrough
	case "ml-punctuation":
		fallthrough
	case "emojis":
		fallthrough
	case "gender-pronouns":
		fallthrough
	case "sensitivity":
		fallthrough
	case "plagiarism":
		fallthrough
	case "readability":
		fallthrough
	case "sentence-complexity":
		fallthrough
	case "vocabulary":
		fallthrough
	case "paragraph-length":
		fallthrough
	case "plain-language":
		fallthrough
	case "healthy-commn":
		fallthrough
	case "confidence":
		fallthrough
	case "data-loss-prevention":
		fallthrough
	case "hate-speech":
		fallthrough
	case "content-safeguards":
		fallthrough
	case "feedback":
		fallthrough
	case "claim":
		fallthrough
	case "quote":
		fallthrough
	case "gender-nouns":
		fallthrough
	case "gender-tone":
		fallthrough
	case "grammar":
		fallthrough
	case "punctuation-dark":
		fallthrough
	case "formatting":
		fallthrough
	case "twitter":
		fallthrough
	case "gec-dark":
		fallthrough
	case "gec-gpt3":
		*e = Service(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Service: %v", v)
	}
}

type ContentIssue struct {
	Description *string     `json:"description,omitempty"`
	From        int64       `json:"from"`
	Meta        interface{} `json:"meta,omitempty"`
	Service     Service     `json:"service"`
	Suggestions []string    `json:"suggestions,omitempty"`
	Until       int64       `json:"until"`
}

func (o *ContentIssue) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ContentIssue) GetFrom() int64 {
	if o == nil {
		return 0
	}
	return o.From
}

func (o *ContentIssue) GetMeta() interface{} {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *ContentIssue) GetService() Service {
	if o == nil {
		return Service("")
	}
	return o.Service
}

func (o *ContentIssue) GetSuggestions() []string {
	if o == nil {
		return nil
	}
	return o.Suggestions
}

func (o *ContentIssue) GetUntil() int64 {
	if o == nil {
		return 0
	}
	return o.Until
}
