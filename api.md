# Shared Response Types

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared#MessageAbortedError">MessageAbortedError</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared#ProviderAuthError">ProviderAuthError</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/shared#UnknownError">UnknownError</a>

# Event

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#EventListResponse">EventListResponse</a>

The `EventListResponse` union has 49 variants covering session lifecycle, messages, permissions, questions, TUI control, MCP, PTY, workspace, worktree, VCS, and infrastructure events. Use `AsUnion()` and type-switch to handle specific event types.

Methods:

- <code title="get /event">client.Event.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#EventService.ListStreaming">ListStreaming</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#EventListParams">EventListParams</a>) *<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/packages/ssestream">ssestream</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/packages/ssestream#Stream">Stream</a>[<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#EventListResponse">EventListResponse</a>]</code>

# Path

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Path">Path</a>

Note: `Path` now includes `Home` field.

Methods:

- <code title="get /path">client.Path.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PathService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PathGetParams">PathGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Path">Path</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# App

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ModelCapabilities">ModelCapabilities</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ModelCapabilitiesModality">ModelCapabilitiesModality</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ModelCapabilitiesInterleaved">ModelCapabilitiesInterleaved</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ModelAPI">ModelAPI</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ModelCostCache">ModelCostCache</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Provider">Provider</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderSource">ProviderSource</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigProvidersResponse">ConfigProvidersResponse</a>

Note: `AppProvidersResponse` is deprecated; app/config providers use `ConfigProvidersResponse` with `Providers` and `Default` fields. The deprecated `App.Providers` wrapper still accepts both `directory` and `workspace` query params for compatibility.

Methods:

- <code title="post /log">client.App.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AppService.Log">Log</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AppLogParams">AppLogParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config/providers">client.App.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AppService.Providers">Providers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AppProvidersParams">AppProvidersParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigProvidersResponse">ConfigProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code> **Deprecated: use Config.Providers**

# Provider

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderListResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethod</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodType</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPrompt</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptUnion</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptText</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptTextType</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptSelect</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptSelectType</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptSelectOption</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptWhen</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthMethodPromptWhenOp</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthAuthorization</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">ProviderAuthAuthorizationMethod</a>

Methods:

- <code title="get /provider">client.Provider.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderListParams">ProviderListParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderListResponse">ProviderListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /provider/auth">client.Provider.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderService.Auth">Auth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderAuthParams">ProviderAuthParams</a>) (map[string][]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderAuthMethod">ProviderAuthMethod</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /provider/{providerID}/oauth/authorize">client.Provider.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderService.OAuthAuthorize">OAuthAuthorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderOAuthAuthorizeParams">ProviderOAuthAuthorizeParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderAuthAuthorization">ProviderAuthAuthorization</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /provider/{providerID}/oauth/callback">client.Provider.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderService.OAuthCallback">OAuthCallback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProviderOAuthCallbackParams">ProviderOAuthCallbackParams</a>) (bool, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Auth

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">Auth</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">AuthOAuth</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">AuthAPI</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">AuthWellKnown</a>

Methods:

