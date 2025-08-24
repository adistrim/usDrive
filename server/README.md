### Database migrations

Generating migrations might take some time.

Generate migrations
```bash
atlas migrate diff <migration-name> --env default
```

Apply migrations
```bash
atlas migrate apply --env default
```
