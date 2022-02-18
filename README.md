# Refresh Token

CRUD User + Refresh Token Implementation

## Credential Admin

- Username : admin@admin.com
- Password : 1234qwer
- Note : Access token will be expired in 1 minute & Refresh token will be expired in 60 days

## API Documentation

- [CRUD User + Refresh Token](https://app.swaggerhub.com/apis-docs/furqonzt99/UserCRUD/1) - Swagger API Documentation

## Local Installation

- Clone this repo

```bash
git clone https://github.com/furqonzt99/refresh-token.git refresh-token
```

- Go to repository folder

```bash
cd refresh-token
go mod tidy
```

- Create .env file and add the following environment (you can see the variables from .env.example)

- Run this app

```bash
go run .
```

## Kubernetes Deploy

- Install Kubernetes https://kubernetes.io/docs/setup/

- Create Account & Setup Okteto https://www.okteto.com/docs/getting-started/

- Setup ENV Variables in kubernetes.yaml

- Run kubernetes.yaml

```bash
kubectl apply -f kubernetes.yaml
```

- Login to okteto cloud dashboard

![Okteto Deployment](https://github.com/furqonzt99/refresh-token/blob/main/documentation/okteto.png)

## Diagram Architecture

![Diagram Architecture](https://github.com/furqonzt99/refresh-token/blob/main/documentation/diagram.png)
