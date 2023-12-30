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

func (addresses *Addresses) ToString() []string {
	if addresses == nil {
		return nil
	}
	var result []string
	for _, address := range *addresses {
		result = append(result, address.ToString())
	}
	return result
}
