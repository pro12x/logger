message=$1
version=$2

git add .
git commit -m "$message"
git push origin main

git tag $version
git push origin $version