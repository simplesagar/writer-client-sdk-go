// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ContentIssueServiceEnum string

const (
	ContentIssueServiceEnumCommonMistakes     ContentIssueServiceEnum = "common-mistakes"
	ContentIssueServiceEnumBannedWords        ContentIssueServiceEnum = "banned-words"
	ContentIssueServiceEnumDictionary         ContentIssueServiceEnum = "dictionary"
	ContentIssueServiceEnumGec                ContentIssueServiceEnum = "gec"
	ContentIssueServiceEnumInfinitive         ContentIssueServiceEnum = "infinitive"
	ContentIssueServiceEnumSpelling           ContentIssueServiceEnum = "spelling"
	ContentIssueServiceEnumWritingStyle       ContentIssueServiceEnum = "writing-style"
	ContentIssueServiceEnumCustomRules        ContentIssueServiceEnum = "custom-rules"
	ContentIssueServiceEnumSentenceCase       ContentIssueServiceEnum = "sentence-case"
	ContentIssueServiceEnumAcronym            ContentIssueServiceEnum = "acronym"
	ContentIssueServiceEnumOxfordComma        ContentIssueServiceEnum = "oxford-comma"
	ContentIssueServiceEnumMlPunctuation      ContentIssueServiceEnum = "ml-punctuation"
	ContentIssueServiceEnumEmojis             ContentIssueServiceEnum = "emojis"
	ContentIssueServiceEnumGenderPronouns     ContentIssueServiceEnum = "gender-pronouns"
	ContentIssueServiceEnumSensitivity        ContentIssueServiceEnum = "sensitivity"
	ContentIssueServiceEnumPlagiarism         ContentIssueServiceEnum = "plagiarism"
	ContentIssueServiceEnumReadability        ContentIssueServiceEnum = "readability"
	ContentIssueServiceEnumSentenceComplexity ContentIssueServiceEnum = "sentence-complexity"
	ContentIssueServiceEnumVocabulary         ContentIssueServiceEnum = "vocabulary"
	ContentIssueServiceEnumParagraphLength    ContentIssueServiceEnum = "paragraph-length"
	ContentIssueServiceEnumPlainLanguage      ContentIssueServiceEnum = "plain-language"
	ContentIssueServiceEnumHealthyCommn       ContentIssueServiceEnum = "healthy-commn"
	ContentIssueServiceEnumConfidence         ContentIssueServiceEnum = "confidence"
	ContentIssueServiceEnumDataLossPrevention ContentIssueServiceEnum = "data-loss-prevention"
	ContentIssueServiceEnumHateSpeech         ContentIssueServiceEnum = "hate-speech"
	ContentIssueServiceEnumContentSafeguards  ContentIssueServiceEnum = "content-safeguards"
	ContentIssueServiceEnumFeedback           ContentIssueServiceEnum = "feedback"
	ContentIssueServiceEnumClaim              ContentIssueServiceEnum = "claim"
	ContentIssueServiceEnumQuote              ContentIssueServiceEnum = "quote"
	ContentIssueServiceEnumGenderNouns        ContentIssueServiceEnum = "gender-nouns"
	ContentIssueServiceEnumGenderTone         ContentIssueServiceEnum = "gender-tone"
	ContentIssueServiceEnumGrammar            ContentIssueServiceEnum = "grammar"
	ContentIssueServiceEnumPunctuationDark    ContentIssueServiceEnum = "punctuation-dark"
	ContentIssueServiceEnumFormatting         ContentIssueServiceEnum = "formatting"
	ContentIssueServiceEnumTwitter            ContentIssueServiceEnum = "twitter"
	ContentIssueServiceEnumGecDark            ContentIssueServiceEnum = "gec-dark"
	ContentIssueServiceEnumGecGpt3            ContentIssueServiceEnum = "gec-gpt3"
)

func (e ContentIssueServiceEnum) ToPointer() *ContentIssueServiceEnum {
	return &e
}

func (e *ContentIssueServiceEnum) UnmarshalJSON(data []byte) error {
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
		*e = ContentIssueServiceEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ContentIssueServiceEnum: %v", v)
	}
}

type ContentIssue struct {
	Description *string                 `json:"description,omitempty"`
	From        int64                   `json:"from"`
	Meta        interface{}             `json:"meta,omitempty"`
	Service     ContentIssueServiceEnum `json:"service"`
	Suggestions []string                `json:"suggestions,omitempty"`
	Until       int64                   `json:"until"`
}
