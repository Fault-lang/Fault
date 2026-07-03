package format

import (
	"embed"
	"encoding/json"
	"fault/ast"
	"fault/generator/scenario"
	"fault/runner"
	"fmt"
	"os"
	"strings"
	"text/template"
)

//go:embed templates/*.tmpl
var builtinTemplates embed.FS

// ResultData is the template-friendly view of a CompilationOutput.
// It is passed as the root data object (.) when rendering a user template.
type ResultData struct {
	// SystemName is the name declared in the spec or system file.
	SystemName string

	// Message is a top-level status string (e.g. "Fault could not find a failure case. All good!")
	// set when the solver returned unsat. Empty when results are present.
	Message string

	// Warnings are compiler warnings emitted during compilation.
	Warnings []string

	// HasResults is true when the solver returned sat and a scenario trace is available.
	HasResults bool

	// Assertions is the list of assert/assume statements with their pass/fail status.
	Assertions []AssertionResult

	// Trace is the human-readable scenario trace produced by the default formatter.
	// Useful when you want the existing trace text embedded in a larger template.
	Trace string

	// Results is the raw solver output: SSA variable name → string value.
	// Useful when writing templates that need to look up values by variable name.
	Results map[string]string

	// Events is the structured list of scenario events after dead-branch pruning.
	// Use this to build a completely custom trace representation.
	Events []TraceEvent

	// JSON is the pre-built structured result for the json format template.
	// Also available to user templates via {{toJSON .JSON}}.
	JSON *JSONResult
}

// AssertionResult represents one assert or assume statement after model checking.
type AssertionResult struct {
	// Text is the assertion body, e.g. "assert queue_size >= 0".
	Text string

	// Violated is true when the assertion was falsified by the solver.
	Violated bool

	// Kind is "assert" or "assume".
	Kind string

	// Status is "FAILED" or "OK".
	Status string
}

// TraceEvent is one event in the scenario trace.
type TraceEvent struct {
	// Kind is one of: "function_entry", "function_exit", "variable", "solvable", "message", "choice".
	Kind string

	// Function is the function name for function_entry and function_exit events.
	Function string

	// CallType is "Entry" or "Exit" for function events.
	CallType string

	// Variable is the variable name (system prefix stripped) for variable and solvable events.
	Variable string

	// Value is the solver-assigned value for variable and solvable events.
	Value string

	// Round is the simulation round string for the event.
	Round string

	// Text is the message body for message events.
	Text string

	// Dead is true for events pruned from the solution (unreachable branches, etc.).
	Dead bool
}

// JSONResult is the structured output for the json format.
type JSONResult struct {
	Sat       bool              `json:"sat"`
	Error     *string           `json:"error"`
	Warnings  []string          `json:"warnings"`
	Variables map[string]string `json:"variables"`
	Steps     map[string]string `json:"steps"`
	Asserts   []JSONAssertion   `json:"asserts"`
}

// JSONAssertion is one assert/assume entry in JSONResult.
type JSONAssertion struct {
	Kind     string `json:"kind"`
	Text     string `json:"text"`
	Violated bool   `json:"violated"`
}

