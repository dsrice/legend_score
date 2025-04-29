# bowling_news
大会速報サービス

## CI/CD

### Backend Tests and Coverage

When a pull request is created or updated that affects files in the `be/` directory, GitHub Actions will automatically:

1. Run all backend tests
2. Generate a test coverage report
3. Comment on the PR with the total coverage percentage

This helps ensure code quality and provides visibility into test coverage for each change.

To run the tests locally:

```bash
cd be
./test.sh
```

This will generate a coverage report at `be/cover_file.html` that you can open in your browser.