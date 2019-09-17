# End to end testing

## Prerequisites

This project uses [endly](https://github.com/viant/endly/) end to end test runner.

1. Install latest [runner](https://github.com/viant/endly/releases) or use [endly docker image](https://github.com/viant/endly/tree/master/docker)
2. Create dedicated GCP project for  end to end testing.
3. Create e2e service account with admin permission on e2e test project
4. Generate and download [google secrets](https://github.com/viant/endly/tree/master/doc/secrets#gc) to ~/.secret/gcp-e2e.json
5. Create slack OAuth token and store in the ~/secret/slack-e2e.json in the following format:

```json
{
  "Token": "MY_OAUTH_SLACK_TOKEN"
}
```

If you do not have slack testing token just create a file with dummy data.

 
6. Checkout the this project:
```bash
git clone https://bqtail.git
cd bqtail/e2e
```
6. Update Bucket name in config file (it has to be globally unique)
[config/config.json](config/config.json)



## Use cases

To run all test use the following command:

```bash
endly run.yaml
```

To run individual use cases run first init task to upload configs, and deploy cloud functions, followed by individual case run.

```json
endly -t=init
```

- [Synchronous data files ingestionn](regression/cases/001_tail_nop/README.md)

```bash
    endly -t=test -i=tail_nop
```


- [Asynchronous data ingestion](regression/cases/002_tail_async/README.md)

```bash
    endly -t=test -i=tail_async
```

- [Synchronous in batch data ingestion](regression/cases/003_tail_batch_sync/README.md)

```bash
    endly -t=test -i=tail_batch_sync
```

- [Asynchronous in batch data ingestion](regression/cases/004_tail_batch_async/README.md)

```bash
    endly -t=test -i=tail_batch_async
```

- [Ingestion with transient dataset](regression/cases/005_tail_transient)

```bash
    endly -t=test -i=bqtail_transient
```

- [Ingestion with data deduplication](regression/cases/006_tail_batch_dedupe)

```bash
    endly -t=test -i=tail_batch_dedupe
```

- [Ingestion with nested data deduplication](regression/cases/007_tail_dedupe_nested)

```bash
    endly -t=test -i=tail_dedupe_nested
```


- [Export data on taret table modification](regression/cases/008_dispatch_export)

```bash
    endly -t=test -i=dispatch_export
```


- [Copy data on target table modification](regression/cases/009_dispatch_copy)

```bash
    endly -t=test -i=dispatch_copy
```



- [Copy data on target table modification](regression/cases/009_dispatch_copy)

```bash
    endly -t=test -i=dispatch_copy
```

- [Deduplication with summary task in sync mode](regression/cases/010_query_task)

```bash
    endly -t=test -i=query_task
```

- [Deduplication with summary task in async mode](regression/cases/011_async_query_task)

```bash
    endly -t=test -i=async_query_task
```


- [Error notification with slack](regression/cases/012_slack_notification)

```bash
    endly -t=test -i=slack_notification
```

