(set-logic QF_NRA)
(declare-fun choose1_a_choice_0 () Bool)
(declare-fun choose1_a_option1_0 () Bool)
(declare-fun choose1_a_option2_0 () Bool)
(declare-fun choose1_a_choice_1 () Bool)
(declare-fun choose1_a_option1_1 () Bool)
(declare-fun choose1_a_option2_1 () Bool)
(declare-fun choose1_a_option1_2 () Bool)
(declare-fun choose1_a_option2_2 () Bool)
(declare-fun choose1_a_option1_3 () Bool)
(declare-fun choose1_a_option2_3 () Bool)
(declare-fun choose1_a_option1_4 () Bool)
(declare-fun choose1_a_option2_4 () Bool)
(declare-fun choose1_a_option1_5 () Bool)
(declare-fun choose1_a_option1_6 () Bool)
(declare-fun choose1_a_option2_5 () Bool)
(declare-fun choose1_a_option2_6 () Bool)
(assert (= choose1_a_choice_0 false))
(assert (= choose1_a_option1_0 false))
(assert (= choose1_a_option2_0 false))
(assert (= choose1_a_choice_1 true))

(assert (or (and (and (= choose1_a_option1_1 true) (not (= choose1_a_option2_1 true)))
(= choose1_a_option1_3 choose1_a_option1_1)
(= choose1_a_option2_3 choose1_a_option2_1))
(and (and (not (= choose1_a_option1_2 true)) (= choose1_a_option2_2 true))
(= choose1_a_option2_3 choose1_a_option2_2)
(= choose1_a_option1_3 choose1_a_option1_2))))

(assert (ite (= choose1_a_choice_1 true) (and (= choose1_a_option1_4 choose1_a_option1_3) (= choose1_a_option2_4 choose1_a_option2_3)) (and (= choose1_a_option1_4 choose1_a_option1_0) (= choose1_a_option2_4 choose1_a_option2_0))))



(assert (= choose1_a_option1_5 true))

(assert (ite (= choose1_a_option1_4 true) (= choose1_a_option1_6 choose1_a_option1_5) (= choose1_a_option1_6 choose1_a_option1_4)))



(assert (= choose1_a_option2_5 true))

(assert (ite (= choose1_a_option2_4 true) (= choose1_a_option2_6 choose1_a_option2_5) (= choose1_a_option2_6 choose1_a_option2_4)))