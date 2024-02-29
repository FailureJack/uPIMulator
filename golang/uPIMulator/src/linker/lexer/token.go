package lexer

type TokenType int

const (
	END_OF_FILE TokenType = iota

	NEW_LINE

	IDENTIFIER
	POSITIVIE_NUMBER
	HEX_NUMBER
	STRING
	GP_REG
	PAIR_REG

	ACQUIRE
	RELEASE
	BOOT
	RESUME

	ADD
	ADDC
	AND
	ANDN
	ASR
	CMPB4
	LSL
	LSL1
	LSL1X
	LSLX
	LSR
	LSR1
	LSR1X
	LSRX
	MUL_SH_SH
	MUL_SH_SL
	MUL_SH_UH
	MUL_SH_UL
	MUL_SL_SH
	MUL_SL_SL
	MUL_SL_UH
	MUL_SL_UL
	MUL_UH_UH
	MUL_UH_UL
	MUL_UL_UH
	MUL_UL_UL
	NAND
	NOR
	NXOR
	OR
	ORN
	ROL
	ROR
	RSUB
	RSUBC
	SUB
	SUBC
	XOR
	CALL
	HASH

	CAO
	CLO
	CLS
	CLZ
	EXTSB
	EXTSH
	EXTUB
	EXTUH
	SATS
	TIME_CFG

	DIV_STEP
	MUL_STEP

	LSL_ADD
	LSL_SUB
	LSR_ADD
	ROL_ADD

	TIME
	NOP

	STOP

	FAULT

	MOVD
	SWAPD

	LBS
	LBU
	LD
	LHS
	LHU
	LW

	SB
	SB_ID
	SD
	SD_ID
	SH
	SH_ID
	SW
	SW_ID

	LDMA
	LDMAI
	SDMA

	MOVE
	NEG
	NOT
	BKP

	JEQ
	JNEQ
	JZ
	JNZ
	JLTU
	JGTU
	JLEU
	JGEU
	JLTS
	JGTS
	JLES
	JGES
	JUMP

	S
	U

	ATOMIC
	BSS
	DATA
	DEBUG_ABBREV
	DEBUG_FRAME
	DEBUG_INFO
	DEBUG_LINE
	DEBUG_LOC
	DEBUG_RANGES
	DEBUG_STR
	DPU_HOST
	MRAM
	RODATA
	STACK_SIZES
	TEXT

	PROGBITS
	NOBITS

	FUNCTION
	OBJECT

	TRUE
	FALSE
	Z
	NZ
	E
	O
	PL
	MI
	OV
	NOV
	C
	NC
	SZ
	SNZ
	SPL
	SMI
	SO
	SE
	NC5
	NC6
	NC7
	NC8
	NC9
	NC10
	NC11
	NC12
	NC13
	NC14
	MAX
	NMAX
	SH32
	NSH32
	EQ
	NEQ
	LTU
	LEU
	GTU
	GEU
	LTS
	LES
	GTS
	GES
	XZ
	XNZ
	XLEU
	XGTU
	XLES
	XGTS
	SMALL
	LARGE

	LITTLE
	BIG

	ZERO_REG
	ONE
	ID
	ID2
	ID4
	ID8
	LNEG
	MNEG

	ADDRSIG
	ADDRSIG_SYM
	ASCII
	ASCIZ
	BYTE
	CFI_DEF_CFA_OFFSET
	CFI_ENDPROC
	CFI_OFFSET
	CFI_SECTIONS
	CFI_STARTPROC
	FILE
	GLOBL
	LOC
	LONG
	P2ALIGN
	QUAD
	SECTION
	SET
	SHORT
	SIZE
	TYPE
	WEAK
	ZERO_DIRECTIVE

	IS_STMT
	PROLOGUE_END

	COLON
	COMMA
	PLUS
	MINUS
)

type Token struct {
	token_type TokenType
	attribute  string
}

func (this *Token) Init(token_type TokenType, attribute string) {
	this.token_type = token_type
	this.attribute = attribute
}

func (this *Token) TokenType() TokenType {
	return this.token_type
}

func (this *Token) Attribute() string {
	return this.attribute
}
