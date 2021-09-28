; ModuleID = 'testdata/bathtub2.ll'
source_filename = "testdata/bathtub2.ll"

define void @__run() {
block-0:
  %bathtub_drawn_water_level = alloca double, align 8
  store double 5.000000e+00, double* %bathtub_drawn_water_level, align 8
  call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\3114905d4a4c10fbc4b5c21a78c55b23a !0
  call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\3114905d4a4c10fbc4b5c21a78c55b23a !0
  call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !f6310547ab23560514f17d3aceeb856a !0
  call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !f6310547ab23560514f17d3aceeb856a !0
  call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\3774fb5b072ce98c5739c1a1d79f3a750 !0
  call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\3774fb5b072ce98c5739c1a1d79f3a750 !0
  call void @bathtub_drawn_in(double* %bathtub_drawn_water_level), !\30ededaa6a0dcd9894de44fee69ad2679 !0
  call void @bathtub_drawn_out(double* %bathtub_drawn_water_level), !\30ededaa6a0dcd9894de44fee69ad2679 !0
  ret void
}

define void @bathtub_drawn_in(double* %bathtub_drawn_water_level) {
block-1:
  %0 = load double, double* %bathtub_drawn_water_level, align 8
  %1 = fadd double %0, 1.000000e+01
  store double %1, double* %bathtub_drawn_water_level, align 8
  ret void
}

define void @bathtub_drawn_out(double* %bathtub_drawn_water_level) {
block-2:
  %0 = load double, double* %bathtub_drawn_water_level, align 8
  %1 = fsub double %0, 2.000000e+01
  store double %1, double* %bathtub_drawn_water_level, align 8
  ret void
}

!0 = !DIBasicType(tag: DW_TAG_string_type)