- <code title="put /auth/{providerID}">client.Auth.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AuthService.Set">Set</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AuthSetParams">AuthSetParams</a>) (bool, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /auth/{providerID}">client.Auth.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AuthService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AuthRemoveParams">AuthRemoveParams</a>) (bool, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agent

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Agent">Agent</a>

Note: Agent now uses `[]PermissionRule` for permissions (defined in Session response types).

Methods:

- <code title="get /agent">client.Agent.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AgentListParams">AgentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Find

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Symbol">Symbol</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindTextResponse">FindTextResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindFilesParamsDirs">FindFilesParamsDirs</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindFilesParamsType">FindFilesParamsType</a>

Note: `FindFilesParams` now includes `Dirs`, `Type`, and `Limit` fields.

Methods:

- <code title="get /find/file">client.Find.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindService.Files">Files</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindFilesParams">FindFilesParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /find/symbol">client.Find.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindService.Symbols">Symbols</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindSymbolsParams">FindSymbolsParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Symbol">Symbol</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /find">client.Find.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindService.Text">Text</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindTextParams">FindTextParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FindTextResponse">FindTextResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# File

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#File">File</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileNode">FileNode</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileReadResponse">FileReadResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileReadResponseType">FileReadResponseType</a>

Note: `FileReadResponseType` now includes `binary`.

Methods:

- <code title="get /file">client.File.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileListParams">FileListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileNode">FileNode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /file/content">client.File.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileService.Read">Read</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileReadParams">FileReadParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileReadResponse">FileReadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /file/status">client.File.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileStatusParams">FileStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Config">Config</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ServerConfig">ServerConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpLocalConfig">McpLocalConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpRemoteConfig">McpRemoteConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigSkills">ConfigSkills</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigCompaction">ConfigCompaction</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigEnterprise">ConfigEnterprise</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigAgentEntry">ConfigAgentEntry</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigAgentEntryMode">ConfigAgentEntryMode</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigLogLevel">ConfigLogLevel</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigAutoupdateUnion">ConfigAutoupdateUnion</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigAutoupdateString">ConfigAutoupdateString</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigPluginItem">ConfigPluginItem</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionActionConfig">PermissionActionConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionObjectConfig">PermissionObjectConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionConfigObject">PermissionConfigObject</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionConfig">PermissionConfig</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigProvidersResponse">ConfigProvidersResponse</a>

Note: deprecated `Config.Mode` remains available as a typed compatibility wrapper with `Build` and `Plan` fields.

Methods:

- <code title="get /config">client.Config.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigGetParams">ConfigGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Config">Config</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /config">client.Config.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigUpdateParams">ConfigUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Config">Config</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config/providers">client.Config.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigService.Providers">Providers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigProvidersParams">ConfigProvidersParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConfigProvidersResponse">ConfigProvidersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Command

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Command">Command</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#CommandSource">CommandSource</a>

Note: `Command` now includes `Source` and `Hints` fields.

Methods:

- <code title="get /command">client.Command.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#CommandService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#CommandListParams">CommandListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Command">Command</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Project

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectIcon">ProjectIcon</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectCommands">ProjectCommands</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectSummary">ProjectSummary</a>

Methods:

- <code title="get /project">client.Project.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectListParams">ProjectListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/current">client.Project.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectService.Current">Current</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectCurrentParams">ProjectCurrentParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /project/{projectID}">client.Project.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /project/git/init">client.Project.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectService.InitGit">InitGit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ProjectInitGitParams">ProjectInitGitParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Session

Params Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AgentPartInputParam">AgentPartInputParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePartInputParam">FilePartInputParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePartSourceUnionParam">FilePartSourceUnionParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePartSourceTextParam">FilePartSourceTextParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileSourceParam">FileSourceParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SymbolSourceParam">SymbolSourceParam</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TextPartInputParam">TextPartInputParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AgentPart">AgentPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#AssistantMessage">AssistantMessage</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePart">FilePart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePartSource">FilePartSource</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FilePartSourceText">FilePartSourceText</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FileSource">FileSource</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Message">Message</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Part">Part</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ReasoningPart">ReasoningPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SnapshotPart">SnapshotPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#StepFinishPart">StepFinishPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#StepStartPart">StepStartPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SymbolSource">SymbolSource</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TextPart">TextPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolPart">ToolPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolStateCompleted">ToolStateCompleted</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolStateError">ToolStateError</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolStatePending">ToolStatePending</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolStateRunning">ToolStateRunning</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#UserMessage">UserMessage</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionCommandResponse">SessionCommandResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessageResponse">SessionMessageResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessagesResponse">SessionMessagesResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPromptResponse">SessionPromptResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionShellResponse">SessionShellResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SubtaskPart">SubtaskPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#CompactionPart">CompactionPart</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionRule">PermissionRule</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SnapshotFileDiff">SnapshotFileDiff</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionStatus">SessionStatus</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Todo">Todo</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#StructuredOutputError">StructuredOutputError</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ContextOverflowError">ContextOverflowError</a>

Methods:

- <code title="post /session">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionNewParams">SessionNewParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionUpdateParams">SessionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionListParams">SessionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionDeleteParams">SessionDeleteParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/abort">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Abort">Abort</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionAbortParams">SessionAbortParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/children">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Children">Children</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionChildrenParams">SessionChildrenParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/command">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Command">Command</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionCommandParams">SessionCommandParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionCommandResponse">SessionCommandResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

Note: `SessionCommandParams.Parts` is limited to file-part inputs (`type: "file"`) for compatibility with the current OpenAPI schema.
- <code title="get /session/{id}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionGetParams">SessionGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/init">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Init">Init</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionInitParams">SessionInitParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/message/{messageID}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Message">Message</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessageParams">SessionMessageParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessageResponse">SessionMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/message">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Messages">Messages</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessagesParams">SessionMessagesParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionMessagesResponse">SessionMessagesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/message">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Prompt">Prompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPromptParams">SessionPromptParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPromptResponse">SessionPromptResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/revert">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Revert">Revert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionRevertParams">SessionRevertParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/share">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Share">Share</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionShareParams">SessionShareParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/shell">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Shell">Shell</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionShellParams">SessionShellParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionShellResponse">SessionShellResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/summarize">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Summarize">Summarize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionSummarizeParams">SessionSummarizeParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/unrevert">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Unrevert">Unrevert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionUnrevertParams">SessionUnrevertParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}/share">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Unshare">Unshare</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionUnshareParams">SessionUnshareParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}/message/{messageID}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.DeleteMessage">DeleteMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionDeleteMessageParams">SessionDeleteMessageParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/diff">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Diff">Diff</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionDiffParams">SessionDiffParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SnapshotFileDiff">SnapshotFileDiff</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/fork">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Fork">Fork</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionForkParams">SessionForkParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Session">Session</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /session/{id}/prompt_async">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.PromptAsync">PromptAsync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPromptAsyncParams">SessionPromptAsyncParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

