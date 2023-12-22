package common

func (hash *Hash) ToString() string {
	if hash == nil {
		return ""
	}
	returnVal := string(hash[:])
	return returnVal
}

func (branchKey *BranchKey) ToString() string {
	if branchKey == nil {
		return ""
	}
	return string(branchKey[:])
}
