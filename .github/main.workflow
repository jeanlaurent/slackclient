workflow "New workflow" {
  on = "push"
  resolves = ["Push"]
}

action "Build" {
  uses = "actions/docker/cli@76ff57a"
  args = "build -t jeanlaurent/slack ."
}

action "Login to Hub" {
  uses = "actions/docker/login@76ff57a"
  needs = ["Build"]
  secrets = ["DOCKER_USERNAME", "DOCKER_PASSWORD"]
}

action "Push" {
  uses = "actions/docker/cli@76ff57a"
  args = "push jeanlaurent/slack"
  needs = ["Login to Hub"]
}
