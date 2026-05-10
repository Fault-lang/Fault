(set-logic QF_NRA)
(declare-fun adand_a_choice_0 () Bool)
(declare-fun adand_a_option1_0 () Bool)
(declare-fun adand_a_option2_0 () Bool)
(declare-fun adand_a_option3_0 () Bool)
(declare-fun adand_a_choice_1 () Bool)
(declare-fun adand_a_option1_1 () Bool)
(declare-fun adand_a_option2_1 () Bool)
(declare-fun adand_a_option3_1 () Bool)
(declare-fun adand_a_option1_2 () Bool)
(declare-fun adand_a_option2_2 () Bool)
(declare-fun adand_a_option3_2 () Bool)
(declare-fun block52true_1 () Bool)
(declare-fun block52false_1 () Bool)
(declare-fun adand_a_option1_3 () Bool)
(declare-fun adand_a_option1_4 () Bool)
(declare-fun block56true_1 () Bool)
(declare-fun block56false_1 () Bool)
(declare-fun adand_a_option2_3 () Bool)
(declare-fun adand_a_option2_4 () Bool)
(declare-fun block60true_1 () Bool)
(declare-fun block60false_1 () Bool)
(declare-fun adand_a_option3_3 () Bool)
(declare-fun adand_a_option3_4 () Bool)
(declare-fun block63true_1 () Bool)
(declare-fun block63false_1 () Bool)
(assert (= adand_a_choice_0 false))
(assert (= adand_a_option1_0 false))
(assert (= adand_a_option2_0 false))
(assert (= adand_a_option3_0 false))
(assert (= adand_a_choice_1 true))

(assert (and (= adand_a_option1_1 true) (= adand_a_option2_1 true) (= adand_a_option3_1 true)))

(assert (ite (= adand_a_choice_1 true) (and (= block52true_1 true) (= block52false_1 false) (and (= adand_a_option1_2 adand_a_option1_1) (= adand_a_option2_2 adand_a_option2_1) (= adand_a_option3_2 adand_a_option3_1))) (and (= block52true_1 false) (= block52false_1 true) (and (= adand_a_option1_2 adand_a_option1_0)
(= adand_a_option2_2 adand_a_option2_0)
(= adand_a_option3_2 adand_a_option3_0)))))
(assert (or (and block52true_1
(not block52false_1))
(and (not block52true_1)
block52false_1)))


(assert (= adand_a_option1_3 true))

(assert (ite (= adand_a_option1_2 true) (and (= block56true_1 true) (= block56false_1 false) (= adand_a_option1_4 adand_a_option1_3)) (and (= block56true_1 false) (= block56false_1 true) (= adand_a_option1_4 adand_a_option1_2))))
(assert (or (and block56true_1
(not block56false_1))
(and (not block56true_1)
block56false_1)))


(assert (= adand_a_option2_3 true))

(assert (ite (= adand_a_option2_2 true) (and (= block60true_1 true) (= block60false_1 false) (= adand_a_option2_4 adand_a_option2_3)) (and (= block60true_1 false) (= block60false_1 true) (= adand_a_option2_4 adand_a_option2_2))))
(assert (or (and block60true_1
(not block60false_1))
(and (not block60true_1)
block60false_1)))


(assert (= adand_a_option3_3 true))

(assert (ite (= adand_a_option3_2 true) (and (= block63true_1 true) (= block63false_1 false) (= adand_a_option3_4 adand_a_option3_3)) (and (= block63true_1 false) (= block63false_1 true) (= adand_a_option3_4 adand_a_option3_2))))
(assert (or (and block63true_1
(not block63false_1))
(and (not block63true_1)
block63false_1)))
