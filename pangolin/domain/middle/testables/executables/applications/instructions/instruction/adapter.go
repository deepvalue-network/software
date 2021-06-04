package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/registry"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/value"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type adapter struct {
	stackFrameBuilder       stackframe.Builder
	skipBuilder             stackframe.SkipBuilder
	saveBuilder             stackframe.SaveBuilder
	conditionBuilder        condition.Builder
	propositionBuilder      condition.PropositionBuilder
	remainingBuilder        remaining.Builder
	standardBuilder         standard.Builder
	valueBuilder            value.Builder
	varValueAdapter         var_value.Adapter
	varValueFactory         var_value.Factory
	varVariableBuilder      var_variable.Builder
	callBuilder             call.Builder
	exitBuilder             exit.Builder
	registryIndexBuilder    registry.IndexBuilder
	registryRegisterBuilder registry.RegisterBuilder
	registerFetchBuilder    registry.FetchBuilder
	registryBuilder         registry.Builder
	builder                 Builder
}

func createAdapter(
	stackFrameBuilder stackframe.Builder,
	skipBuilder stackframe.SkipBuilder,
	saveBuilder stackframe.SaveBuilder,
	conditionBuilder condition.Builder,
	propositionBuilder condition.PropositionBuilder,
	remainingBuilder remaining.Builder,
	standardBuilder standard.Builder,
	valueBuilder value.Builder,
	varValueAdapter var_value.Adapter,
	varValueFactory var_value.Factory,
	varVariableBuilder var_variable.Builder,
	callBuilder call.Builder,
	exitBuilder exit.Builder,
	registryIndexBuilder registry.IndexBuilder,
	registryRegisterBuilder registry.RegisterBuilder,
	registerFetchBuilder registry.FetchBuilder,
	registryBuilder registry.Builder,
	builder Builder,
) Adapter {
	out := adapter{
		stackFrameBuilder:       stackFrameBuilder,
		skipBuilder:             skipBuilder,
		saveBuilder:             saveBuilder,
		conditionBuilder:        conditionBuilder,
		propositionBuilder:      propositionBuilder,
		remainingBuilder:        remainingBuilder,
		standardBuilder:         standardBuilder,
		valueBuilder:            valueBuilder,
		varValueAdapter:         varValueAdapter,
		varValueFactory:         varValueFactory,
		varVariableBuilder:      varVariableBuilder,
		callBuilder:             callBuilder,
		exitBuilder:             exitBuilder,
		registryIndexBuilder:    registryIndexBuilder,
		registryRegisterBuilder: registryRegisterBuilder,
		registerFetchBuilder:    registerFetchBuilder,
		registryBuilder:         registryBuilder,
		builder:                 builder,
	}

	return &out
}