Note: `PromptAsync` is a `204 No Content` endpoint and reports success via `error` only.
- <code title="get /session/status">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionStatusParams">SessionStatusParams</a>) (map[string]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionStatus">SessionStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /session/{id}/todo">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.Todo">Todo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionTodoParams">SessionTodoParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Todo">Todo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /session/{id}/message/{messageID}/part/{partID}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.UpdatePart">UpdatePart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, partID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionUpdatePartParams">SessionUpdatePartParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Part">Part</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /session/{id}/message/{messageID}/part/{partID}">client.Session.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionService.DeletePart">DeletePart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, messageID <a href="https://pkg.go.dev/builtin#string">string</a>, partID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionDeletePartParams">SessionDeletePartParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Permissions

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Permission">Permission</a>

Methods:

- <code title="post /session/{id}/permissions/{permissionID}">client.Session.Permissions.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPermissionService.Respond">Respond</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, permissionID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SessionPermissionRespondParams">SessionPermissionRespondParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Permission

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionRequest">PermissionRequest</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionRequestTool">PermissionRequestTool</a>

Methods:

- <code title="get /permission">client.Permission.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionListParams">PermissionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionRequest">PermissionRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /permission/{requestID}/reply">client.Permission.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionService.Reply">Reply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PermissionReplyParams">PermissionReplyParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Question

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionRequest">QuestionRequest</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionInfo">QuestionInfo</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionOption">QuestionOption</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionTool">QuestionTool</a>

Methods:

- <code title="get /question">client.Question.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionListParams">QuestionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionRequest">QuestionRequest</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /question/{requestID}/reply">client.Question.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionService.Reply">Reply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionReplyParams">QuestionReplyParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /question/{requestID}/reject">client.Question.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionService.Reject">Reject</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#QuestionRejectParams">QuestionRejectParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tui

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlNextResponse">TuiControlNextResponse</a>

Methods:

- <code title="post /tui/append-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.AppendPrompt">AppendPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiAppendPromptParams">TuiAppendPromptParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/clear-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.ClearPrompt">ClearPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiClearPromptParams">TuiClearPromptParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/execute-command">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.ExecuteCommand">ExecuteCommand</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiExecuteCommandParams">TuiExecuteCommandParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-help">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.OpenHelp">OpenHelp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiOpenHelpParams">TuiOpenHelpParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-models">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.OpenModels">OpenModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiOpenModelsParams">TuiOpenModelsParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-sessions">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.OpenSessions">OpenSessions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiOpenSessionsParams">TuiOpenSessionsParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/open-themes">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.OpenThemes">OpenThemes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiOpenThemesParams">TuiOpenThemesParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/publish">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.Publish">Publish</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiPublishParams">TuiPublishParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/select-session">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.SelectSession">SelectSession</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiSelectSessionParams">TuiSelectSessionParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/show-toast">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.ShowToast">ShowToast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiShowToastParams">TuiShowToastParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/submit-prompt">client.Tui.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiService.SubmitPrompt">SubmitPrompt</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiSubmitPromptParams">TuiSubmitPromptParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Control

