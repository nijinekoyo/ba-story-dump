// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flat

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ScenarioCharacterNameExcel struct {
	_tab flatbuffers.Table
}

func GetRootAsScenarioCharacterNameExcel(buf []byte, offset flatbuffers.UOffsetT) *ScenarioCharacterNameExcel {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ScenarioCharacterNameExcel{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsScenarioCharacterNameExcel(buf []byte, offset flatbuffers.UOffsetT) *ScenarioCharacterNameExcel {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &ScenarioCharacterNameExcel{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *ScenarioCharacterNameExcel) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ScenarioCharacterNameExcel) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ScenarioCharacterNameExcel) CharacterName() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ScenarioCharacterNameExcel) MutateCharacterName(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ScenarioCharacterNameExcel) ProductionStep() ProductionStep {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return ProductionStep(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ScenarioCharacterNameExcel) MutateProductionStep(n ProductionStep) bool {
	return rcv._tab.MutateInt32Slot(6, int32(n))
}

func (rcv *ScenarioCharacterNameExcel) NameKr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScenarioCharacterNameExcel) NicknameKr() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScenarioCharacterNameExcel) NameJp() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScenarioCharacterNameExcel) NicknameJp() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScenarioCharacterNameExcel) Shape() ScenarioCharacterShapes {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return ScenarioCharacterShapes(rcv._tab.GetInt32(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *ScenarioCharacterNameExcel) MutateShape(n ScenarioCharacterShapes) bool {
	return rcv._tab.MutateInt32Slot(16, int32(n))
}

func (rcv *ScenarioCharacterNameExcel) SpinePrefabName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ScenarioCharacterNameExcel) SmallPortrait() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ScenarioCharacterNameExcelStart(builder *flatbuffers.Builder) {
	builder.StartObject(9)
}
func ScenarioCharacterNameExcelAddCharacterName(builder *flatbuffers.Builder, characterName uint32) {
	builder.PrependUint32Slot(0, characterName, 0)
}
func ScenarioCharacterNameExcelAddProductionStep(builder *flatbuffers.Builder, productionStep ProductionStep) {
	builder.PrependInt32Slot(1, int32(productionStep), 0)
}
func ScenarioCharacterNameExcelAddNameKr(builder *flatbuffers.Builder, nameKr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(nameKr), 0)
}
func ScenarioCharacterNameExcelAddNicknameKr(builder *flatbuffers.Builder, nicknameKr flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(nicknameKr), 0)
}
func ScenarioCharacterNameExcelAddNameJp(builder *flatbuffers.Builder, nameJp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(nameJp), 0)
}
func ScenarioCharacterNameExcelAddNicknameJp(builder *flatbuffers.Builder, nicknameJp flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(nicknameJp), 0)
}
func ScenarioCharacterNameExcelAddShape(builder *flatbuffers.Builder, shape ScenarioCharacterShapes) {
	builder.PrependInt32Slot(6, int32(shape), 0)
}
func ScenarioCharacterNameExcelAddSpinePrefabName(builder *flatbuffers.Builder, spinePrefabName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(spinePrefabName), 0)
}
func ScenarioCharacterNameExcelAddSmallPortrait(builder *flatbuffers.Builder, smallPortrait flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(smallPortrait), 0)
}
func ScenarioCharacterNameExcelEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
