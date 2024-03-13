package goi18n

type Direction uint

const (
	DRCTN_UNKNOWN Direction = iota
	DRCTN_FORWARD
	DRCTN_BACKWARD
	DRCTN_MIXED
)

type Axis uint

const (
	AX_UNKNOWN Axis = iota
	AX_HORIZONTAL
	AX_VERTICAL
	AX_EITHER
)

type Arrow uint

const (
	ARR_UNKNOWN Arrow = iota
	ARR_LEFT_TO_RIGHT
	ARR_RIGHT_TO_LEFT
	ARR_TOP_TO_BOTTOM
	ARR_BOTTOM_TO_TOP
	ARR_HORIZONTAL_MIXED
	ARR_VERTICAL_MIXED
	ARR_T2B_OR_L2R
	ARR_R2L_OR_T2B
)

func(arr Arrow) Split() (Axis, Direction) {
	switch arr {
		case ARR_LEFT_TO_RIGHT:
			return AX_HORIZONTAL, DRCTN_FORWARD
		case ARR_RIGHT_TO_LEFT:
			return AX_HORIZONTAL, DRCTN_BACKWARD
		case ARR_TOP_TO_BOTTOM:
			return AX_VERTICAL, DRCTN_FORWARD
		case ARR_BOTTOM_TO_TOP:
			return AX_VERTICAL, DRCTN_BACKWARD
		case ARR_HORIZONTAL_MIXED:
			return AX_HORIZONTAL, DRCTN_MIXED
		case ARR_VERTICAL_MIXED:
			return AX_VERTICAL, DRCTN_MIXED
		case ARR_T2B_OR_L2R:
			return AX_EITHER, DRCTN_FORWARD
		case ARR_R2L_OR_T2B:
			return AX_EITHER, DRCTN_UNKNOWN
		default:
			return AX_UNKNOWN, DRCTN_UNKNOWN
	}
}

type Directionality uint

const (
	DIR_UNKNOWN Directionality = iota
	DIR_L2R_T2B
	DIR_L2R_B2T
	DIR_R2L_T2B
	DIR_R2L_B2T
	DIR_T2B_L2R
	DIR_T2B_R2L
	DIR_B2T_L2R
	DIR_B2T_R2L
	DIR_BOUSTROPHEDON
	DIR_MIXED
	DIR_VARIES
	DIR_T2B_R2L_OR_L2R_T2B // "Japanese-style"
)

func(dir Directionality) NativeMajorAxis() Axis {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T, DIR_R2L_T2B, DIR_R2L_B2T:
			return AX_VERTICAL
		case DIR_T2B_L2R, DIR_T2B_R2L, DIR_B2T_L2R, DIR_B2T_R2L:
			return AX_HORIZONTAL
		case DIR_BOUSTROPHEDON:
			return AX_VERTICAL
		case DIR_MIXED, DIR_VARIES:
			return AX_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return AX_HORIZONTAL
		default:
			return AX_UNKNOWN
	}
}

func(dir Directionality) AlternativeMajorAxis() Axis {
	if dir == DIR_T2B_R2L_OR_L2R_T2B {
		return AX_VERTICAL
	} else {
		return AX_UNKNOWN
	}
}

func(dir Directionality) MajorAxis() Axis {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T, DIR_R2L_T2B, DIR_R2L_B2T:
			return AX_VERTICAL
		case DIR_T2B_L2R, DIR_T2B_R2L, DIR_B2T_L2R, DIR_B2T_R2L:
			return AX_HORIZONTAL
		case DIR_BOUSTROPHEDON:
			return AX_VERTICAL
		case DIR_MIXED, DIR_VARIES:
			return AX_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return AX_EITHER
		default:
			return AX_UNKNOWN
	}
}

func(dir Directionality) NativeMinorAxis() Axis {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T, DIR_R2L_T2B, DIR_R2L_B2T:
			return AX_HORIZONTAL
		case DIR_T2B_L2R, DIR_T2B_R2L, DIR_B2T_L2R, DIR_B2T_R2L:
			return AX_VERTICAL
		case DIR_BOUSTROPHEDON:
			return AX_HORIZONTAL
		case DIR_MIXED, DIR_VARIES:
			return AX_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return AX_VERTICAL
		default:
			return AX_UNKNOWN
	}
}

func(dir Directionality) AlternativeMinorAxis() Axis {
	if dir == DIR_T2B_R2L_OR_L2R_T2B {
		return AX_HORIZONTAL
	} else {
		return AX_UNKNOWN
	}
}

func(dir Directionality) MinorAxis() Axis {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T, DIR_R2L_T2B, DIR_R2L_B2T:
			return AX_HORIZONTAL
		case DIR_T2B_L2R, DIR_T2B_R2L, DIR_B2T_L2R, DIR_B2T_R2L:
			return AX_VERTICAL
		case DIR_BOUSTROPHEDON:
			return AX_HORIZONTAL
		case DIR_MIXED, DIR_VARIES:
			return AX_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return AX_EITHER
		default:
			return AX_UNKNOWN
	}
}

func(dir Directionality) NativeMajorArrow() Arrow {
	switch dir {
		case DIR_L2R_T2B, DIR_R2L_T2B:
			return ARR_TOP_TO_BOTTOM
		case DIR_L2R_B2T, DIR_R2L_B2T:
			return ARR_BOTTOM_TO_TOP
		case DIR_T2B_L2R, DIR_B2T_L2R:
			return ARR_LEFT_TO_RIGHT
		case DIR_T2B_R2L, DIR_B2T_R2L:
			return ARR_RIGHT_TO_LEFT
		case DIR_BOUSTROPHEDON:
			return ARR_TOP_TO_BOTTOM
		case DIR_MIXED, DIR_VARIES:
			return ARR_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return ARR_RIGHT_TO_LEFT
		default:
			return ARR_UNKNOWN
	}
}

