package summarizer

import (
	"errors"

	"github.com/artemzi/summarizer/helpers"
)

// Summarizer instance, used for extracting summary from raw texts and urls
type Summarizer struct {
	url            string
	title          string
	fullText       string
	summarizedText string
	images         []string
	summarized     bool
}

// CreateFromURL creates summarizer instance, using the url parameter for summarizing
func CreateFromURL(url string) *Summarizer {
	var summarizer = new(Summarizer)
	summarizer.url = url
	return summarizer
}

// CreateFromText creates summarizer instance, using the text parameter for summarizing
func CreateFromText(text string) *Summarizer {
	var summarizer = new(Summarizer)
	summarizer.fullText = text
	return summarizer
}

// Summarize returns summary of the text, extracted from the url or the saved text
func (s *Summarizer) Summarize() (string, error) {
	if s.IsSummarized() {
		return s.summarizedText, nil
	}

	if s.fullText == "" && s.url == "" {
		return "", errors.New("You must submit text or url for summarizing")
	}

	if s.url != "" {
		s.GetMainTextFromURL()
	}

	var summarizedText = s.summarizeFromText()
	if len(summarizedText) == 0 {
		return "", errors.New("Something happened while summarizing. Please try again")
	}

	s.summarizedText = summarizedText
	s.summarized = true
	if s.title != "" {
		return s.title + "\n\n" + s.summarizedText, nil
	}

	return s.summarizedText, nil
}

// GetMainTextFromURL parses the summarizer object URL property and returns the main text
// from the website without ads, unnecessary images and other not important data
func (s *Summarizer) GetMainTextFromURL() (string, error) {
	if s.url == "" {
		return "", errors.New("You must use summarizer from URL")
	}

	if s.fullText != "" {
		return s.fullText, nil
	}

	extractedTitle, extractedText, extractedImages, err := helpers.ExtractMainInfoFromURL(s.url)
	if err != nil {
		return "", err
	}

	s.title = extractedTitle
	s.fullText = extractedText
	s.images = extractedImages

	return extractedTitle + "\n\n" + extractedText, nil
}

func (s *Summarizer) summarizeFromText() string {
	// Build the summary with the sentences dictionary
	var summary = helpers.GetSummary(s.fullText)
	return summary
}

// GetSummaryInfo returns summary information statistics if the text is summarized and an error if not
func (s *Summarizer) GetSummaryInfo() (string, error) {
	if !s.IsSummarized() {
		return "", errors.New("You must first summarize the text in order to get information for it")
	}

	var summaryInfo = helpers.GetSummaryInfo(s.fullText, s.summarizedText, len(s.images))
	return summaryInfo, nil
}

// IsSummarized checks if the instance was already summarized
func (s *Summarizer) IsSummarized() bool {
	return s.summarized
}

// StoreToFile stores the summarized text to the file from the given path
func (s *Summarizer) StoreToFile(filePath string) (bool, error) {
	if !s.IsSummarized() {
		return false, errors.New("You must first summarize the text in order to save the summary to a file")
	}

	stored, err := helpers.StoreTextToFile(filePath, s.title, s.summarizedText, s.images)
	return stored, err
}
