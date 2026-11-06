(set-logic QF_LRA)
(declare-fun bathtub_drawn_water_level_0 () Real)
(declare-fun bathtub_pipe_water_level_0 () Real)
(declare-fun bathtub_drawn_water_level_1 () Real)
(declare-fun bathtub_pipe_water_level_1 () Real)
(declare-fun bathtub_drawn_water_level_2 () Real)
(declare-fun bathtub_pipe_water_level_2 () Real)
(declare-fun @__run_0 () Bool)
(declare-fun bathtub_pipe_water_level_3 () Real)
(declare-fun bathtub_drawn_water_level_3 () Real)
(declare-fun bathtub_drawn_water_level_4 () Real)
(declare-fun bathtub_pipe_water_level_4 () Real)
(declare-fun @__run_1 () Bool)
(declare-fun bathtub_drawn_water_level_5 () Real)
(declare-fun bathtub_pipe_water_level_5 () Real)
(declare-fun bathtub_drawn_water_level_6 () Real)
(declare-fun bathtub_pipe_water_level_6 () Real)
(declare-fun bathtub_pipe_water_level_7 () Real)
(declare-fun bathtub_drawn_water_level_7 () Real)
(declare-fun bathtub_drawn_water_level_8 () Real)
(declare-fun bathtub_pipe_water_level_8 () Real)
(declare-fun bathtub_drawn_water_level_9 () Real)
(declare-fun bathtub_pipe_water_level_9 () Real)
(declare-fun bathtub_drawn_water_level_10 () Real)
(declare-fun bathtub_pipe_water_level_10 () Real)
(declare-fun bathtub_pipe_water_level_11 () Real)
(declare-fun bathtub_drawn_water_level_11 () Real)
(declare-fun bathtub_drawn_water_level_12 () Real)
(declare-fun bathtub_pipe_water_level_12 () Real)
(declare-fun bathtub_drawn_water_level_13 () Real)
(declare-fun bathtub_pipe_water_level_13 () Real)
(declare-fun bathtub_drawn_water_level_14 () Real)
(declare-fun bathtub_pipe_water_level_14 () Real)
(declare-fun bathtub_pipe_water_level_15 () Real)
(declare-fun bathtub_drawn_water_level_15 () Real)
(declare-fun bathtub_drawn_water_level_16 () Real)
(declare-fun bathtub_pipe_water_level_16 () Real)
(assert (= bathtub_drawn_water_level_0 5.0))
(assert (= bathtub_pipe_water_level_0 5.0))

(assert (= bathtub_drawn_water_level_1 (+ bathtub_drawn_water_level_0 10.0)))


(assert (= bathtub_pipe_water_level_1 (- bathtub_pipe_water_level_0 20.0)))

(assert (= @__run_0 (and (= bathtub_drawn_water_level_2 bathtub_drawn_water_level_1) (= bathtub_pipe_water_level_2 bathtub_pipe_water_level_1))))

(assert (= bathtub_pipe_water_level_3 (- bathtub_pipe_water_level_0 20.0)))


(assert (= bathtub_drawn_water_level_3 (+ bathtub_drawn_water_level_0 10.0)))

(assert (= @__run_1 (and (and (= bathtub_drawn_water_level_4 bathtub_drawn_water_level_1) (= bathtub_pipe_water_level_4 bathtub_pipe_water_level_1))
(and (= bathtub_drawn_water_level_4 bathtub_drawn_water_level_3) (= bathtub_pipe_water_level_4 bathtub_pipe_water_level_3)))))
(assert (distinct @__run_0 @__run_1))

(assert (= bathtub_drawn_water_level_5 (+ bathtub_drawn_water_level_4 10.0)))


(assert (= bathtub_pipe_water_level_5 (- bathtub_pipe_water_level_4 20.0)))

(assert (= @__run_0 (and (= bathtub_drawn_water_level_6 bathtub_drawn_water_level_5) (= bathtub_pipe_water_level_6 bathtub_pipe_water_level_5))))

(assert (= bathtub_pipe_water_level_7 (- bathtub_pipe_water_level_4 20.0)))


(assert (= bathtub_drawn_water_level_7 (+ bathtub_drawn_water_level_4 10.0)))

(assert (= @__run_1 (and (and (= bathtub_drawn_water_level_8 bathtub_drawn_water_level_5) (= bathtub_pipe_water_level_8 bathtub_pipe_water_level_5))
(and (= bathtub_pipe_water_level_8 bathtub_pipe_water_level_7) (= bathtub_drawn_water_level_8 bathtub_drawn_water_level_7)))))
(assert (distinct @__run_0 @__run_1))

(assert (= bathtub_drawn_water_level_9 (+ bathtub_drawn_water_level_8 10.0)))


(assert (= bathtub_pipe_water_level_9 (- bathtub_pipe_water_level_8 20.0)))

(assert (= @__run_0 (and (= bathtub_drawn_water_level_10 bathtub_drawn_water_level_9) (= bathtub_pipe_water_level_10 bathtub_pipe_water_level_9))))

(assert (= bathtub_pipe_water_level_11 (- bathtub_pipe_water_level_8 20.0)))


(assert (= bathtub_drawn_water_level_11 (+ bathtub_drawn_water_level_8 10.0)))

(assert (= @__run_1 (and (and (= bathtub_drawn_water_level_12 bathtub_drawn_water_level_9) (= bathtub_pipe_water_level_12 bathtub_pipe_water_level_9))
(and (= bathtub_pipe_water_level_12 bathtub_pipe_water_level_11) (= bathtub_drawn_water_level_12 bathtub_drawn_water_level_11)))))
(assert (distinct @__run_0 @__run_1))

(assert (= bathtub_drawn_water_level_13 (+ bathtub_drawn_water_level_12 10.0)))


(assert (= bathtub_pipe_water_level_13 (- bathtub_pipe_water_level_12 20.0)))

(assert (= @__run_0 (and (= bathtub_drawn_water_level_14 bathtub_drawn_water_level_13) (= bathtub_pipe_water_level_14 bathtub_pipe_water_level_13))))

(assert (= bathtub_pipe_water_level_15 (- bathtub_pipe_water_level_12 20.0)))


(assert (= bathtub_drawn_water_level_15 (+ bathtub_drawn_water_level_12 10.0)))

(assert (= @__run_1 (and (and (= bathtub_drawn_water_level_16 bathtub_drawn_water_level_13) (= bathtub_pipe_water_level_16 bathtub_pipe_water_level_13))
(and (= bathtub_drawn_water_level_16 bathtub_drawn_water_level_15) (= bathtub_pipe_water_level_16 bathtub_pipe_water_level_15)))))
(assert (distinct @__run_0 @__run_1))