// BuildJSON converts a CompilationOutput into a JSONResult.
func BuildJSON(output *runner.CompilationOutput) *JSONResult {
	r := &JSONResult{
		Warnings:  output.Warnings,
		Variables: map[string]string{},
		Steps:     map[string]string{},
		Asserts:   []JSONAssertion{},
	}

	if output.Error != nil {
		msg := output.Error.Error()
		r.Error = &msg
		return r
	}

	if output.ResultLog == nil {
		// unsat — no scenario found
		r.Sat = false
		sysPrefix := ""
		for _, a := range output.Asserts {
			r.Asserts = append(r.Asserts, buildJSONAssertion(a, sysPrefix))
		}
		return r
	}

	r.Sat = true
	log := output.ResultLog
	sysPrefix := log.SystemName + "_"

	// Clean variables: strip SSA suffix and system prefix, keep last value per base name.
	seen := map[string]bool{}
	for _, e := range log.Events {
		var rawVar string
		switch ev := e.(type) {
		case *scenario.VariableUpdate:
			if e.IsDead() || log.IsInternalVariable(ev.Variable) {
				continue
			}
			rawVar = ev.Variable
		case *scenario.Solvable:
			if e.IsDead() || log.IsInternalVariable(ev.Variable) {
				continue
			}
			rawVar = ev.Variable
		default:
			continue
		}
		base := getBase(rawVar)
		clean := strings.TrimPrefix(base, sysPrefix)
		if !seen[clean] {
			seen[clean] = true
		}
		r.Variables[clean] = log.Results[rawVar]
	}

	// Synthesis steps: collect synth_N entry events in order.
	for _, e := range log.Events {
		fc, ok := e.(*scenario.FunctionCall)
		if !ok || fc.Type != "Entry" || !strings.HasPrefix(fc.FunctionName, "synth_") {
			continue
		}
		chosen := synthChoice(log, fc.FunctionName)
		step := strings.TrimPrefix(fc.FunctionName, "synth_")
		step = strings.TrimPrefix(step, sysPrefix)
		if chosen != "" {
			r.Steps[step] = chosen
		}
	}

	for _, a := range output.Asserts {
		r.Asserts = append(r.Asserts, buildJSONAssertion(a, sysPrefix))
	}

	return r
}

func buildJSONAssertion(a *ast.AssertionStatement, sysPrefix string) JSONAssertion {
	ar := buildAssertion(a, sysPrefix)
	return JSONAssertion{Kind: ar.Kind, Text: ar.Text, Violated: a.Violated}
}

// synthChoice replicates scenario.Logger.synthChoice for use in this package.
func synthChoice(log *scenario.Logger, slotName string) string {
	prefix := slotName + "_"
	for varWithSSA, val := range log.Results {
		if val != "true" {
			continue
		}
		base := getBase(varWithSSA)
		if strings.HasPrefix(base, prefix) {
			return strings.TrimPrefix(base, prefix)
		}
	}
	return ""
}

// getBase strips the SSA numeric suffix (_N) from a variable name.
func getBase(varWithSSA string) string {
	idx := strings.LastIndex(varWithSSA, "_")
	if idx < 0 {
		return varWithSSA
	}
	suffix := varWithSSA[idx+1:]
	for _, c := range suffix {
		if c < '0' || c > '9' {
			return varWithSSA
		}
	}
	return varWithSSA[:idx]
}

// Build converts a CompilationOutput into a ResultData ready for template rendering.
// sysPrefix is used to strip the system name from assertion text; pass an empty string
// to leave names unmodified.
func Build(output *runner.CompilationOutput) *ResultData {
	d := &ResultData{
		Warnings:   output.Warnings,
		Message:    output.Message,
		HasResults: output.ResultLog != nil,
		Results:    make(map[string]string),
	}

	if output.ResultLog != nil {
		d.SystemName = output.ResultLog.SystemName
		d.Trace = output.ResultLog.String()
		d.Results = output.ResultLog.Results
		d.Events = buildEvents(output.ResultLog)
	}

	sysPrefix := d.SystemName + "_"
	for _, a := range output.Asserts {
		d.Assertions = append(d.Assertions, buildAssertion(a, sysPrefix))
	}

	d.JSON = BuildJSON(output)

	return d
}

func buildAssertion(a *ast.AssertionStatement, sysPrefix string) AssertionResult {
	ar := AssertionResult{Violated: a.Violated}
	if a.Violated {
		ar.Status = "FAILED"
	} else {
		ar.Status = "OK"
	}
	if a.Assume {
		ar.Kind = "assume"
	} else {
		ar.Kind = "assert"
	}

	full := a.EvLogString(true)
	// Strip system name prefix from variable names in the assertion text.
	if sysPrefix != "_" {
		full = strings.ReplaceAll(full, sysPrefix, "")
	}
	// Strip the "OK  " or "FAILED  " status prefix to get just the assertion text.
	text := full
	if strings.HasPrefix(text, "FAILED  ") {
		text = text[8:]
	} else if strings.HasPrefix(text, "OK  ") {
		text = text[4:]
	}
	ar.Text = strings.TrimSuffix(text, ";")
	return ar
}

