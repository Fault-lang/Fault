(set-logic QF_LRA)
(declare-fun bathtub2_drawn_water_level_0 () Real)
(declare-fun bathtub2_drawn_water_level_1 () Real)
(declare-fun bathtub2_drawn_water_level_2 () Real)
(declare-fun bathtub2_drawn_water_level_3 () Real)
(declare-fun @__run_0 () Bool)
(declare-fun bathtub2_drawn_water_level_4 () Real)
(declare-fun bathtub2_drawn_water_level_5 () Real)
(declare-fun bathtub2_drawn_water_level_6 () Real)
(declare-fun @__run_1 () Bool)
(declare-fun bathtub2_drawn_water_level_7 () Real)
(declare-fun bathtub2_drawn_water_level_8 () Real)
(declare-fun bathtub2_drawn_water_level_9 () Real)
(declare-fun bathtub2_drawn_water_level_10 () Real)
(declare-fun bathtub2_drawn_water_level_11 () Real)
(declare-fun bathtub2_drawn_water_level_12 () Real)
(declare-fun bathtub2_drawn_water_level_13 () Real)
(declare-fun bathtub2_drawn_water_level_14 () Real)
(declare-fun bathtub2_drawn_water_level_15 () Real)
(declare-fun bathtub2_drawn_water_level_16 () Real)
(declare-fun bathtub2_drawn_water_level_17 () Real)
(declare-fun bathtub2_drawn_water_level_18 () Real)
(assert (= bathtub2_drawn_water_level_0 5.0))

(assert (= bathtub2_drawn_water_level_1 (+ bathtub2_drawn_water_level_0 10.0)))


(assert (= bathtub2_drawn_water_level_2 (- bathtub2_drawn_water_level_1 20.0)))

(assert (= @__run_0 (= bathtub2_drawn_water_level_3 bathtub2_drawn_water_level_2)))

(assert (= bathtub2_drawn_water_level_4 (- bathtub2_drawn_water_level_0 20.0)))


(assert (= bathtub2_drawn_water_level_5 (+ bathtub2_drawn_water_level_4 10.0)))

(assert (= @__run_1 (and (= bathtub2_drawn_water_level_6 bathtub2_drawn_water_level_2)
(= bathtub2_drawn_water_level_6 bathtub2_drawn_water_level_5))))
(assert (distinct @__run_0 @__run_1))

(assert (= bathtub2_drawn_water_level_7 (+ bathtub2_drawn_water_level_6 10.0)))


(assert (= bathtub2_drawn_water_level_8 (- bathtub2_drawn_water_level_7 20.0)))

(assert (= @__run_0 (= bathtub2_drawn_water_level_9 bathtub2_drawn_water_level_8)))

(assert (= bathtub2_drawn_water_level_10 (- bathtub2_drawn_water_level_6 20.0)))


(assert (= bathtub2_drawn_water_level_11 (+ bathtub2_drawn_water_level_10 10.0)))

(assert (= @__run_1 (and (= bathtub2_drawn_water_level_12 bathtub2_drawn_water_level_8)
(= bathtub2_drawn_water_level_12 bathtub2_drawn_water_level_11))))
(assert (distinct @__run_0 @__run_1))

(assert (= bathtub2_drawn_water_level_13 (+ bathtub2_drawn_water_level_12 10.0)))


(assert (= bathtub2_drawn_water_level_14 (- bathtub2_drawn_water_level_13 20.0)))

(assert (= @__run_0 (= bathtub2_drawn_water_level_15 bathtub2_drawn_water_level_14)))

(assert (= bathtub2_drawn_water_level_16 (- bathtub2_drawn_water_level_12 20.0)))


(assert (= bathtub2_drawn_water_level_17 (+ bathtub2_drawn_water_level_16 10.0)))

(assert (= @__run_1 (and (= bathtub2_drawn_water_level_18 bathtub2_drawn_water_level_14)
(= bathtub2_drawn_water_level_18 bathtub2_drawn_water_level_17))))
(assert (distinct @__run_0 @__run_1))