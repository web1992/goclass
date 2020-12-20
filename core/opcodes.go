package core

const (

	// Reserved Opcodes
	impdep1    = 254 // 0xfe
	impdep2    = 255 // 0xff
	breakpoint = 202 // 0xca

	JVM_OPC_nop             = 0
	JVM_OPC_aconst_null     = 1
	JVM_OPC_iconst_m1       = 2
	JVM_OPC_iconst_0        = 3
	JVM_OPC_iconst_1        = 4
	JVM_OPC_iconst_2        = 5
	JVM_OPC_iconst_3        = 6
	JVM_OPC_iconst_4        = 7
	JVM_OPC_iconst_5        = 8
	JVM_OPC_lconst_0        = 9
	JVM_OPC_lconst_1        = 10
	JVM_OPC_fconst_0        = 11
	JVM_OPC_fconst_1        = 12
	JVM_OPC_fconst_2        = 13
	JVM_OPC_dconst_0        = 14
	JVM_OPC_dconst_1        = 15
	JVM_OPC_bipush          = 16
	JVM_OPC_sipush          = 17
	JVM_OPC_ldc             = 18
	JVM_OPC_ldc_w           = 19
	JVM_OPC_ldc2_w          = 20
	JVM_OPC_iload           = 21
	JVM_OPC_lload           = 22
	JVM_OPC_fload           = 23
	JVM_OPC_dload           = 24
	JVM_OPC_aload           = 25
	JVM_OPC_iload_0         = 26
	JVM_OPC_iload_1         = 27
	JVM_OPC_iload_2         = 28
	JVM_OPC_iload_3         = 29
	JVM_OPC_lload_0         = 30
	JVM_OPC_lload_1         = 31
	JVM_OPC_lload_2         = 32
	JVM_OPC_lload_3         = 33
	JVM_OPC_fload_0         = 34
	JVM_OPC_fload_1         = 35
	JVM_OPC_fload_2         = 36
	JVM_OPC_fload_3         = 37
	JVM_OPC_dload_0         = 38
	JVM_OPC_dload_1         = 39
	JVM_OPC_dload_2         = 40
	JVM_OPC_dload_3         = 41
	JVM_OPC_aload_0         = 42
	JVM_OPC_aload_1         = 43
	JVM_OPC_aload_2         = 44
	JVM_OPC_aload_3         = 45
	JVM_OPC_iaload          = 46
	JVM_OPC_laload          = 47
	JVM_OPC_faload          = 48
	JVM_OPC_daload          = 49
	JVM_OPC_aaload          = 50
	JVM_OPC_baload          = 51
	JVM_OPC_caload          = 52
	JVM_OPC_saload          = 53
	JVM_OPC_istore          = 54
	JVM_OPC_lstore          = 55
	JVM_OPC_fstore          = 56
	JVM_OPC_dstore          = 57
	JVM_OPC_astore          = 58
	JVM_OPC_istore_0        = 59
	JVM_OPC_istore_1        = 60
	JVM_OPC_istore_2        = 61
	JVM_OPC_istore_3        = 62
	JVM_OPC_lstore_0        = 63
	JVM_OPC_lstore_1        = 64
	JVM_OPC_lstore_2        = 65
	JVM_OPC_lstore_3        = 66
	JVM_OPC_fstore_0        = 67
	JVM_OPC_fstore_1        = 68
	JVM_OPC_fstore_2        = 69
	JVM_OPC_fstore_3        = 70
	JVM_OPC_dstore_0        = 71
	JVM_OPC_dstore_1        = 72
	JVM_OPC_dstore_2        = 73
	JVM_OPC_dstore_3        = 74
	JVM_OPC_astore_0        = 75
	JVM_OPC_astore_1        = 76
	JVM_OPC_astore_2        = 77
	JVM_OPC_astore_3        = 78
	JVM_OPC_iastore         = 79
	JVM_OPC_lastore         = 80
	JVM_OPC_fastore         = 81
	JVM_OPC_dastore         = 82
	JVM_OPC_aastore         = 83
	JVM_OPC_bastore         = 84
	JVM_OPC_castore         = 85
	JVM_OPC_sastore         = 86
	JVM_OPC_pop             = 87
	JVM_OPC_pop2            = 88
	JVM_OPC_dup             = 89
	JVM_OPC_dup_x1          = 90
	JVM_OPC_dup_x2          = 91
	JVM_OPC_dup2            = 92
	JVM_OPC_dup2_x1         = 93
	JVM_OPC_dup2_x2         = 94
	JVM_OPC_swap            = 95
	JVM_OPC_iadd            = 96
	JVM_OPC_ladd            = 97
	JVM_OPC_fadd            = 98
	JVM_OPC_dadd            = 99
	JVM_OPC_isub            = 100
	JVM_OPC_lsub            = 101
	JVM_OPC_fsub            = 102
	JVM_OPC_dsub            = 103
	JVM_OPC_imul            = 104
	JVM_OPC_lmul            = 105
	JVM_OPC_fmul            = 106
	JVM_OPC_dmul            = 107
	JVM_OPC_idiv            = 108
	JVM_OPC_ldiv            = 109
	JVM_OPC_fdiv            = 110
	JVM_OPC_ddiv            = 111
	JVM_OPC_irem            = 112
	JVM_OPC_lrem            = 113
	JVM_OPC_frem            = 114
	JVM_OPC_drem            = 115
	JVM_OPC_ineg            = 116
	JVM_OPC_lneg            = 117
	JVM_OPC_fneg            = 118
	JVM_OPC_dneg            = 119
	JVM_OPC_ishl            = 120
	JVM_OPC_lshl            = 121
	JVM_OPC_ishr            = 122
	JVM_OPC_lshr            = 123
	JVM_OPC_iushr           = 124
	JVM_OPC_lushr           = 125
	JVM_OPC_iand            = 126
	JVM_OPC_land            = 127
	JVM_OPC_ior             = 128
	JVM_OPC_lor             = 129
	JVM_OPC_ixor            = 130
	JVM_OPC_lxor            = 131
	JVM_OPC_iinc            = 132
	JVM_OPC_i2l             = 133
	JVM_OPC_i2f             = 134
	JVM_OPC_i2d             = 135
	JVM_OPC_l2i             = 136
	JVM_OPC_l2f             = 137
	JVM_OPC_l2d             = 138
	JVM_OPC_f2i             = 139
	JVM_OPC_f2l             = 140
	JVM_OPC_f2d             = 141
	JVM_OPC_d2i             = 142
	JVM_OPC_d2l             = 143
	JVM_OPC_d2f             = 144
	JVM_OPC_i2b             = 145
	JVM_OPC_i2c             = 146
	JVM_OPC_i2s             = 147
	JVM_OPC_lcmp            = 148
	JVM_OPC_fcmpl           = 149
	JVM_OPC_fcmpg           = 150
	JVM_OPC_dcmpl           = 151
	JVM_OPC_dcmpg           = 152
	JVM_OPC_ifeq            = 153
	JVM_OPC_ifne            = 154
	JVM_OPC_iflt            = 155
	JVM_OPC_ifge            = 156
	JVM_OPC_ifgt            = 157
	JVM_OPC_ifle            = 158
	JVM_OPC_if_icmpeq       = 159
	JVM_OPC_if_icmpne       = 160
	JVM_OPC_if_icmplt       = 161
	JVM_OPC_if_icmpge       = 162
	JVM_OPC_if_icmpgt       = 163
	JVM_OPC_if_icmple       = 164
	JVM_OPC_if_acmpeq       = 165
	JVM_OPC_if_acmpne       = 166
	JVM_OPC_goto            = 167
	JVM_OPC_jsr             = 168
	JVM_OPC_ret             = 169
	JVM_OPC_tableswitch     = 170
	JVM_OPC_lookupswitch    = 171
	JVM_OPC_ireturn         = 172
	JVM_OPC_lreturn         = 173
	JVM_OPC_freturn         = 174
	JVM_OPC_dreturn         = 175
	JVM_OPC_areturn         = 176
	JVM_OPC_return          = 177
	JVM_OPC_getstatic       = 178
	JVM_OPC_putstatic       = 179
	JVM_OPC_getfield        = 180
	JVM_OPC_putfield        = 181
	JVM_OPC_invokevirtual   = 182
	JVM_OPC_invokespecial   = 183
	JVM_OPC_invokestatic    = 184
	JVM_OPC_invokeinterface = 185
	JVM_OPC_invokedynamic   = 186
	JVM_OPC_new             = 187
	JVM_OPC_newarray        = 188
	JVM_OPC_anewarray       = 189
	JVM_OPC_arraylength     = 190
	JVM_OPC_athrow          = 191
	JVM_OPC_checkcast       = 192
	JVM_OPC_instanceof      = 193
	JVM_OPC_monitorenter    = 194
	JVM_OPC_monitorexit     = 195
	JVM_OPC_wide            = 196
	JVM_OPC_multianewarray  = 197
	JVM_OPC_ifnull          = 198
	JVM_OPC_ifnonnull       = 199
	JVM_OPC_goto_w          = 200
	JVM_OPC_jsr_w           = 201
	JVM_OPC_MAX             = 201
)

type OpCodes []interface{}

type OpCode struct {
	Desc string
	Opc  int32
	Args []int32
}

type OpCodeInvoke struct {
	Desc string
	Opc  int32
	Args []int32
}

func (op *OpCodeInvoke) ReadObj(bytes []byte) int {
	return 0
}
func (op *OpCodeInvoke) ObjLen() int {
	return 0
}
