define void @__run() {
block-0:
	%bathtub_drawn_water_level = alloca double
	store double 5.0, double* %bathtub_drawn_water_level
	call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\3114905d4a4c10fbc4b5c21a78c55b23a !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\3114905d4a4c10fbc4b5c21a78c55b23a !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !f6310547ab23560514f17d3aceeb856a !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !f6310547ab23560514f17d3aceeb856a !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\3774fb5b072ce98c5739c1a1d79f3a750 !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\3774fb5b072ce98c5739c1a1d79f3a750 !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\30ededaa6a0dcd9894de44fee69ad2679 !DIBasicType(tag: DW_TAG_string_type)
	call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\30ededaa6a0dcd9894de44fee69ad2679 !DIBasicType(tag: DW_TAG_string_type)
	ret void
}

define void @bathtub_drawn_in(double* %bathtub_drawn_water_level) {
block-1:
	%0 = load double, double* %bathtub_drawn_water_level
	%1 = fadd double %0, 10.0
	store double %1, double* %bathtub_drawn_water_level
	ret void
}

define void @bathtub_drawn_out(double* %bathtub_drawn_water_level) {
block-2:
	%0 = load double, double* %bathtub_drawn_water_level
	%1 = fsub double %0, 20.0
	store double %1, double* %bathtub_drawn_water_level
	ret void
}