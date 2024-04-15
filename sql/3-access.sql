CREATE ROLE MovieReader

GRANT SELECT ON dbo.Movie to MovieReader

-- Create a SQL USER from your managed identity.
-- Replace <your user name> with your actual user name (e.g. nst).
CREATE USER [id-edc23-radix-workshop-<your user name>] FROM EXTERNAL PROVIDER

-- Replace <your user name> with your actual user name (e.g. nst).
ALTER ROLE MovieReader ADD MEMBER [id-edc23-radix-workshop-<your user name>]