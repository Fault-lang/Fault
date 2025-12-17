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
(declare-fun choose1_a_choice__state-%8_0 () Bool)
(declare-fun choose1_a_choice__state-%8_1 () Bool)
(declare-fun choose1_a_option1_4 () Bool)
(declare-fun choose1_a_option2_4 () Bool)
(declare-fun block162true_0 () Bool)
(declare-fun block162false_0 () Bool)
(declare-fun choose1_a_option1_5 () Bool)
(declare-fun choose1_a_option1_6 () Bool)
(declare-fun block166true_0 () Bool)
(declare-fun block166false_0 () Bool)
(declare-fun choose1_a_option2_5 () Bool)
(declare-fun choose1_a_option2_6 () Bool)
(declare-fun block170true_0 () Bool)
(declare-fun block170false_0 () Bool)
(assert (= choose1_a_choice_0 false))
(assert (= choose1_a_option1_0 false))
(assert (= choose1_a_option2_0 false))
(assert (= choose1_a_choice_1 true))

(assert (or (and choose1_a_choice__state-%8_0
(not choose1_a_choice__state-%8_1))
(and (not choose1_a_choice__state-%8_0)
choose1_a_choice__state-%8_1)))

(assert (ite (= choose1_a_choice_1 true) (= block162true_0 (and (= choose1_a_option1_4 choose1_a_option1_3) (= choose1_a_option2_4 choose1_a_option2_3))) (= block162false_0 (and (= choose1_a_option1_4 choose1_a_option1_0)
(= choose1_a_option2_4 choose1_a_option2_0)))))
(assert (or (and block162true_0
(not block162false_0))
(and (not block162true_0)
block162false_0)))


(assert (= choose1_a_option1_5 true))

(assert (ite (= choose1_a_option1_4 true) (= block166true_0 (= choose1_a_option1_6 choose1_a_option1_5)) (= block166false_0 (= choose1_a_option1_6 choose1_a_option1_4))))
(assert (or (and block166true_0
(not block166false_0))
(and (not block166true_0)
block166false_0)))


(assert (= choose1_a_option2_5 true))

(assert (ite (= choose1_a_option2_4 true) (= block170true_0 (= choose1_a_option2_6 choose1_a_option2_5)) (= block170false_0 (= choose1_a_option2_6 choose1_a_option2_4))))
(assert (or (and block170true_0
(not block170false_0))
(and (not block170true_0)
block170false_0)))
