(set-logic QF_NRA)
(declare-fun adand_a_choice_0 () Bool)
(declare-fun adand_a_option1_0 () Bool)
(declare-fun adand_a_option2_0 () Bool)
(declare-fun adand_a_option3_0 () Bool)
(declare-fun adand_a_choice_1 () Bool)
(declare-fun adand_a_option1_1 () Bool)
(declare-fun adand_a_option2_1 () Bool)
(declare-fun adand_a_option3_1 () Bool)
(declare-fun adand_a_option3_2 () Bool)
(declare-fun adand_a_option1_2 () Bool)
(declare-fun adand_a_option2_2 () Bool)
(declare-fun block112true_0 () Bool)
(declare-fun block112false_0 () Bool)
(declare-fun adand_a_option1_3 () Bool)
(declare-fun adand_a_option1_4 () Bool)
(declare-fun block116true_0 () Bool)
(declare-fun block116false_0 () Bool)
(declare-fun adand_a_option2_3 () Bool)
(declare-fun adand_a_option2_4 () Bool)
(declare-fun block120true_0 () Bool)
(declare-fun block120false_0 () Bool)
(declare-fun adand_a_option3_3 () Bool)
(declare-fun adand_a_option3_4 () Bool)
(declare-fun block123true_0 () Bool)
(declare-fun block123false_0 () Bool)
(assert (= adand_a_choice_0 false))
(assert (= adand_a_option1_0 false))
(assert (= adand_a_option2_0 false))
(assert (= adand_a_option3_0 false))
(assert (= adand_a_choice_1 true))

(assert (and (= adand_a_option1_1 true) (= adand_a_option2_1 true) (= adand_a_option3_1 true)))

(assert (ite (= adand_a_choice_1 true) (= block112true_0 (and (= adand_a_option3_2 adand_a_option3_1) (= adand_a_option1_2 adand_a_option1_1) (= adand_a_option2_2 adand_a_option2_1))) (= block112false_0 (and (= adand_a_option2_2 adand_a_option2_0)
(= adand_a_option3_2 adand_a_option3_0)
(= adand_a_option1_2 adand_a_option1_0)))))
(assert (or (and block112true_0
(not block112false_0))
(and (not block112true_0)
block112false_0)))


(assert (= adand_a_option1_3 true))

(assert (ite (= adand_a_option1_2 true) (= block116true_0 (= adand_a_option1_4 adand_a_option1_3)) (= block116false_0 (= adand_a_option1_4 adand_a_option1_2))))
(assert (or (and block116true_0
(not block116false_0))
(and (not block116true_0)
block116false_0)))


(assert (= adand_a_option2_3 true))

(assert (ite (= adand_a_option2_2 true) (= block120true_0 (= adand_a_option2_4 adand_a_option2_3)) (= block120false_0 (= adand_a_option2_4 adand_a_option2_2))))
(assert (or (and block120true_0
(not block120false_0))
(and (not block120true_0)
block120false_0)))


(assert (= adand_a_option3_3 true))

(assert (ite (= adand_a_option3_2 true) (= block123true_0 (= adand_a_option3_4 adand_a_option3_3)) (= block123false_0 (= adand_a_option3_4 adand_a_option3_2))))
(assert (or (and block123true_0
(not block123false_0))
(and (not block123true_0)
block123false_0)))