// ToInstruction converts a parsed instruction to an optimized instruction
func (app *adapter) ToInstruction(parsed parsers.Instruction) (Instruction, error) {
	builder := app.builder.Create()

	if parsed.IsPrint() {
		parsedVal := parsed.Print().Value()
		varValue, err := app.varValueAdapter.ToValue(parsedVal)
		if err != nil {
			return nil, err
		}

		val, err := app.valueBuilder.Create().WithValue(varValue).IsPrint().Now()
		if err != nil {
			return nil, err
		}

		builder.WithValue(val)
	}

	if parsed.IsVariable() {
		vr := parsed.Variable()
		if vr.IsAssignment() {
			ass := vr.Assignment()
			parsedVal := ass.Value()
			val, err := app.varValueAdapter.ToValue(parsedVal)
			if err != nil {
				return nil, err
			}

			parsedVariableName := ass.Variable()
			ins, err := app.varVariableBuilder.Create().WithValue(val).WithName(parsedVariableName).Now()
			if err != nil {
				return nil, err
			}

			builder.WithSave(ins)
		}

		if vr.IsDeclaration() {
			decl := vr.Declaration()
			ins, err := app.insert(decl)
			if err != nil {
				return nil, err
			}

			builder.WithInsert(ins)
		}

		if vr.IsConcatenation() {
			concat := vr.Concatenation().Operation()
			st, err := app.standard(concat).IsConcatenation().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStandard(st)
		}

		if vr.IsDelete() {
			del := vr.Delete()
			builder.WithDelete(del)
		}
	}

	if parsed.IsOperation() {
		op := parsed.Operation()
		if op.IsArythmetic() {
			ary := op.Arythmetic()
			if ary.IsAdd() {
				add := ary.Add()
				st, err := app.standard(add).IsAdd().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if ary.IsDiv() {
				div := ary.Div()
				rem, err := app.remaining(div).IsDiv().Now()
				if err != nil {
					return nil, err
				}

				builder.WithRemaining(rem)
			}

			if ary.IsMul() {
				mul := ary.Mul()
				st, err := app.standard(mul).IsMul().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if ary.IsSub() {
				sub := ary.Sub()
				st, err := app.standard(sub).IsSub().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}

		if op.IsRelational() {
			rel := op.Relational()
			if rel.IsEqual() {
				eq := rel.Equal()
				st, err := app.standard(eq).IsEqual().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if rel.IsLessThan() {
				lessThan := rel.LessThan()
				st, err := app.standard(lessThan).IsLessThan().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if rel.IsNotEqual() {
				notEqual := rel.NotEqual()
				st, err := app.standard(notEqual).IsNotEqual().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}

		if op.IsLogical() {
			logical := op.Logical()
			if logical.IsAnd() {
				and := logical.And()
				st, err := app.standard(and).IsAnd().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}

			if logical.IsOr() {
				or := logical.Or()
				st, err := app.standard(or).IsOr().Now()
				if err != nil {
					return nil, err
				}

				builder.WithStandard(st)
			}
		}
	}

	if parsed.IsStackFrame() {
		stackFrame := parsed.StackFrame()
		if stackFrame.IsPush() {
			stackframe, err := app.stackFrameBuilder.Create().IsPush().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsPop() {
			stackframe, err := app.stackFrameBuilder.Create().IsPop().Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsSkip() {
			skipBuilder := app.skipBuilder.Create()
			ptr := stackFrame.Skip().Pointer()
			if ptr.IsInt() {
				intVal := ptr.Int()
				skipBuilder.WithInt(intVal)
			}

			if ptr.IsVariable() {
				variable := ptr.Variable()
				skipBuilder.WithVariable(variable)
			}

			skip, err := skipBuilder.Now()
			if err != nil {
				return nil, err
			}

			stackframe, err := app.stackFrameBuilder.Create().WithSkip(skip).Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}

		if stackFrame.IsIndex() {
			variable := stackFrame.Index().Variable()
			stackframe, err := app.stackFrameBuilder.Create().WithIndex(variable).Now()
			if err != nil {
				return nil, err
			}

			builder.WithStackframe(stackframe)
		}
	}

	if parsed.IsJump() {
		jmp := parsed.Jump()
		condition, err := app.conditionFromJump(jmp)
		if err != nil {
			return nil, err
		}

		builder.WithCondition(condition)
	}

	if parsed.IsExit() {
		parsedExit := parsed.Exit()
		exit, err := app.exit(parsedExit)
		if err != nil {
			return nil, err
		}

		builder.WithExit(exit)
	}

	if parsed.IsCall() {
		parsedCall := parsed.Call()
		call, err := app.call(parsedCall)
		if err != nil {
			return nil, err
		}

		builder.WithCall(call)
	}

	if parsed.IsRegistry() {
		parsedRegistry := parsed.Registry()
		registry, err := app.registry(parsedRegistry)
		if err != nil {
			return nil, err
		}

		builder.WithRegistry(registry)
	}

	if parsed.IsSave() {
		save := parsed.Save()
		to := save.To()
		saveBuilder := app.saveBuilder.Create().To(to)
		if save.HasFrom() {
			from := save.From()
			saveBuilder.From(from)
		}

		save, err := saveBuilder.Now()
		if err != nil {
			return nil, err
		}

		stackFrame, err := app.stackFrameBuilder.Create().WithSave(save).Now()
		if err != nil {
			return nil, err
		}

		builder.WithStackframe(stackFrame)
	}

	if parsed.IsSwitch() {
		variable := parsed.Switch().Variable()
		stackFrame, err := app.stackFrameBuilder.Create().WithSwitch(variable).Now()
		if err != nil {
			return nil, err
		}

		builder.WithStackframe(stackFrame)
	}

	return builder.Now()
}

func (app *adapter) insert(parsed parsers.Declaration) (var_variable.Variable, error) {
	typ := parsed.Type()
	name := parsed.Variable()
	val, err := app.varValueFactory.Create(typ)
	if err != nil {
		return nil, err
	}

	return app.varVariableBuilder.Create().WithName(name).WithValue(val).Now()
}

func (app *adapter) registry(parsed parsers.Registry) (registry.Registry, error) {
	builder := app.registryBuilder.Create()
	if parsed.IsFetch() {
		parsedFetch := parsed.Fetch()
		to := parsedFetch.To()
		from := parsedFetch.From()
		fetchBuilder := app.registerFetchBuilder.Create().From(from).To(to)
		if parsedFetch.HasIndex() {
			parsedIndex := parsedFetch.Index()
			index, err := app.registryIndex(parsedIndex)
			if err != nil {
				return nil, err
			}

			fetchBuilder.WithIndex(index)
		}

		fetch, err := fetchBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithFetch(fetch)
	}

	if parsed.IsRegister() {
		parsedRegister := parsed.Register()
		vr := parsedRegister.Variable()
		registerBuilder := app.registryRegisterBuilder.Create().WithVariable(vr)
		if parsedRegister.HasIndex() {
			parsedIndex := parsedRegister.Index()
			index, err := app.registryIndex(parsedIndex)
			if err != nil {
				return nil, err
			}

			registerBuilder.WithIndex(index)
		}

		register, err := registerBuilder.Now()
		if err != nil {
			return nil, err
		}

		builder.WithRegister(register)
	}

	if parsed.IsUnregister() {
		unregister := parsed.Unregister().Variable()
		builder.WithUnregister(unregister)
	}

	return builder.Now()
}

func (app *adapter) registryIndex(parsed parsers.IntPointer) (registry.Index, error) {
	builder := app.registryIndexBuilder.Create()
	if parsed.IsInt() {
		intVal := parsed.Int()
		builder.WithInt(intVal)
	}

	if parsed.IsVariable() {
		vr := parsed.Variable()
		builder.WithVariable(vr)
	}

	return builder.Now()
}

func (app *adapter) call(parsed parsers.Call) (call.Call, error) {
	name := parsed.Name()
	stackFrame := parsed.StackFrame()
	builder := app.callBuilder.Create().WithName(name).WithStackFrame(stackFrame)
	if parsed.HasCondition() {
		condition := parsed.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *adapter) exit(parsed parsers.Exit) (exit.Exit, error) {
	builder := app.exitBuilder.Create()
	if parsed.HasCondition() {
		condition := parsed.Condition()
		builder.WithCondition(condition)
	}

	return builder.Now()
}

func (app *adapter) conditionFromJump(parsed parsers.Jump) (condition.Condition, error) {
	label := parsed.Label()
	propositionBuilder := app.propositionBuilder.Create().WithName(label)
	if parsed.HasCondition() {
		condition := parsed.Condition()
		propositionBuilder.WithCondition(condition)
	}

	proposition, err := propositionBuilder.Now()
	if err != nil {
		return nil, err
	}

	return app.conditionBuilder.Create().IsJump().WithProposition(proposition).Now()
}

func (app *adapter) remaining(parsed parsers.RemainingOperation) remaining.Builder {
	builder := app.remainingBuilder.Create()
	first := parsed.First()
	builder.WithFirst(first)

	second := parsed.Second()
	builder.WithSecond(second)

	result := parsed.Result()
	builder.WithResult(result)

	remaining := parsed.Remaining()
	builder.WithRemaining(remaining)

	return builder
}

func (app *adapter) standard(parsed parsers.StandardOperation) standard.Builder {
	builder := app.standardBuilder.Create()
	first := parsed.First()
	builder.WithFirst(first)

	second := parsed.Second()
	builder.WithSecond(second)

	result := parsed.Result()
	builder.WithResult(result)

	return builder
}
