workflow "Github Tests" {
  resolves = ["Execute"]
  on = "pull_request"
}

action "Execute" {
  uses = "skx/github-action-tester@master"
}
