(set-logic QF_NRA)
(declare-fun mixed_a_option2_1 () Bool)
(declare-fun mixed_a_choice_2 () Bool)
(declare-fun mixed_a_option3_1 () Bool)
(declare-fun mixed_a_choice_3 () Bool)
(declare-fun mixed_a_option1_1 () Bool)
(declare-fun mixed_a_choice_4 () Bool)
(declare-fun mixed_a_option1_2 () Bool)
(declare-fun mixed_a_choice_5 () Bool)
(declare-fun mixed_a_option3_2 () Bool)
(declare-fun mixed_a_option2_2 () Bool)
(declare-fun mixed_a_option2_3 () Bool)
(declare-fun mixed_a_choice_6 () Bool)
(declare-fun mixed_a_option3_3 () Bool)
(declare-fun mixed_a_option1_3 () Bool)
(declare-fun mixed_a_option1_4 () Bool)
(declare-fun mixed_a_option1_5 () Bool)
(declare-fun mixed_a_option2_4 () Bool)
(declare-fun mixed_a_option1_6 () Bool)
(declare-fun mixed_a_option3_4 () Bool)
(declare-fun mixed_a_option1_7 () Bool)
(declare-fun mixed_a_option3_5 () Bool)
(declare-fun mixed_a_option1_8 () Bool)
(declare-fun mixed_a_option2_5 () Bool)
(declare-fun mixed_a_option1_9 () Bool)
(declare-fun mixed_a_option2_6 () Bool)
(declare-fun mixed_a_option3_6 () Bool)
(declare-fun mixed_a_option1_10 () Bool)
(declare-fun mixed_a_option2_7 () Bool)
(declare-fun mixed_a_option3_7 () Bool)
(declare-fun mixed_a_option2_8 () Bool)
(declare-fun mixed_a_option1_11 () Bool)
(declare-fun mixed_a_option2_9 () Bool)
(declare-fun mixed_a_option3_8 () Bool)
(declare-fun mixed_a_option3_9 () Bool)
(declare-fun mixed_a_option2_10 () Bool)
(declare-fun mixed_a_option1_12 () Bool)
(declare-fun mixed_a_option1_13 () Bool)
(declare-fun mixed_a_option3_10 () Bool)
(declare-fun mixed_a_option2_11 () Bool)
(declare-fun mixed_a_option3_11 () Bool)
(declare-fun mixed_a_option3_12 () Bool)
(declare-fun mixed_a_option3_13 () Bool)
(declare-fun mixed_a_option3_14 () Bool)
(declare-fun mixed_a_option1_14 () Bool)
(declare-fun mixed_a_option2_12 () Bool)
(declare-fun mixed_a_option1_15 () Bool)
(declare-fun mixed_a_option3_15 () Bool)
(declare-fun mixed_a_option2_13 () Bool)
(declare-fun mixed_a_option1_16 () Bool)
(declare-fun mixed_a_option3_16 () Bool)
(declare-fun mixed_a_option2_14 () Bool)
(declare-fun mixed_a_option3_17 () Bool)
(declare-fun mixed_a_option1_17 () Bool)
(declare-fun mixed_a_option3_18 () Bool)
(declare-fun mixed_a_option2_15 () Bool)
(declare-fun mixed_a_option2_16 () Bool)
(declare-fun mixed_a_option3_19 () Bool)
(declare-fun mixed_a_option1_18 () Bool)
(declare-fun mixed_fl_active_0 () Bool)
(declare-fun mixed_fl_vault_value_0 () Real)
(declare-fun mixed_a_choice_0 () Bool)
(declare-fun mixed_a_option1_0 () Bool)
(declare-fun mixed_a_option2_0 () Bool)
(declare-fun mixed_a_option3_0 () Bool)
(declare-fun mixed_a_choice_1 () Bool)(assert (= mixed_fl_active_0 false))
(assert (= mixed_fl_vault_value_0 30.0))
(assert (= mixed_a_choice_0 false))
(assert (= mixed_a_option1_0 false))
(assert (= mixed_a_option2_0 false))
(assert (= mixed_a_option3_0 false))
(assert (= mixed_a_choice_1 true))
(assert (or (and (= mixed_a_option2_1 true)(= mixed_a_choice_2 false)(= mixed_a_option3_1 true)(= mixed_a_choice_3 false)(= mixed_a_option3_2 mixed_a_option3_1)(= mixed_a_choice_5 mixed_a_choice_3)(= mixed_a_option2_2 mixed_a_option2_1)(= mixed_a_option1_2 mixed_a_option1_0))(and (= mixed_a_option1_1 true)(= mixed_a_choice_4 false)(= mixed_a_option1_2 mixed_a_option1_1)(= mixed_a_choice_5 mixed_a_choice_4)(= mixed_a_option2_2 mixed_a_option2_0)(= mixed_a_option3_2 mixed_a_option3_0))))
(assert (ite (= mixed_a_choice_1 true) (and (= mixed_a_option2_3 mixed_a_option2_2) (= mixed_a_choice_6 mixed_a_choice_5) (= mixed_a_option3_3 mixed_a_option3_2) (= mixed_a_option1_3 mixed_a_option1_2)) (and (= mixed_a_option2_3 mixed_a_option2_0) (= mixed_a_choice_6 mixed_a_choice_1) (= mixed_a_option3_3 mixed_a_option3_0) (= mixed_a_option1_3 mixed_a_option1_0))))
(assert (or (and (= mixed_a_option1_4 true)(= mixed_a_option1_5 false)(= mixed_a_option2_4 true)(= mixed_a_option1_6 false)(= mixed_a_option1_8 mixed_a_option1_6)(= mixed_a_option2_5 mixed_a_option2_4)(= mixed_a_option3_5 mixed_a_option3_3))(and (= mixed_a_option3_4 true)(= mixed_a_option1_7 false)(= mixed_a_option3_5 mixed_a_option3_4)(= mixed_a_option1_8 mixed_a_option1_7)(= mixed_a_option2_5 mixed_a_option2_3))))
(assert (ite (= mixed_a_option1_3 true) (and (= mixed_a_option1_9 mixed_a_option1_8) (= mixed_a_option2_6 mixed_a_option2_5) (= mixed_a_option3_6 mixed_a_option3_5)) (and (= mixed_a_option1_9 mixed_a_option1_3) (= mixed_a_option2_6 mixed_a_option2_3) (= mixed_a_option3_6 mixed_a_option3_3))))
(assert (or (and (= mixed_a_option3_7 true)(= mixed_a_option2_8 false)(= mixed_a_option3_8 mixed_a_option3_7)(= mixed_a_option2_9 mixed_a_option2_8)(= mixed_a_option1_11 mixed_a_option1_9))(and (= mixed_a_option1_10 true)(= mixed_a_option2_7 false)(= mixed_a_option1_11 mixed_a_option1_10)(= mixed_a_option2_9 mixed_a_option2_7)(= mixed_a_option3_8 mixed_a_option3_6))))
(assert (ite (and (= mixed_a_option2_6 true) (= mixed_fl_active_0 true)) (and (= mixed_a_option3_9 mixed_a_option3_8) (= mixed_a_option2_10 mixed_a_option2_9) (= mixed_a_option1_12 mixed_a_option1_11)) (and (= mixed_a_option3_9 mixed_a_option3_6) (= mixed_a_option2_10 mixed_a_option2_6) (= mixed_a_option1_12 mixed_a_option1_9))))
(assert (or (and (= mixed_a_option1_13 true)(= mixed_a_option3_10 false)(= mixed_a_option2_11 true)(= mixed_a_option3_11 false)(= mixed_a_option1_14 mixed_a_option1_13)(= mixed_a_option2_12 mixed_a_option2_11)(= mixed_a_option3_14 mixed_a_option3_11))(and (= mixed_a_option3_12 true)(= mixed_a_option3_13 false)(= mixed_a_option3_14 mixed_a_option3_13)(= mixed_a_option2_12 mixed_a_option2_10)(= mixed_a_option1_14 mixed_a_option1_12))))
(assert (ite (= mixed_a_option3_9 true) (and (= mixed_a_option1_15 mixed_a_option1_14) (= mixed_a_option3_15 mixed_a_option3_14) (= mixed_a_option2_13 mixed_a_option2_12)) (and (= mixed_a_option3_15 mixed_a_option3_9) (= mixed_a_option2_13 mixed_a_option2_10) (= mixed_a_option1_15 mixed_a_option1_12))))
(assert (or (and (= mixed_a_option2_14 true)(= mixed_a_option3_17 false)(= mixed_a_option3_18 mixed_a_option3_17)(= mixed_a_option2_15 mixed_a_option2_14)(= mixed_a_option1_17 mixed_a_option1_15))(and (= mixed_a_option1_16 true)(= mixed_a_option3_16 false)(= mixed_a_option1_17 mixed_a_option1_16)(= mixed_a_option3_18 mixed_a_option3_16)(= mixed_a_option2_15 mixed_a_option2_13))))
(assert (ite (and (= mixed_a_option3_15 true) (not mixed_fl_active_0)) (and (= mixed_a_option2_16 mixed_a_option2_15) (= mixed_a_option3_19 mixed_a_option3_18) (= mixed_a_option1_18 mixed_a_option1_17)) (and (= mixed_a_option2_16 mixed_a_option2_13) (= mixed_a_option3_19 mixed_a_option3_15) (= mixed_a_option1_18 mixed_a_option1_15))))