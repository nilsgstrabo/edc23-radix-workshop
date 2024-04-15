# Golang SQL Server driver using WorkloadIdentityCredential as token provider

The Azure SQL Server driver for Golang does not have built-in support for workload identity authentication. It does however have support for authentication using access token, and it is up to the developer to provide the access token.

See example code in directory `/src/goweb`, file `sql.go`.

Replace `src: /src/web` in radixconfig.yaml with `src: /src/goweb`. Commit changes and wait for build-deploy job complete. Open the URL for `web` component and verify that list of movies is returned.

