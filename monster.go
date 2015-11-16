package main

import (
	"fmt"
	example "./MyGame"
	flatbuffers "github.com/google/flatbuffers/go"
)

func MakeMonster(builder *flatbuffers.Builder) []byte {
	builder.Reset()

	name := []byte("One Sample Monster")
	name_position := builder.CreateByteString(name)

	example.MonsterStart(builder)
	example.MonsterAddPos(builder, example.CreateVec3(builder, 1.0, 2.0, 3.0))
	example.MonsterAddHp(builder, 80)
	example.MonsterAddName(builder, name_position)

	monster_position := example.MonsterEnd(builder)

	builder.Finish(monster_position)

	return builder.Bytes[builder.Head():]
}

func main() {
	builder := flatbuffers.NewBuilder(0)
	buf := MakeMonster(builder)

	monster := example.GetRootAsMonster(buf, 0)
	monster_name := monster.Name()

	fmt.Printf("Monster named '%s' encoded in %d bytes.\n", monster_name, len(buf))
}

