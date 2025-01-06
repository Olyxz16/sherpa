# Sherpa
An environment storing and sharing tool

### DOING
- HEAVY REFACTOR
  - db Scan on struct
  - define convention for masterkey and filekey
  - put abstraction over encoded values to avoid issues

### TODO 
- refactor getUserRepos
- Refactor db functions to avoid using cookies without multiple db calls
- Create error views #feat
- Save data #back
- Fetch data from file on selection #front
- error handling /handlers/auth/getUserRepos #fix
- error handling /handlers/auth/AuthGithubLogin #fix
- handle cookie expiration /handlers/auth/generateUserCookie #fix
- handle empty columns /database/TokenFromCookie #fix
- handle cookie collision /database/GetUserOrCreateFromAuth

### DONE
- Authentication
- Select file #front
- Select repository #front
- Login view #front
- Main page #front

### QUESTIONS
- Add file encryption key (derived from masterkey) in UserAuth ?
