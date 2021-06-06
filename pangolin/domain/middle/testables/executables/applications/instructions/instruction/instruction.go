package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/module"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/registry"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable"
)

type instruction struct {
	stackframe stackframe.Stackframe
	condition  condition.Condition
	standard   standard.Standard
	remaining  remaining.Remaining
	value      value.Value
	insert     variable.Variable
	save       variable.Variable
	del        string
	call       call.Call
	module     module.Module
	exit       exit.Exit
	reg        registry.Registry
}

func createInstructionWithStackframe(stackframe stackframe.Stackframe) Instruction {
	return createInstructionInternally(stackframe, nil, nil, nil, nil, nil, nil, "", nil, nil, nil, nil)
}

func createInstructionWithCondition(condition condition.Condition) Instruction {
	return createInstructionInternally(nil, condition, nil, nil, nil, nil, nil, "", nil, nil, nil, nil)
}

func createInstructionWithStandard(standard standard.Standard) Instruction {
	return createInstructionInternally(nil, nil, standard, nil, nil, nil, nil, "", nil, nil, nil, nil)
}

func createInstructionWithRemaining(remaining remaining.Remaining) Instruction {
	return createInstructionInternally(nil, nil, nil, remaining, nil, nil, nil, "", nil, nil, nil, nil)
}

func createInstructionWithValue(value value.Value) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, value, nil, nil, "", nil, nil, nil, nil)
}

func createInstructionWithInsert(insert variable.Variable) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, insert, nil, "", nil, nil, nil, nil)
}

func createInstructionWithSave(save variable.Variable) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, save, "", nil, nil, nil, nil)
}

func createInstructionWithDelete(del string) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, del, nil, nil, nil, nil)
}

func createInstructionWithCall(call call.Call) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, "", call, nil, nil, nil)
}

func createInstructionWithModule(module module.Module) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, "", nil, module, nil, nil)
}

func createInstructionWithExit(exit exit.Exit) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, "", nil, nil, exit, nil)
}

func createInstructionWithRegistry(reg registry.Registry) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, "", nil, nil, nil, reg)
}

func createInstructionInternally(
	stackframe stackframe.Stackframe,
	condition condition.Condition,
	standard standard.Standard,
	remaining remaining.Remaining,
	value value.Value,
	insert variable.Variable,
	save variable.Variable,
	del string,
	call call.Call,
	module module.Module,
	exit exit.Exit,
	reg registry.Registry,
) Instruction {
	out := instruction{
		stackframe: stackframe,
		condition:  condition,
		standard:   standard,
		remaining:  remaining,
		value:      value,
		insert:     insert,
		save:       save,
		del:        del,
		call:       call,
		module:     module,
		exit:       exit,
		reg:        reg,
	}

	return &out
}

// IsStackframe returns true if there is a stackframe, false otherwise
func (obj *instruction) IsStackframe() bool {
	return obj.stackframe != nil
}

// Stackframe returns the stackframe, if any
func (obj *instruction) Stackframe() stackframe.Stackframe {
	return obj.stackframe
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *instruction) Condition() condition.Condition {
	return obj.condition
}

// IsStandard returns true if there is a standard, false otherwise
func (obj *instruction) IsStandard() bool {
	return obj.standard != nil
}

// Standard returns the standard, if any
func (obj *instruction) Standard() standard.Standard {
	return obj.standard
}

// IsRemaining returns true if there is a remaining, false otherwise
func (obj *instruction) IsRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining, if any
func (obj *instruction) Remaining() remaining.Remaining {
	return obj.remaining
}

// IsValue returns true if there is a value, false otherwise
func (obj *instruction) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *instruction) Value() value.Value {
	return obj.value
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *instruction) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *instruction) Insert() variable.Variable {
	return obj.insert
}

// IsSave returns true if there is a save, false otherwise
func (obj *instruction) IsSave() bool {
	return obj.save != nil
}

// Save returns the save, if any
func (obj *instruction) Save() variable.Variable {
	return obj.save
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *instruction) IsDelete() bool {
	return obj.del != ""
}

// Delete returns the delete, if any
func (obj *instruction) Delete() string {
	return obj.del
}

// IsCall returns true if there is a call, false otherwise
func (obj *instruction) IsCall() bool {
	return obj.call != nil
}

// Call returns the call, if any
func (obj *instruction) Call() call.Call {
	return obj.call
}

// IsModule returns true if there is a module, false otherwise
func (obj *instruction) IsModule() bool {
	return obj.module != nil
}

// Module returns the module, if any
func (obj *instruction) Module() module.Module {
	return obj.module
}

// IsExit returns true if there is an exit, false otherwise
func (obj *instruction) IsExit() bool {
	return obj.exit != nil
}

// Exit returns the exit, if any
func (obj *instruction) Exit() exit.Exit {
	return obj.exit
}

// IsRegistry returns true if there is a registry, false otherwise
func (obj *instruction) IsRegistry() bool {
	return obj.reg != nil
}

// Registry returns the registry, if any
func (obj *instruction) Registry() registry.Registry {
	return obj.reg
}
