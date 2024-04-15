# Workload Identity in Radix

In this workshop we will explore [Radix Workload Identity](https://www.radix.equinor.com/guides/workload-identity/), and how we can use it in our applications to access protected Azure resources without the use of long-lived client secrets or passwords.

More specifically, we will use the [Azure.SDK](https://azure.microsoft.com/en-us/downloads/) in a .NET web application to access blobs in an Azure Storage container using a managed identity, and connect to a Azure SQL database using the same managed identity.

[Brief introduction to how workload identity works](docs/workloadidentity_howitworks.md)



## Prerequisites

We recommend using VS Code with [Dev Containers](https://code.visualstudio.com/docs/devcontainers/containers) or [GitHub Codespaces](https://docs.github.com/en/codespaces/overview) for this workshop.

- Dev Containers is an [extension](https://code.visualstudio.com/docs/devcontainers/containers) to [VS Code](https://code.visualstudio.com/docs/setup/setup-overview) and requires [Docker Desktop](https://www.docker.com/products/docker-desktop/).
- GitHub Codespaces can be run in the browser or in VS Code with the the GitHub Codespaces [extension](https://marketplace.visualstudio.com/items?itemName=GitHub.codespaces).


## Steps

1. [Configure GitHub reposistory](docs/create_repository.md)
1. [Configure application in Radix Playground](docs/configure_radix_application.md)
1. [Configure Workload Identity](docs/configure-workload-identity.md)
1. [Configure Azure SQL Database](docs/configure-azure-sql.md)
1. [Configure Movie page to access SQL database](docs/configure-page-sqlconnection.md)
1. [Configure Movie page to access Storage Account](docs/configure-page-storageaccount.md)
1. [Optional: Golang SQL Server driver using Azure.Identity library](docs/golang-sql-server-driver.md)

