# Job scrapper

Query indeed.com for jobs, generates data table with job information and links to your local machine.

Config file format


```yaml
baseurl   : https://www.indeed.com/jobs?
# Amount of jobs returned from each page.
baselimit : &limit=50
# Total amounts of jobs returned
maxresults : 200

jobs:
- job:
  keyword: JobKeyword1 JobKeyword2
  # You can use City ST, zipcode, remote
  # Location is optional
  location:
    - City ST
    - City ST
    - City ST
    - City ST
    - City ST

- job:
  keyword: golang
  location:
    - remote

- job:
  keyword: full stack developer
  location:
    - Austin TX
    - Denver CO
```