Methods:

- <code title="get /tui/control/next">client.Tui.Control.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlService.Next">Next</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlNextParams">TuiControlNextParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlNextResponse">TuiControlNextResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /tui/control/response">client.Tui.Control.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlService.Response">Response</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#TuiControlResponseParams">TuiControlResponseParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lsp

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#LspStatus">LspStatus</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#LspStatusStatus">LspStatusStatus</a>

Methods:

- <code title="get /lsp">client.Lsp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#LspService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#LspStatusParams">LspStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#LspStatus">LspStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Formatter

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FormatterStatus">FormatterStatus</a>

Methods:

- <code title="get /formatter">client.Formatter.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FormatterService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FormatterStatusParams">FormatterStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#FormatterStatus">FormatterStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Global

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalHealthResponse">GlobalHealthResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalEvent">GlobalEvent</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalUpgradeResponse">GlobalUpgradeResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalUpgradeResponseSuccess">GlobalUpgradeResponseSuccess</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalUpgradeResponseFailure">GlobalUpgradeResponseFailure</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalSession">GlobalSession</a>

Methods:

- <code title="get /global/health">client.Global.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalHealthResponse">GlobalHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /global/event">client.Global.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalService.Event">Event</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) *<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/packages/ssestream">ssestream</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go/packages/ssestream#Stream">Stream</a>[<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalEvent">GlobalEvent</a>]</code>
- <code title="get /global/config">client.Global.Config.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalConfigService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Config">Config</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /global/config">client.Global.Config.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalConfigService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalConfigUpdateParams">GlobalConfigUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Config">Config</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /global/dispose">client.Global.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalService.Dispose">Dispose</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /global/upgrade">client.Global.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalService.Upgrade">Upgrade</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalUpgradeParams">GlobalUpgradeParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalUpgradeResponse">GlobalUpgradeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Instance

Methods:

- <code title="post /instance/dispose">client.Instance.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#InstanceService.Dispose">Dispose</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#InstanceDisposeParams">InstanceDisposeParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Mcp

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatus">McpStatus</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusConnected">McpStatusConnected</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusDisabled">McpStatusDisabled</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusFailed">McpStatusFailed</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusNeedsAuth">McpStatusNeedsAuth</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusNeedsClientRegistration">McpStatusNeedsClientRegistration</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthStartResponse">McpAuthStartResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthRemoveResponse">McpAuthRemoveResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpResource">McpResource</a>

Methods:

- <code title="get /mcp">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatusParams">McpStatusParams</a>) (map[string]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatus">McpStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAddParams">McpAddParams</a>) (map[string]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatus">McpStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp/{name}/connect">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpConnectParams">McpConnectParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp/{name}/disconnect">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.Disconnect">Disconnect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpDisconnectParams">McpDisconnectParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp/{name}/auth/start">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.AuthStart">AuthStart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthStartParams">McpAuthStartParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthStartResponse">McpAuthStartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /mcp/{name}/auth">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.AuthRemove">AuthRemove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthRemoveParams">McpAuthRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthRemoveResponse">McpAuthRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp/{name}/auth/authenticate">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.AuthAuthenticate">AuthAuthenticate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthAuthenticateParams">McpAuthAuthenticateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatus">McpStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /mcp/{name}/auth/callback">client.Mcp.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpService.AuthCallback">AuthCallback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, name <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpAuthCallbackParams">McpAuthCallbackParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpStatus">McpStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pty

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Pty">Pty</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyStatus">PtyStatus</a>

Methods:

- <code title="get /pty">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyListParams">PtyListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Pty">Pty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pty">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.Create">Create</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyCreateParams">PtyCreateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Pty">Pty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pty/{ptyID}">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ptyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyGetParams">PtyGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Pty">Pty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pty/{ptyID}">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ptyID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyUpdateParams">PtyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Pty">Pty</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pty/{ptyID}">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ptyID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyRemoveParams">PtyRemoveParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pty/{ptyID}/connect">client.Pty.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ptyID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#PtyConnectParams">PtyConnectParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Vcs

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsInfo">VcsInfo</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsFileDiff">VcsFileDiff</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsFileDiffStatus">VcsFileDiffStatus</a>

