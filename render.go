package main

import (
	"encoding/json"
	"fault/execute"
	"fmt"
	"os"
	ospath "path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func newRenderCmd() *cobra.Command {
	var paramsFile string
	var paramOverrides []string

	cmd := &cobra.Command{
		Use:   "render",
		Short: "Render a .smt2.tmpl template and run model checking",
		Long: `Render substitutes param values into a .smt2.tmpl template produced by
'fault -m template', then sends the result to the solver.

Param values are loaded from the accompanying .params.json file by default.
Override individual values with --param name=value.

Example:
  fault -m template -f=model.fspec
  fault render -f=model.smt2.tmpl --param threshold=0.8`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tmplFile, _ := cmd.Flags().GetString("file")
			if tmplFile == "" {
				return fmt.Errorf("required flag \"file\" not set")
			}
			return runRender(tmplFile, paramsFile, paramOverrides)
		},
	}

	cmd.Flags().StringP("file", "f", "", "path to .smt2.tmpl file")
	cmd.Flags().StringVar(&paramsFile, "params", "", "path to .params.json file (default: <template>.params.json)")
	cmd.Flags().StringArrayVar(&paramOverrides, "param", nil, "override a param value: name=value (repeatable)")

	return cmd
}

func runRender(tmplFile, paramsFile string, overrides []string) error {
	// Read template
	tmplData, err := os.ReadFile(tmplFile)
	if err != nil {
		return fmt.Errorf("could not read template: %w", err)
	}

	// Locate params file
	if paramsFile == "" {
		paramsFile = strings.TrimSuffix(tmplFile, ".smt2.tmpl") + ".params.json"
	}

	manifest, err := loadParamsManifest(paramsFile)
	if err != nil {
		return err
	}

	// Parse overrides
	values, err := parseParamOverrides(overrides, manifest)
	if err != nil {
		return err
	}

	// Check all params are provided
	var missing []string
	for token := range manifest {
		name := tokenToName(token)
		if _, ok := values[name]; !ok {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing values for params: %s\nProvide them with --param name=value", strings.Join(missing, ", "))
	}

	// Substitute tokens
	smt := string(tmplData)
	for token, sort := range manifest {
		name := tokenToName(token)
		val := values[name]
		if err := validateParamValue(val, sort); err != nil {
			return fmt.Errorf("param %s: %w", name, err)
		}
		smt = strings.ReplaceAll(smt, token, val)
	}

	// Run solver
	mc, err := execute.NewModelChecker()
	if err != nil {
		return fmt.Errorf("solver not configured: %w", err)
	}
	mc.LoadModel(smt, nil, nil)

	ok, err := mc.Check()
	if err != nil {
		return fmt.Errorf("model checker failed: %w", err)
	}
	if !ok {
		fmt.Println("Fault could not find a failure case.")
		return nil
	}

	if err := mc.Solve(); err != nil {
		return fmt.Errorf("error fetching solution from solver: %w", err)
	}

	if mc.Log != nil {
		mc.Log.Print()
	}

	return nil
}

// loadParamsManifest reads the .params.json file produced by template mode.
// Returns a map of __PARAM_name__ token → SMT sort (Real, Int, Bool).
func loadParamsManifest(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("params file not found: %s\nGenerate it with: fault -m template -f=<spec>", ospath.Base(path))
		}
		return nil, fmt.Errorf("could not read params file: %w", err)
	}
	var manifest map[string]string
	if err := json.Unmarshal(data, &manifest); err != nil {
		return nil, fmt.Errorf("could not parse params file: %w", err)
	}
	return manifest, nil
}

// parseParamOverrides parses --param name=value flags and merges them with
// any defaults that can be inferred. Unrecognised names are rejected early.
func parseParamOverrides(overrides []string, manifest map[string]string) (map[string]string, error) {
	// Build a name→token reverse index so we can validate names
	nameToToken := make(map[string]string, len(manifest))
	for token := range manifest {
		nameToToken[tokenToName(token)] = token
	}

	values := make(map[string]string)
	for _, o := range overrides {
		name, val, ok := strings.Cut(o, "=")
		if !ok {
			return nil, fmt.Errorf("invalid --param format %q: expected name=value", o)
		}
		if _, known := nameToToken[name]; !known {
			return nil, fmt.Errorf("unknown param %q — check %v for valid names", name, keys(nameToToken))
		}
		values[name] = val
	}
	return values, nil
}

// validateParamValue does a light sanity check that the provided value is
// compatible with the SMT sort declared in the manifest.
func validateParamValue(val, sort string) error {
	switch sort {
	case "Bool":
		if val != "true" && val != "false" {
			return fmt.Errorf("expected true or false for Bool param, got %q", val)
		}
	case "Int":
		if strings.ContainsAny(val, ".") {
			return fmt.Errorf("expected integer for Int param, got %q", val)
		}
	}
	return nil
}

// tokenToName strips the __PARAM_ prefix and trailing __ from a token.
func tokenToName(token string) string {
	name := strings.TrimPrefix(token, "__PARAM_")
	name = strings.TrimSuffix(name, "__")
	return name
}

func keys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
