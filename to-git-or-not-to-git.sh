curl -s https://01.kood.tech/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Adeola\"}}){id}}"}'| jq -r '.data | .user | .[] | .id'
