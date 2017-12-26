#!/usr/bin/env sh

function die() {
  echo "$@"
  exit 1
}


(which gx > /dev/null) || die "Please 'go get -u github.com/whyrusleeping/gx'"
(which gx-go > /dev/null) || die "Please 'go get -u github.com/whyrusleeping/gx-go'"
[[ $(git remote get-url upstream) == *"ipfs/go-ipfs" ]] || die "Please 'git remote add upstream github.com/ipfs/go-ipfs'"

echo "> Pulling upstream repo..."
git pull upstream master -X theirs
echo "> Installing gx dependencies..."
gx i
echo "> Rewritting import paths..."
gx-go uw
echo "> Replace ifps/go-ipfs to getlantern/go-ipfs..."
find . -name "*.go" | xargs perl -pi -e "s%github.com/ipfs/go-ipfs/%github.com/getlantern/go-ipfs/%g"
echo "> Committing changes..."
git commit -am "rewrite to non-ipfs import paths"


