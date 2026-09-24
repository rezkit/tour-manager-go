# Examples

Small, runnable programs demonstrating this client's patterns. Each is a
standalone `package main` under its own directory — run one directly with:

```sh
REZKIT_API_KEY=... go run ./examples/<name>
```

| Example | Demonstrates |
|---|---|
| [`list-holidays`](list-holidays/main.go) | `Client.Holidays().All` auto-pagination (`iter.Seq2`), `text/tabwriter` for tabular output, bounding calls with a `context` timeout |

This suite is expected to grow alongside the client's resource coverage —
see the repository `README.md`'s "Known gaps" section and
`.claude/skills/update-tour-manager-go-client/SKILL.md` for what's
implemented so far.
