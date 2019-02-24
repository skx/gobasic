# pushes
workflow "Push Event" {
  on = "push"
  resolves = ["Execute"]
}

# pull-requests
workflow "Pull Request" {
  on = "pull_request"
  resolves = ["Execute"]
}

# Run the magic
action "Execute" {
  uses = "skx/github-action-tester@master"
}