func buildEvents(log *scenario.Logger) []TraceEvent {
	sysPrefix := log.SystemName + "_"
	var events []TraceEvent
	for _, e := range log.Events {
		switch ev := e.(type) {
		case *scenario.FunctionCall:
			kind := "function_entry"
			if ev.Type == "Exit" {
				kind = "function_exit"
			}
			name := ev.FunctionName
			if sysPrefix != "_" {
				name = strings.TrimPrefix(name, sysPrefix)
			}
			events = append(events, TraceEvent{
				Kind:     kind,
				Function: name,
				CallType: ev.Type,
				Round:    ev.Round,
				Dead:     ev.IsDead(),
			})
		case *scenario.VariableUpdate:
			varName := ev.Variable
			if sysPrefix != "_" {
				varName = strings.TrimPrefix(varName, sysPrefix)
			}
			events = append(events, TraceEvent{
				Kind:     "variable",
				Variable: varName,
				Value:    log.Results[ev.Variable],
				Round:    ev.Round,
				Dead:     ev.IsDead(),
			})
		case *scenario.Solvable:
			varName := ev.Variable
			if sysPrefix != "_" {
				varName = strings.TrimPrefix(varName, sysPrefix)
			}
			events = append(events, TraceEvent{
				Kind:     "solvable",
				Variable: varName,
				Value:    log.Results[ev.Variable],
				Round:    ev.Round,
				Dead:     ev.IsDead(),
			})
		case *scenario.Message:
			events = append(events, TraceEvent{
				Kind:  "message",
				Text:  ev.Text,
				Round: ev.Round,
				Dead:  ev.IsDead(),
			})
		case *scenario.Choice:
			events = append(events, TraceEvent{
				Kind: "choice",
				Dead: ev.IsDead(),
			})
		}
	}
	return events
}

// RenderTemplate renders a Go template file against data and returns the result.
func RenderTemplate(data *ResultData, tmplPath string) (string, error) {
	tmplBytes, err := os.ReadFile(tmplPath)
	if err != nil {
		return "", fmt.Errorf("reading format template %q: %w", tmplPath, err)
	}
	return RenderTemplateString(data, string(tmplBytes))
}

// BuiltinNames lists the available built-in format templates.
var BuiltinNames = []string{"default", "json", "compact"}

// RenderBuiltin renders one of the built-in templates by name.
// Valid names are those in BuiltinNames.
func RenderBuiltin(data *ResultData, name string) (string, error) {
	b, err := builtinTemplates.ReadFile("templates/" + name + ".tmpl")
	if err != nil {
		return "", fmt.Errorf("unknown built-in format %q (available: default, json, compact)", name)
	}
	return RenderTemplateString(data, string(b))
}

// RenderTemplateString renders an inline Go template string against data.
func RenderTemplateString(data *ResultData, tmpl string) (string, error) {
	t, err := template.New("output").Funcs(templateFuncs()).Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("parsing format template: %w", err)
	}
	var buf strings.Builder
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing format template: %w", err)
	}
	return buf.String(), nil
}

// templateFuncs returns the helper functions available inside user templates.
func templateFuncs() template.FuncMap {
	return template.FuncMap{
		// toJSON marshals v to indented JSON.
		// Example: {{toJSON .}} or {{toJSON .Assertions}}
		"toJSON": func(v any) (string, error) {
			b, err := json.MarshalIndent(v, "", "  ")
			return string(b), err
		},
		// toJSONCompact marshals v to compact (single-line) JSON.
		"toJSONCompact": func(v any) (string, error) {
			b, err := json.Marshal(v)
			return string(b), err
		},
		// join concatenates elems with sep (mirrors strings.Join).
		// Example: {{join .Warnings "\n"}}
		"join": strings.Join,
		// upper converts s to upper case.
		"upper": strings.ToUpper,
		// lower converts s to lower case.
		"lower": strings.ToLower,
		// trimPrefix removes prefix from s.
		"trimPrefix": strings.TrimPrefix,
		// trimSuffix removes suffix from s.
		"trimSuffix": strings.TrimSuffix,
		// contains reports whether substr is in s.
		"contains": strings.Contains,
		// replace replaces all occurrences of old with new in s.
		"replace": strings.ReplaceAll,
		// default returns val if non-empty, otherwise def.
		// Example: {{default "none" .Message}}
		"default": func(def, val string) string {
			if val == "" {
				return def
			}
			return val
		},
	}
}
