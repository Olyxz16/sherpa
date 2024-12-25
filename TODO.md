Authenticate
Save data

Select repository
Fetch data from file on selection

Store previous nav data ?

handle error : handlers/auth/getUserRepos


AUTHENTICATION :
- select authentication method
- check account cookie
- add account
- if cookie
  - if already linked to account, refresh ?
  - if linked to other account, error
  - if not linked, link to current account
- if not cookie
  - if account already linked, get users cookie
  - else create new user account,
    link platform account,
    return cookie,
    prompt for account password
