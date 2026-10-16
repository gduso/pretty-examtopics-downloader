package fetch

import (
	"strings"
	"testing"

	"examtopics-downloader/internal/models"

	"github.com/PuerkitoBio/goquery"
)

func TestExtractInlineOptions(t *testing.T) {
	html := `
<div class="question-body" data-id="1">
  <p class="card-text">
    A company runs an app.<br />
A. sentence that starts like a choice is part of the stem.<br />
Which solution meets the requirement?<br>
    <br>A.
      Use Amazon SQS.<br>
    <br>B.
      Use Amazon SNS.<br>
    <br>C.
      Use <span>AWS</span> Lambda<br />with a second line.<br>
  </p>
</div>
<p class="card-text">Login modal text</p>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}

	stem, options, ok := extractInlineOptions(doc)
	if !ok {
		t.Fatalf("expected inline options to be found")
	}

	wantStem := "A company runs an app.\nA. sentence that starts like a choice is part of the stem.\nWhich solution meets the requirement?"
	if stem != wantStem {
		t.Fatalf("unexpected stem:\nwant %q\ngot  %q", wantStem, stem)
	}

	wantOptions := []string{
		"A. Use Amazon SQS.",
		"B. Use Amazon SNS.",
		"C. Use AWS Lambda with a second line.",
	}
	if strings.Join(options, "|") != strings.Join(wantOptions, "|") {
		t.Fatalf("unexpected options:\nwant %q\ngot  %q", wantOptions, options)
	}
}

func TestExtractInlineOptionsWithoutChoices(t *testing.T) {
	html := `<div class="question-body"><p class="card-text">HOTSPOT -<br>Select the answer.<br><img src="/a.png"></p></div>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}

	if _, options, ok := extractInlineOptions(doc); ok {
		t.Fatalf("expected no inline options, got %q", options)
	}
}

func TestExtractDiscussionCommentsSelectedAnswerBadge(t *testing.T) {
	html := `
<div class="discussion-container">
  <div class="media comment-container">
    <div class="media-body">
      <div><h5 class="comment-username">alice</h5><span class="badge badge-primary">Most Recent</span></div>
      <div class="comment-body">
        <div class="badge badge-warning">Selected Answer: <span><strong>D</strong></span></div>
        <div class="comment-content">D is the answer.</div>
      </div>
    </div>
  </div>
  <div class="media comment-container">
    <div class="media-body">
      <div><h5 class="comment-username">bob</h5></div>
      <div class="comment-body">
        <div class="badge badge-warning">Selected Answer: <span><strong>BD</strong></span></div>
        <div class="comment-content">Pick two.</div>
      </div>
    </div>
  </div>
  <div class="media comment-container">
    <div class="media-body">
      <div><h5 class="comment-username">carol</h5></div>
      <div class="comment-body"><div class="comment-content">No vote here.</div></div>
    </div>
  </div>
</div>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("failed to parse html: %v", err)
	}

	comments := extractDiscussionComments(doc)
	if len(comments) != 3 {
		t.Fatalf("expected 3 comments, got %d", len(comments))
	}
	want := []string{"D", "BD", ""}
	for i, c := range comments {
		if c.Answer != want[i] {
			t.Fatalf("comment %d: want answer %q, got %q", i, want[i], c.Answer)
		}
	}
}

func TestMostSelectedCommentAnswer(t *testing.T) {
	comments := []models.CommentData{
		{Answer: "A"},
		{Answer: "D"},
		{Answer: ""},
		{Answer: "D"},
		{Answer: "A"},
	}
	if got := mostSelectedCommentAnswer(comments); got != "D" {
		t.Fatalf("expected D, got %q", got)
	}
	if got := mostSelectedCommentAnswer(nil); got != "" {
		t.Fatalf("expected empty answer, got %q", got)
	}
}
