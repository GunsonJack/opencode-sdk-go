// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/GunsonJack/opencode-sdk-go"
	"github.com/GunsonJack/opencode-sdk-go/internal/testutil"
	"github.com/GunsonJack/opencode-sdk-go/option"
)

func TestTuiAppendPromptWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.AppendPrompt(context.TODO(), opencode.TuiAppendPromptParams{
		Text:      opencode.F("text"),
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiClearPromptWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.ClearPrompt(context.TODO(), opencode.TuiClearPromptParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiExecuteCommandWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.ExecuteCommand(context.TODO(), opencode.TuiExecuteCommandParams{
		Command:   opencode.F("prompt.submit"),
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiOpenHelpWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.OpenHelp(context.TODO(), opencode.TuiOpenHelpParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiOpenModelsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.OpenModels(context.TODO(), opencode.TuiOpenModelsParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiOpenSessionsWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.OpenSessions(context.TODO(), opencode.TuiOpenSessionsParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiOpenThemesWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.OpenThemes(context.TODO(), opencode.TuiOpenThemesParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiPublish(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Tui.Publish(context.TODO(), opencode.TuiPublishParams{
		Body: opencode.F[opencode.TuiPublishBody](opencode.TuiPublishBodyToastShow{
			Properties: opencode.F(opencode.TuiPublishBodyToastShowProperties{
				Message: opencode.F("message"),
				Variant: opencode.F(opencode.TuiShowToastParamsVariantInfo),
			}),
		}),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiPublishBodyMarshalingMatchesSpec(t *testing.T) {
	tests := []struct {
		name string
		body opencode.TuiPublishBody
		want string
	}{
		{
			name: "prompt append",
			body: opencode.TuiPublishBodyPromptAppend{
				Properties: opencode.F(opencode.TuiPublishBodyPromptAppendProperties{Text: opencode.F("hello")}),
			},
			want: `{"properties":{"text":"hello"},"type":"tui.prompt.append"}`,
		},
		{
			name: "command execute",
			body: opencode.TuiPublishBodyCommandExecute{
				Properties: opencode.F(opencode.TuiPublishBodyCommandExecuteProperties{Command: opencode.F(opencode.TuiCommandSessionNew)}),
			},
			want: `{"properties":{"command":"session.new"},"type":"tui.command.execute"}`,
		},
		{
			name: "toast show",
			body: opencode.TuiPublishBodyToastShow{
				Properties: opencode.F(opencode.TuiPublishBodyToastShowProperties{
					Message: opencode.F("Done!"),
					Variant: opencode.F(opencode.TuiShowToastParamsVariantSuccess),
					Title:   opencode.F("Result"),
				}),
			},
			want: `{"properties":{"message":"Done!","title":"Result","variant":"success"},"type":"tui.toast.show"}`,
		},
		{
			name: "session select",
			body: opencode.TuiPublishBodySessionSelect{
				Properties: opencode.F(opencode.TuiPublishBodySessionSelectProperties{SessionID: opencode.F("ses_123")}),
			},
			want: `{"properties":{"sessionID":"ses_123"},"type":"tui.session.select"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.body)
			if err != nil {
				t.Fatalf("json.Marshal() error = %v", err)
			}
			if string(got) != tt.want {
				t.Fatalf("json.Marshal() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestTuiPublishBodyMarshalingRequiresProperties(t *testing.T) {
	tests := []struct {
		name string
		body opencode.TuiPublishBody
		want string
	}{
		{
			name: "prompt append text",
			body: opencode.TuiPublishBodyPromptAppend{Properties: opencode.F(opencode.TuiPublishBodyPromptAppendProperties{})},
			want: "missing required properties.text for tui.prompt.append",
		},
		{
			name: "command execute command",
			body: opencode.TuiPublishBodyCommandExecute{Properties: opencode.F(opencode.TuiPublishBodyCommandExecuteProperties{})},
			want: "missing required properties.command for tui.command.execute",
		},
		{
			name: "toast show message",
			body: opencode.TuiPublishBodyToastShow{Properties: opencode.F(opencode.TuiPublishBodyToastShowProperties{Variant: opencode.F(opencode.TuiShowToastParamsVariantInfo)})},
			want: "missing required properties.message for tui.toast.show",
		},
		{
			name: "toast show variant",
			body: opencode.TuiPublishBodyToastShow{Properties: opencode.F(opencode.TuiPublishBodyToastShowProperties{Message: opencode.F("Done!")})},
			want: "missing required properties.variant for tui.toast.show",
		},
		{
			name: "session select sessionID",
			body: opencode.TuiPublishBodySessionSelect{Properties: opencode.F(opencode.TuiPublishBodySessionSelectProperties{})},
			want: "missing required properties.sessionID for tui.session.select",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := json.Marshal(tt.body)
			if err == nil {
				t.Fatal("json.Marshal() error = nil, want error")
			}
			if !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("json.Marshal() error = %q, want substring %q", err.Error(), tt.want)
			}
		})
	}
}

func TestTuiSelectSession(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Tui.SelectSession(context.TODO(), opencode.TuiSelectSessionParams{
		SessionID: opencode.F("ses_123"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiShowToastWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.ShowToast(context.TODO(), opencode.TuiShowToastParams{
		Message:   opencode.F("message"),
		Variant:   opencode.F(opencode.TuiShowToastParamsVariantInfo),
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
		Title:     opencode.F("title"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiSubmitPromptWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Tui.SubmitPrompt(context.TODO(), opencode.TuiSubmitPromptParams{
		Directory: opencode.F("directory"),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiControlNext(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Tui.Control.Next(context.TODO(), opencode.TuiControlNextParams{
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTuiControlResponse(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(option.WithBaseURL(baseURL))
	_, err := client.Tui.Control.Response(context.TODO(), opencode.TuiControlResponseParams{
		Body:      opencode.F[interface{}](map[string]interface{}{"ok": true}),
		Workspace: opencode.F("workspace"),
	})
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
