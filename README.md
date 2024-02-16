## Simple Job Portal

The Simple Job Portal is an API-based platform designed to connect job seekers (talents)

### Talent Flow:

Talents can view available jobs.
Talents can apply for jobs.
Talents can track the status of their applications.

### Employer Flow:

Employers can post job listings, including job details and requirements.
Employers can view applications submitted for their posted jobs.
Employers can process applications by reviewing applicant details and making hiring
decisions. (Interview, Accept, Reject)

### Security:

Authentication: Using JWT tokens for user authentication and authorization. Only authorized users can access specific endpoints and perform relevant actions.
JWT tokens generated every time user login and stored inside cookies. It will be expired after 24 hours.

### Notes:

- Use this setting to create DB `host=localhost user=postgres password=123 dbname=bosshire port=5432 sslmode=disable`
- Or you can edit it at `database/postgres.go` file
- JWT secret is `secret`

### Testing method:

- Go to root folder
- Run it with `go run .`

### Documentation:

- After application started, open `localhost:8000/swagger/index.html` to access the documentation
- Can try the API inside swagger as well.

### TODO:

- Create config file to store environment parameters
- Add unit tests
