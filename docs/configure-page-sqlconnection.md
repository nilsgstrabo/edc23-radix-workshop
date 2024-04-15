# Configure application to read data from Azure SQL Database

Follow these steps to update our application's Index page to read data from the Movie table in our database using workload identity credentials:

## 1. Update `DB_CONNECTION_STRING` variable in radixconfig.yaml
- Open radixconfig.yaml.
- In the `DB_CONNECTION_STRING` variable, replace `<insert name here>` with the name (FQDN) of your Azure SQL Server, similar to this:
   ```yaml
   DB_CONNECTION_STRING: "Server=sql-edc23-radix-workshop-nst.database.windows.net; Database=moviedb; Authentication=Active Directory Managed Identity; Encrypt=True"
   ```
   You can find the name in the Azure Portal, or by running the following in the VS Code terminal:
   ```
   az sql server show --subscription d1775405-6d42-4fba-99ac-3cae223d9087 \
   --resource-group EDC23-Radix-Workshop \
   --name sql-edc23-radix-workshop-$(az ad signed-in-user show -otsv --query mail | awk -F"@" '{print tolower($1)}') \
   --query fullyQualifiedDomainName \
   -otsv
   ```
## 2. Change the Index page to connect to and read movies from the database
- Open `/src/web/Pages/Index.cshtml.cs`.
- Delete the existing `OnGet` method and replace it with the following `OnGetAsync` method:
   ```csharp
   public async Task OnGetAsync()
   {
      string connString = _config["DB_CONNECTION_STRING"] ?? "";

      using(var conn = new SqlConnection(connString)) {
         await conn.OpenAsync();
         Movies = (await conn.QueryAsync<Movie>("SELECT * FROM dbo.Movie")).ToList();
      }
   }
   ```
## 3. Commit and push the changes to your GitHub repo
   - Wait for the `build-deploy` job to finish.
   - Open the web page and verify that the application is able to connect to and read data from the database.

<br/>

---

<br/>

The .NET SQL Server driver has support for multiple Azure AD authentication methods, ref [Setting Azure Active Directory authentication](https://learn.microsoft.com/en-us/sql/connect/ado-net/sql/azure-active-directory-authentication?view=sql-server-ver16#setting-azure-active-directory-authentication).  
The `Authentication=Active Directory Managed Identity` setting in the connection string will try to connect using workload identity. If the runtime environment is not configured for this authentication method, like when running locally, the connection will fail. We can change the setting to `Authentication=Active Directory Default` to run locally, which will try to acquire credentials from various source, like VS Code/Visual Studio secrets and Azure CLI.
Try to add `"DB_CONNECTION_STRING": "Server=<your server name>; Database=moviedb; Authentication=Active Directory Default; Encrypt=True"` (replace `<your server name>` with the real server name) to the file `appsettings.Development.json` and run the application locally with the command `dotnet watch run --project src/web/edc23-radix-workshop-1.csproj`. The application will now connect to the database using your credentials from Azure CLI.