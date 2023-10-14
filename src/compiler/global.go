package compiler

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

type (
	StringGlobal struct {
		glob *ir.Global
		typ  *types.ArrayType
	}
)

func (g *generator) genStrings() {
	for i, str := range g.program.Strings {
		val := str.Value
		if _, ok := g.strs[val]; ok {
			continue
		}

		arr := constant.NewCharArrayFromString(val)
		strGlob := g.mod.NewGlobalDef(".str"+fmt.Sprint(i), arr)
		strGlob.Linkage = enum.LinkagePrivate
		strGlob.UnnamedAddr = enum.UnnamedAddrUnnamedAddr
		strGlob.Immutable = true
		strGlob.Align = 1

		g.strs[val] = StringGlobal{
			strGlob,
			arr.Typ,
		}
	}
}

func (g *generator) genFuncs() {
	for i, fun := range g.program.Functions {
		name := fun.Module + "." + fun.Name

		params := []*ir.Param{}
		for i, param := range fun.Params {
			p := ir.NewParam("", g.lltyp(param.Type))
			fun.Params[i].Ir = p
			params = append(params, p)
		}

		fun.Ir = g.mod.NewFunc(
			name,
			g.lltyp(fun.Return),
			params...,
		)

		g.program.Functions[i] = fun
		g.builtins.funcs[fun.Name] = &g.program.Functions[i]
	}
}

func (g *generator) genBinOps() {
	for i, binop := range g.program.BinaryOps {
		name := binop.Module + "." + binop.Op.OperatorName() + "." + binop.Left.String() + "_" + binop.Right.String()
		if g.complex(binop.Return) {
			binop.Ir = g.mod.NewFunc(
				name,
				g.lltyp(binop.Return),
				ir.NewParam("", g.lltyp(binop.Left)),
				ir.NewParam("", g.lltyp(binop.Right)),
			)
			binop.Complex = true
		}

		hash := binop.Op.OperatorName() + " " + binop.Left.String() + " " + binop.Right.String()
		g.program.BinaryOps[i] = binop
		g.builtins.binops[hash] = &g.program.BinaryOps[i]
	}
}

func (g *generator) genUnOps() {
	for i, unop := range g.program.UnaryOps {
		name := unop.Module + "." + unop.Op.OperatorName() + "." + unop.Value.String()
		if g.complex(unop.Return) {
			unop.Ir = g.mod.NewFunc(
				name,
				g.lltyp(unop.Return),
				ir.NewParam("", g.lltyp(unop.Value)),
			)
			unop.Complex = true
		}

		hash := unop.Op.OperatorName() + " " + unop.Value.String()
		g.program.UnaryOps[i] = unop
		g.builtins.unops[hash] = &g.program.UnaryOps[i]
	}
}

func (g *generator) genIncDecs() {
	for i, incdec := range g.program.IncDecs {
		name := incdec.Module + "." + incdec.Op.OperatorName() + "." + incdec.Var.String()
		if g.complex(incdec.Var) {
			incdec.Ir = g.mod.NewFunc(
				name,
				g.lltyp(incdec.Var),
			)
			incdec.Complex = true
		}

		hash := incdec.Op.OperatorName() + " " + incdec.Var.String()
		g.program.IncDecs[i] = incdec
		g.builtins.incdecs[hash] = &g.program.IncDecs[i]
	}
}

func (g *generator) genComps() {
	for i, comp := range g.program.Comparisons {
		// TODO: Complex comparisons

		hash := comp.Comp.OperatorName() + " " + comp.Left.String() + " " + comp.Right.String()
		g.program.Comparisons[i] = comp
		g.builtins.comps[hash] = &g.program.Comparisons[i]
	}
}

func (g *generator) genTypeConvs() {
	for i, conv := range g.program.TypeConvs {
		name := conv.Module + ".conv." + string(conv.From) + "_" + string(conv.To)
		if g.complex(conv.To) {
			conv.Ir = g.mod.NewFunc(
				name,
				g.lltyp(conv.To),
				ir.NewParam("", g.lltyp(conv.From)),
			)
			conv.Complex = true
		}

		hash := "conv " + string(conv.From) + " " + string(conv.To)
		g.program.TypeConvs[i] = conv
		g.builtins.convs[hash] = &g.program.TypeConvs[i]
	}
}
