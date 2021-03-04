# Job scrapper

Query indeed.com for jobs, returns results in csv format with links to each job posting.

Config file format


```yaml
baseurl   : https://www.indeed.com/jobs?
baselimit : &limit=50

jobs:
- job:
  keyword: q=JobKeyword1+JobKeyword2
  location:
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
- job:
  keyword: JOB
  location:
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
    - l=CITY%2C+ST
```
