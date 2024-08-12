package main

import (
	_ "notable_notes/examples/gomistakes" // in this way we can call init of the package .. (there is no direct use of public function)
	"notable_notes/examples/goprogramminglanguage/datatypes"
)

// func init() {
// 	fmt.Println("init from the main")
// }

func main() {
	// s, sep := "", ""
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }
	// fmt.Println(s)

	// comparetimes.ConcatStr()
	// comparetimes.StrJoinFnc()
	// goprogramminglanguage.IdentifyDuplicate()
	// goprogramminglanguage.IdentifyDuplicateTwo()
	// goprogramminglanguage.FetchUrl()
	// formats.PointerThings()
	// formats.FlagImplementation()
	// formats.Declaration()
	// datatypes.UtilizeStrings()
	// garbagecollector.One()
	// datatypes.PrintIotas()
	// datatypes.ArrayManipulation()
	// concurencyparallelism.GoRoutineSample()
	// gomistakes.VariableShadowing()
	// gomistakes.LearnInit()
	// gettersetters.GettersSetters()
	// interfaces.InterfaceFunctions()
	// interfaces.TypeCorecions("ksjfkajf")
	// interfaces.AnimalSpeaker()
	// fmt.Printf("%T\n",interfaces.GetString("mosses"))
	// dirstructure.PrintFilePaths()
	// generics.ShowResultsBeforeGeneric()
	// generics.RestructuredPrintSums()
	// datatypes.UseIntegers()
	// datatypes.FloatingPoints()
	// datatypes.SlicesPractice()
	// datatypes.NilVsEmptySlice()
	// datatypes.CheckingNilEmptySlices()
	// datatypes.CopySlices()
	// datatypes.UnexpectedSideEffects()

	// datatypes.LargeSlice()
	// datatypes.CheckMemory()
	datatypes.MapFunctions()
}

// Notes

// Since strings in Go are immutable, the concatenation operation creates a new string that combines the previous value of s,
// a space character, and the current arg. This new string is then assigned back to the variable s.
// Each time the += statement is executed, a new string is created in memory.
// The previous value of s becomes eligible for garbage collection because it is no longer referenced by any variables.