func(dir Directionality) AlternativeMajorArrow() Arrow {
	if dir == DIR_T2B_R2L_OR_L2R_T2B {
		return ARR_TOP_TO_BOTTOM
	} else {
		return ARR_UNKNOWN
	}
}

func(dir Directionality) MajorArrow() Arrow {
	switch dir {
		case DIR_L2R_T2B, DIR_R2L_T2B:
			return ARR_TOP_TO_BOTTOM
		case DIR_L2R_B2T, DIR_R2L_B2T:
			return ARR_BOTTOM_TO_TOP
		case DIR_T2B_L2R, DIR_B2T_L2R:
			return ARR_LEFT_TO_RIGHT
		case DIR_T2B_R2L, DIR_B2T_R2L:
			return ARR_RIGHT_TO_LEFT
		case DIR_BOUSTROPHEDON:
			return ARR_TOP_TO_BOTTOM
		case DIR_MIXED, DIR_VARIES:
			return ARR_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return ARR_R2L_OR_T2B
		default:
			return ARR_UNKNOWN
	}
}

func(dir Directionality) NativeMinorArrow() Arrow {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T:
			return ARR_LEFT_TO_RIGHT
		case DIR_R2L_T2B, DIR_R2L_B2T:
			return ARR_RIGHT_TO_LEFT
		case DIR_T2B_L2R, DIR_T2B_R2L:
			return ARR_TOP_TO_BOTTOM
		case DIR_B2T_L2R, DIR_B2T_R2L:
			return ARR_BOTTOM_TO_TOP
		case DIR_BOUSTROPHEDON:
			return ARR_HORIZONTAL_MIXED
		case DIR_MIXED, DIR_VARIES:
			return ARR_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return ARR_TOP_TO_BOTTOM
		default:
			return ARR_UNKNOWN
	}
}

func(dir Directionality) AlternativeMinorArrow() Arrow {
	if dir == DIR_T2B_R2L_OR_L2R_T2B {
		return ARR_LEFT_TO_RIGHT
	} else {
		return ARR_UNKNOWN
	}
}

func(dir Directionality) MinorArrow() Arrow {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T:
			return ARR_LEFT_TO_RIGHT
		case DIR_R2L_T2B, DIR_R2L_B2T:
			return ARR_RIGHT_TO_LEFT
		case DIR_T2B_L2R, DIR_T2B_R2L:
			return ARR_TOP_TO_BOTTOM
		case DIR_B2T_L2R, DIR_B2T_R2L:
			return ARR_BOTTOM_TO_TOP
		case DIR_BOUSTROPHEDON:
			return ARR_HORIZONTAL_MIXED
		case DIR_MIXED, DIR_VARIES:
			return ARR_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			return ARR_T2B_OR_L2R
		default:
			return ARR_UNKNOWN
	}
}

func(dir Directionality) MinorAxisByMajorAxis(major Axis) Axis {
	switch dir {
		case DIR_L2R_T2B, DIR_R2L_T2B, DIR_L2R_B2T, DIR_R2L_B2T, DIR_BOUSTROPHEDON:
			if major == AX_VERTICAL {
				return AX_HORIZONTAL
			} else {
				return AX_UNKNOWN
			}
		case DIR_T2B_L2R, DIR_T2B_R2L, DIR_B2T_L2R, DIR_B2T_R2L:
			if major == AX_HORIZONTAL {
				return AX_VERTICAL
			} else {
				return AX_UNKNOWN
			}
		case DIR_MIXED, DIR_VARIES:
			return AX_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			switch major {
				case AX_HORIZONTAL:
					return AX_VERTICAL
				case AX_VERTICAL:
					return AX_HORIZONTAL
				default:
					return AX_UNKNOWN
			}
		default:
			return AX_UNKNOWN
	}
}

func(dir Directionality) MinorArrowByMajorAxis(major Axis) Arrow {
	switch dir {
		case DIR_L2R_T2B, DIR_L2R_B2T:
			if major == AX_VERTICAL {
				return ARR_LEFT_TO_RIGHT
			} else {
				return ARR_UNKNOWN
			}
		case DIR_R2L_T2B, DIR_R2L_B2T:
			if major == AX_VERTICAL {
				return ARR_RIGHT_TO_LEFT
			} else {
				return ARR_UNKNOWN
			}
		case DIR_T2B_L2R, DIR_T2B_R2L:
			if major == AX_HORIZONTAL {
				return ARR_TOP_TO_BOTTOM
			} else {
				return ARR_UNKNOWN
			}
		case DIR_B2T_L2R, DIR_B2T_R2L:
			if major == AX_HORIZONTAL {
				return ARR_BOTTOM_TO_TOP
			} else {
				return ARR_UNKNOWN
			}
		case DIR_BOUSTROPHEDON:
			if major == AX_VERTICAL {
				return ARR_HORIZONTAL_MIXED
			} else {
				return ARR_UNKNOWN
			}
		case DIR_MIXED, DIR_VARIES:
			return ARR_UNKNOWN
		case DIR_T2B_R2L_OR_L2R_T2B:
			switch major {
				case AX_HORIZONTAL:
					return ARR_TOP_TO_BOTTOM
				case AX_VERTICAL:
					return ARR_LEFT_TO_RIGHT
				default:
					return ARR_UNKNOWN
			}
		default:
			return ARR_UNKNOWN
	}
}