Methods:

- <code title="get /vcs">client.Vcs.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsGetParams">VcsGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsInfo">VcsInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vcs/diff">client.Vcs.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsService.Diff">Diff</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsDiffParams">VcsDiffParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#VcsFileDiff">VcsFileDiff</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Worktree

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Worktree">Worktree</a>

Methods:

- <code title="post /worktree">client.Worktree.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeService.Create">Create</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeCreateParams">WorktreeCreateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Worktree">Worktree</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /worktree">client.Worktree.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeListParams">WorktreeListParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /worktree">client.Worktree.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeRemoveParams">WorktreeRemoveParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /worktree/reset">client.Worktree.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeService.Reset">Reset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorktreeResetParams">WorktreeResetParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workspace

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Workspace">Workspace</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceStatusResponse">WorkspaceStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceStatusResponseStatus">WorkspaceStatusResponseStatus</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceAdaptor">WorkspaceAdaptor</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceSessionRestoreResponse">WorkspaceSessionRestoreResponse</a>

Methods:

- <code title="post /experimental/workspace">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.Create">Create</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceCreateParams">WorkspaceCreateParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /experimental/workspace">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceListParams">WorkspaceListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /experimental/workspace/status">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceStatusParams">WorkspaceStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceStatusResponse">WorkspaceStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /experimental/workspace/adaptor">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.Adaptors">Adaptors</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceAdaptorsParams">WorkspaceAdaptorsParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceAdaptor">WorkspaceAdaptor</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /experimental/workspace/{id}">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceRemoveParams">WorkspaceRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /experimental/workspace/{id}/session-restore">client.Workspace.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceService.SessionRestore">SessionRestore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceSessionRestoreParams">WorkspaceSessionRestoreParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#WorkspaceSessionRestoreResponse">WorkspaceSessionRestoreResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sync

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncReplayResponse">SyncReplayResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncHistoryEvent">SyncHistoryEvent</a>

Methods:

- <code title="post /sync/start">client.Sync.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncStartParams">SyncStartParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /sync/replay">client.Sync.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncService.Replay">Replay</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncReplayParams">SyncReplayParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncReplayResponse">SyncReplayResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /sync/history">client.Sync.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncService.History">History</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncHistoryParams">SyncHistoryParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SyncHistoryEvent">SyncHistoryEvent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Skill

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SkillItem">SkillItem</a>

Methods:

- <code title="get /skill">client.Skill.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SkillService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SkillListParams">SkillListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#SkillItem">SkillItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tool

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolListItem">ToolListItem</a>

Methods:

- <code title="get /experimental/tool">client.Tool.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolListParams">ToolListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolListItem">ToolListItem</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /experimental/tool/ids">client.Tool.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolService.IDs">IDs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ToolIDsParams">ToolIDsParams</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Resource

Note: Uses `McpResource` type from the Mcp service.

Methods:

- <code title="get /experimental/resource">client.Resource.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ResourceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ResourceListParams">ResourceListParams</a>) (map[string]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#McpResource">McpResource</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Console

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleState">ConsoleState</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleListOrgsResponse">ConsoleListOrgsResponse</a>
- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleOrg">ConsoleOrg</a>

Methods:

- <code title="get /experimental/console">client.Console.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleGetParams">ConsoleGetParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleState">ConsoleState</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /experimental/console/orgs">client.Console.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleService.ListOrgs">ListOrgs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleListOrgsParams">ConsoleListOrgsParams</a>) (<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleListOrgsResponse">ConsoleListOrgsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /experimental/console/switch">client.Console.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleService.SwitchOrg">SwitchOrg</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ConsoleSwitchOrgParams">ConsoleSwitchOrgParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ExperimentalSession

Response Types:

- <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalSession">GlobalSession</a>

Methods:

- <code title="get /experimental/session">client.ExperimentalSession.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ExperimentalSessionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#ExperimentalSessionListParams">ExperimentalSessionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go">opencode</a>.<a href="https://pkg.go.dev/github.com/GunsonJack/opencode-sdk-go#GlobalSession">GlobalSession</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
