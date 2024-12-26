# Sherpa
An environment storing and sharing tool

### DOING
- Authentication #back
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

### TODO 
- Save data #back
- Fetch data from file on selection #front
- Store previous nav data ? #front
- error handling /handlers/auth/getUserRepos #fix
- error handling /handlers/auth/AuthGithubLogin #fix
- handle cookie expiration /handlers/auth/generateUserCookie #fix
- handle empty columns /database/TokenFromCookie #fix
- handle cookie collision /database/GetUserOrCreateFromAuth

### DONE
- Select file #front
- Select repository #front
- Login view #front
- Main page #front

