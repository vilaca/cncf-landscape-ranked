yq '.landscape.[].subcategories.[].items.[].repo_url' < $1 | grep -Fxv null \
    | uniq \
    | cut -d '/' -f4-5 \
    | while read repo; do curl -s -L "https://api.github.com/repos/$repo"  --header "Authorization: Bearer $PAT" \
    | jq -r '"\(.stargazers_count)|\(.full_name)|\(.license.name)|\(.description)"'; done \
    > $2
