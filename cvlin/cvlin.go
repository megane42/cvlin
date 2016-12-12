package cvlin

func Run(rulePath, subjectPath string) (bool, error) {
	rule, err := LoadRule(rulePath)
	if err != nil {
		return false, err
	}

	subj, err := LoadSubject(subjectPath)
	if err != nil {
		return false, err
	}

	return validate(rule, subj)
}

func validate(rule map[string]rule, subject [][]string) (bool, error) {
	for _, row := range subject {
		// TODO
	}
	return false, nil
}
