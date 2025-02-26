package scenario

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Errorf("NewLogger() = %v, want non-nil logger", logger)
	}
	if logger.Events == nil {
		t.Errorf("NewLogger().Events = %v, want non-nil Events", logger.Events)
	}
	if logger.BranchIndexes == nil {
		t.Errorf("NewLogger().BranchIndexes = %v, want non-nil BranchIndexes", logger.BranchIndexes)
	}
	if logger.BranchVars == nil {
		t.Errorf("NewLogger().BranchVars = %v, want non-nil BranchVars", logger.BranchVars)
	}
	if logger.Forks == nil {
		t.Errorf("NewLogger().Forks = %v, want non-nil Forks", logger.Forks)
	}
	if logger.Results == nil {
		t.Errorf("NewLogger().Results = %v, want non-nil Results", logger.Results)
	}
}

func TestLogger_EnterFunction(t *testing.T) {
	logger := NewLogger()
	logger.EnterFunction("foo", 1)
	if len(logger.Events) != 1 {
		t.Errorf("Logger.EnterFunction() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*FunctionCall).FunctionName != "foo" {
		t.Errorf("Logger.EnterFunction().FunctionName = %v, want %v", logger.Events[0].(*FunctionCall).FunctionName, "foo")
	}
	if logger.Events[0].(*FunctionCall).Round != 1 {
		t.Errorf("Logger.EnterFunction().Round = %v, want %v", logger.Events[0].(*FunctionCall).Round, 1)
	}
	if logger.Events[0].(*FunctionCall).Type != "Entry" {
		t.Errorf("Logger.EnterFunction().Type = %v, want %v", logger.Events[0].(*FunctionCall).Type, "Entry")
	}
}

func TestLogger_ExitFunction(t *testing.T) {
	logger := NewLogger()
	logger.ExitFunction("foo", 1)
	if len(logger.Events) != 1 {
		t.Errorf("Logger.ExitFunction() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*FunctionCall).FunctionName != "foo" {
		t.Errorf("Logger.ExitFunction().FunctionName = %v, want %v", logger.Events[0].(*FunctionCall).FunctionName, "foo")
	}
	if logger.Events[0].(*FunctionCall).Round != 1 {
		t.Errorf("Logger.ExitFunction().Round = %v, want %v", logger.Events[0].(*FunctionCall).Round, 1)
	}
	if logger.Events[0].(*FunctionCall).Type != "Exit" {
		t.Errorf("Logger.ExitFunction().Type = %v, want %v", logger.Events[0].(*FunctionCall).Type, "Exit")
	}
}

func TestLogger_UpdateVariable(t *testing.T) {
	logger := NewLogger()
	logger.UpdateVariable("foo")
	if len(logger.Events) != 1 {
		t.Errorf("Logger.UpdateVariable() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*VariableUpdate).Variable != "foo" {
		t.Errorf("Logger.UpdateVariable().Variable = %v, want %v", logger.Events[0].(*VariableUpdate).Variable, "foo")
	}
}

func TestLogger_AddPhiOption(t *testing.T) {
	logger := NewLogger()
	logger.AddPhiOption("foo", "bar")
	if len(logger.Forks["foo"]) != 2 {
		t.Errorf("Logger.AddPhiOption() = %v, want %v", len(logger.Forks["foo"]), 2)
	}
	if logger.Forks["foo"][1] != "bar" {
		t.Errorf("Logger.AddPhiOption() = %v, want %v", logger.Forks["foo"][1], "bar")
	}
}

func TestFunctionCall_MarkDead(t *testing.T) {
	logger := NewLogger()
	logger.EnterFunction("test1", 1)
	logger.UpdateVariable("a_1")
	logger.ExitFunction("test1", 1)
	logger.EnterFunction("test2", 1)
	logger.UpdateVariable("a_2")
	logger.ExitFunction("test2", 1)
	logger.AddPhiOption("a_3", "a_1")
	logger.AddPhiOption("a_3", "a_2")

	logger.Results["a_1"] = "1"
	logger.Results["a_2"] = "5"
	logger.Results["a_3"] = "1"

	logger.Trace()
	logger.Kill()

	if logger.Events[1].IsDead() {
		t.Errorf("FunctionCall.IsDead() = %v, want %v", logger.Events[1].IsDead(), false)
	}
	if !logger.Events[4].IsDead() {
		t.Errorf("FunctionCall.IsDead() = %v, want %v", logger.Events[4].IsDead(), true)
	}
}
