## :warning: Please read these instructions carefully and entirely first
* Clone this repository to your local machine.
* Use your IDE of choice to complete the assignment.
* When you have completed the assignment, you need to  push your code to this repository and [mark the assignment as completed by clicking here](https://app.snapcode.review/submission_links/0b0cdb7e-b3fd-431c-8a99-40e0a1128ceb).
* Once you mark it as completed, your access to this repository will be revoked. Please make sure that you have completed the assignment and pushed all code from your local machine to this repository before you click the link.

## Operability Take-Home Exercise

This project is a solution to the Operability Take-Home Exercise. It involves building an HTTP web server API that interacts with the GitHub API to fetch a list of publicly available Gists for a given user. The solution is packaged into a Docker container and includes automated tests.

### Running

If `docker-compose` is installed, just run:

```
docker-compose up
```

If not:
```
docker build -t equal-experts-test .
docker run -it --rm -p 8080:8080 equal-experts-test
```

### TODOs

There are several enhancements I could have implemented but left out due to time constraints. Here is a non-exhaustive list:

- **Observability:** Implement middleware to add observability signals using OpenTelemetry: Traces, Metrics, and Logs.
- **External Caching:** Add an external implementation of the Cache interface using Redis or Memcached.
- **Expand Testing:** Currently, there is only a happy path test case, and it uses the actual GitHub API, which is not ideal. With more time, I would inject the HTTP client to mock the GitHub API response, covering error paths like 404 and 500. I would also add fixtures with JSON to avoid hitting the live GitHub API, increasing the speed and reliability of the tests.
- **Rate Limiting:** Since we are querying the GitHub API, there are rate limits on their side. Implementing rate limiting on our API might be beneficial, although the cache already helps reduce the number of external calls.
- **User Authentication:** Currently, the API does not use any authentication, which significantly limits the number of requests it can make to GitHub. Allowing users to provide their GitHub Token would address this issue.