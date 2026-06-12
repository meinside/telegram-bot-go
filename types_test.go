// types_test.go
//
// pure (offline) unit tests for un/marshalling of types

package telegrambot

import (
	"encoding/json"
	"log/slog"
	"testing"
)

// RichText can be a plain string.
func TestRichTextUnmarshalString(t *testing.T) {
	slog.Info("testing unmarshalling of plain-text RichText...")

	var rt RichText
	if err := json.Unmarshal([]byte(`"hello"`), &rt); err != nil {
		t.Fatalf("failed to unmarshal plain-text RichText: %s", err)
	}

	s, ok := rt.Value.(string)
	if !ok {
		t.Fatalf("expected RichText.Value to be string, got %T", rt.Value)
	}
	if s != "hello" {
		t.Errorf("expected %q, got %q", "hello", s)
	}
}

// RichText can be an array of RichText.
func TestRichTextUnmarshalArray(t *testing.T) {
	slog.Info("testing unmarshalling of array RichText...")

	var rt RichText
	if err := json.Unmarshal([]byte(`["a", "b"]`), &rt); err != nil {
		t.Fatalf("failed to unmarshal array RichText: %s", err)
	}

	arr, ok := rt.Value.([]RichText)
	if !ok {
		t.Fatalf("expected RichText.Value to be []RichText, got %T", rt.Value)
	}
	if len(arr) != 2 {
		t.Fatalf("expected 2 elements, got %d", len(arr))
	}
	if s, ok := arr[0].Value.(string); !ok || s != "a" {
		t.Errorf("expected first element %q, got %v", "a", arr[0].Value)
	}
}

// RichText can be a typed object discriminated by `type`, with a nested RichText.
func TestRichTextUnmarshalObject(t *testing.T) {
	slog.Info("testing unmarshalling of typed-object RichText...")

	var rt RichText
	if err := json.Unmarshal([]byte(`{"type":"bold","text":"hi"}`), &rt); err != nil {
		t.Fatalf("failed to unmarshal object RichText: %s", err)
	}

	bold, ok := rt.Value.(*RichTextBold)
	if !ok {
		t.Fatalf("expected RichText.Value to be *RichTextBold, got %T", rt.Value)
	}
	if bold.Type != "bold" {
		t.Errorf("expected type %q, got %q", "bold", bold.Type)
	}
	if s, ok := bold.Text.Value.(string); !ok || s != "hi" {
		t.Errorf("expected nested text %q, got %v", "hi", bold.Text.Value)
	}
}

// An unknown RichText object type should decode to a generic map without error.
func TestRichTextUnmarshalUnknownObject(t *testing.T) {
	slog.Info("testing unmarshalling of unknown-type RichText...")

	var rt RichText
	if err := json.Unmarshal([]byte(`{"type":"future_type","text":"x"}`), &rt); err != nil {
		t.Fatalf("failed to unmarshal unknown-type RichText: %s", err)
	}
	if rt.Value == nil {
		t.Errorf("expected a non-nil fallback value for unknown type")
	}
}

// RichBlock is a flat union struct discriminated by `type`.
func TestRichBlockUnmarshalParagraph(t *testing.T) {
	slog.Info("testing unmarshalling of paragraph RichBlock...")

	var rb RichBlock
	if err := json.Unmarshal([]byte(`{"type":"paragraph","text":"hi"}`), &rb); err != nil {
		t.Fatalf("failed to unmarshal paragraph RichBlock: %s", err)
	}
	if rb.Type != "paragraph" {
		t.Errorf("expected type %q, got %q", "paragraph", rb.Type)
	}
	if s, ok := rb.Text.Value.(string); !ok || s != "hi" {
		t.Errorf("expected text %q, got %v", "hi", rb.Text.Value)
	}
}

// A RichMessage (received inside a Message) decodes its blocks recursively.
func TestRichMessageUnmarshal(t *testing.T) {
	slog.Info("testing recursive unmarshalling of RichMessage...")

	var rm RichMessage
	data := `{"blocks":[{"type":"paragraph","text":[{"type":"bold","text":"hi"}]}],"is_rtl":false}`
	if err := json.Unmarshal([]byte(data), &rm); err != nil {
		t.Fatalf("failed to unmarshal RichMessage: %s", err)
	}
	if len(rm.Blocks) != 1 {
		t.Fatalf("expected 1 block, got %d", len(rm.Blocks))
	}
	if rm.Blocks[0].Type != "paragraph" {
		t.Errorf("expected paragraph, got %q", rm.Blocks[0].Type)
	}

	arr, ok := rm.Blocks[0].Text.Value.([]RichText)
	if !ok {
		t.Fatalf("expected paragraph text to be []RichText, got %T", rm.Blocks[0].Text.Value)
	}
	bold, ok := arr[0].Value.(*RichTextBold)
	if !ok {
		t.Fatalf("expected nested *RichTextBold, got %T", arr[0].Value)
	}
	if s, ok := bold.Text.Value.(string); !ok || s != "hi" {
		t.Errorf("expected nested bold text %q, got %v", "hi", bold.Text.Value)
	}
}

// Send-path compatibility: helpers still produce the same JSON as before.
func TestRichTextMarshalCompat(t *testing.T) {
	slog.Info("testing send-path marshalling compatibility of RichText...")

	// plain text -> bare JSON string
	if b, err := json.Marshal(NewRichTextWithText("hello")); err != nil {
		t.Fatalf("failed to marshal plain-text RichText: %s", err)
	} else if string(b) != `"hello"` {
		t.Errorf("expected %q, got %q", `"hello"`, string(b))
	}

	// array -> JSON array of bare strings
	if b, err := json.Marshal(NewRichTextWithRichTexts(
		NewRichTextWithText("a"),
		NewRichTextWithText("b"),
	)); err != nil {
		t.Fatalf("failed to marshal array RichText: %s", err)
	} else if string(b) != `["a","b"]` {
		t.Errorf("expected %q, got %q", `["a","b"]`, string(b))
	}
}
