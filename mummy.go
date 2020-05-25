package main

import (
	"strings"
	)

func stringToArray(text string) []string{
	splited:=strings.Split(text,"")
	return splited
}
func countArray(splited []string ) int{
	count :=len(splited)
	return count
}
func obtenerPorcentajeDeVocales(splited []string, count int) int{
	vocalA:="a"
	vocalE:="e"
	vocalI:="i"
	vocalO:="o"
	vocalU:="u"
	vocalAUpper:="A"
	vocalEUpper:="E"
	vocalIUpper:="I"
	vocalOUpper:="O"
	vocalUUpper:="U"

	i:=0
	contador:=0
	for i<count{
		if (splited[i]==vocalA || splited[i]==vocalE || splited[i]==vocalI || splited[i]==vocalO || splited[i]==vocalU || splited[i]==vocalAUpper || splited[i]==vocalEUpper || splited[i]==vocalIUpper || splited[i]==vocalOUpper || splited[i]==vocalUUpper){
			contador=contador+1
		}
		i=i+1
	}
	percentage:=(contador*100/count)

	return percentage
}
func sustituirPorMummy(splited []string, count int, percentage int) []string{
	if (percentage>=30){
		vocalA:="a"
		vocalE:="e"
		vocalI:="i"
		vocalO:="o"
		vocalU:="u"
		vocalAUpper:="A"
		vocalEUpper:="E"
		vocalIUpper:="I"
		vocalOUpper:="O"
		vocalUUpper:="U"
		i:=0
		for i<count{
			if (splited[i]==vocalA || splited[i]==vocalE || splited[i]==vocalI || splited[i]==vocalO || splited[i]==vocalU || splited[i]==vocalAUpper || splited[i]==vocalEUpper || splited[i]==vocalIUpper || splited[i]==vocalOUpper || splited[i]==vocalUUpper){
				splited[i]="mummy"
			}
			i=i+1
		}
		return splited
	}
	return splited
}
func buscarDobleMummy( splited []string, count int)[]string{
	i:=1
	for i<count{
		if (splited[i]=="mummy"){
			if (splited[i-1]=="mummy"){
				splited[i]=""
			}
		}
		i=i+1
	}
	return splited
}
func arrayToString(splited []string) string{
	stringFinal:=strings.Join(splited, "")
	return stringFinal
}
func todo(text string) string{
	splited:=stringToArray(text)
	count:=countArray(splited)
	porcentaje:=obtenerPorcentajeDeVocales(splited, count)
	sustituir:=sustituirPorMummy(splited,count,porcentaje)
	limpiar:=buscarDobleMummy(sustituir,count)
	stringFinal:=arrayToString(limpiar)

	return stringFinal
}
