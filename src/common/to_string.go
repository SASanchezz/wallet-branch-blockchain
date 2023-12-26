package common

func (hash *Hash) ToString() string {
	if hash == nil {
		return ""
	}
	return string(hash[:])
}

func (branchKey *BranchKey) ToString() string {
	if branchKey == nil {
		return ""
	}
	return string(branchKey[:])
}

func (address *Address) ToString() string {
	if address == nil {
		return ""
	}
	return string(address[:])
}
