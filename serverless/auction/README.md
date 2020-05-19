# Auction

## Setup

```shell
# Install deps
npm install

# Create .envrc to configure AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY_ID
touch .envrc
```

## Run function locally

```shell
npx sls invoke local -f <function-name>
```

## Deploy

```shell
npx sls deploy --stage prod

# Deploy a certain function
sls deploy -f createAuction -v

```
