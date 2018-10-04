package main

import (
	"encoding/json"
	"fmt"
)

type GitHubRepo struct {
	RepoID   int64  `json:"id"`
	RepoURL  string `json:"html_url"`
	RepoName string `json:"full_name"`
	Language string `json:"language"`
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	b := []byte(`{"id":37479167,"node_id":"MDEwOlJlcG9zaXRvcnkzNzQ3OTE2Nw==","name":"sssd","full_name":"SSSD/sssd","private":false,"owner":{"login":"SSSD","id":12898906,"node_id":"MDEyOk9yZ2FuaXphdGlvbjEyODk4OTA2","avatar_url":"https://avatars3.githubusercontent.com/u/12898906?v=4","gravatar_id":"","url":"https://api.github.com/users/SSSD","html_url":"https://github.com/SSSD","followers_url":"https://api.github.com/users/SSSD/followers","following_url":"https://api.github.com/users/SSSD/following{/other_user}","gists_url":"https://api.github.com/users/SSSD/gists{/gist_id}","starred_url":"https://api.github.com/users/SSSD/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/SSSD/subscriptions","organizations_url":"https://api.github.com/users/SSSD/orgs","repos_url":"https://api.github.com/users/SSSD/repos","events_url":"https://api.github.com/users/SSSD/events{/privacy}","received_events_url":"https://api.github.com/users/SSSD/received_events","type":"Organization","site_admin":false},"html_url":"https://github.com/SSSD/sssd","description":"A daemon to manage identity, authentication and authorization for centrally-managed systems.","fork":false,"url":"https://api.github.com/repos/SSSD/sssd","forks_url":"https://api.github.com/repos/SSSD/sssd/forks","keys_url":"https://api.github.com/repos/SSSD/sssd/keys{/key_id}","collaborators_url":"https://api.github.com/repos/SSSD/sssd/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/SSSD/sssd/teams","hooks_url":"https://api.github.com/repos/SSSD/sssd/hooks","issue_events_url":"https://api.github.com/repos/SSSD/sssd/issues/events{/number}","events_url":"https://api.github.com/repos/SSSD/sssd/events","assignees_url":"https://api.github.com/repos/SSSD/sssd/assignees{/user}","branches_url":"https://api.github.com/repos/SSSD/sssd/branches{/branch}","tags_url":"https://api.github.com/repos/SSSD/sssd/tags","blobs_url":"https://api.github.com/repos/SSSD/sssd/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/SSSD/sssd/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/SSSD/sssd/git/refs{/sha}","trees_url":"https://api.github.com/repos/SSSD/sssd/git/trees{/sha}","statuses_url":"https://api.github.com/repos/SSSD/sssd/statuses/{sha}","languages_url":"https://api.github.com/repos/SSSD/sssd/languages","stargazers_url":"https://api.github.com/repos/SSSD/sssd/stargazers","contributors_url":"https://api.github.com/repos/SSSD/sssd/contributors","subscribers_url":"https://api.github.com/repos/SSSD/sssd/subscribers","subscription_url":"https://api.github.com/repos/SSSD/sssd/subscription","commits_url":"https://api.github.com/repos/SSSD/sssd/commits{/sha}","git_commits_url":"https://api.github.com/repos/SSSD/sssd/git/commits{/sha}","comments_url":"https://api.github.com/repos/SSSD/sssd/comments{/number}","issue_comment_url":"https://api.github.com/repos/SSSD/sssd/issues/comments{/number}","contents_url":"https://api.github.com/repos/SSSD/sssd/contents/{+path}","compare_url":"https://api.github.com/repos/SSSD/sssd/compare/{base}...{head}","merges_url":"https://api.github.com/repos/SSSD/sssd/merges","archive_url":"https://api.github.com/repos/SSSD/sssd/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/SSSD/sssd/downloads","issues_url":"https://api.github.com/repos/SSSD/sssd/issues{/number}","pulls_url":"https://api.github.com/repos/SSSD/sssd/pulls{/number}","milestones_url":"https://api.github.com/repos/SSSD/sssd/milestones{/number}","notifications_url":"https://api.github.com/repos/SSSD/sssd/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/SSSD/sssd/labels{/name}","releases_url":"https://api.github.com/repos/SSSD/sssd/releases{/id}","deployments_url":"https://api.github.com/repos/SSSD/sssd/deployments","created_at":"2015-06-15T17:09:48Z","updated_at":"2018-10-03T10:51:11Z","pushed_at":"2018-10-04T09:48:46Z","git_url":"git://github.com/SSSD/sssd.git","ssh_url":"git@github.com:SSSD/sssd.git","clone_url":"https://github.com/SSSD/sssd.git","svn_url":"https://github.com/SSSD/sssd","homepage":"https://pagure.io/SSSD/sssd","size":40564,"stargazers_count":67,"watchers_count":67,"language":"C","has_issues":false,"has_projects":false,"has_downloads":true,"has_wiki":false,"has_pages":false,"forks_count":61,"mirror_url":null,"archived":false,"open_issues_count":21,"license":{"key":"gpl-3.0","name":"GNU General Public License v3.0","spdx_id":"GPL-3.0","url":"https://api.github.com/licenses/gpl-3.0","node_id":"MDc6TGljZW5zZTk="},"forks":61,"open_issues":21,"watchers":67,"default_branch":"master","score":83.650696}`)
	g := GitHubRepo{}
	json.Unmarshal(b, &g)
	fmt.Println(g)
}
