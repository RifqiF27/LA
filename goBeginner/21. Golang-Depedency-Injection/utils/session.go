package utils

func SessionRole() (int, string, bool) {
	session, err := ReadSession()
	if err != nil {
		SendJSONResponse("err", "Unauthorized", nil)
		return 0, "", false
	}

	role, ok := session["Role"].(string)
	userID, okID := session["ID"].(float64)
	if !ok || !okID {
		SendJSONResponse("err", "Forbidden", nil)
		return 0, "", false
	}

	return int(userID), role, true
}
