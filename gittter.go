
type GitHubResponse struct {
	DefaultBranch string `json:"default_branch"`
}

func VerifyBranchName(username string, reponame string, branchName string) bool {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + reponame + "/branches" + branchName)
	if err != nil {
		check(err)
	}
	if response.StatusCode == http.StatusNotFound {
		return false
	}
	return true
}

func GetMainBranchName(username string, reponame string) (string, error) {
	response, err := http.Get("https://api.github.com/repos/" + username + "/" + reponame)
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(response.Body)
	// bodyStr := string(data)
	var obj GitHubResponse
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return "", err

	}

	return obj.DefaultBranch, nil
}

func InitRepo(path string) error {
	gitBin, _ := exec.LookPath("git")

	cmd := &exec.Cmd{
		Path:   gitBin,
		Args:   []string{gitBin, "init", path},
		Stdout: os.Stdout,
		Stdin:  os.Stdin,
	}

	err := cmd.Run()
	return err
}

func IsGitInstalled() bool {
	binPath, _ := exec.LookPath("git")
	if binPath != "" {
		return true
	}
	return false
}
