package binchunk

type header struct {
	signature       [4]byte
	version         byte
	format          byte
	luaData         [6]byte
	cintSize        byte
	sizeSize        byte
	instructionSize byte
	luaIntegerSize  byte
	luaNumberSize   byte
	luacInt         int64
	luacNum         float64
}

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}
