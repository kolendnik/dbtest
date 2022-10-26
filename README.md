# Database connection tester
tool for testing database connection

Usage:
- dbtest driver "DSN"
- dbtest sqlserver "sqlserver://{username}:{password}@{host}/{instance}?database={db}"

Avaible drivers with avaible DSN formats:
- sqlserver
	1. sqlserver://{username}:{password}@{host}/{instance}?database={db}
	2. sqlserver://{username}:{password}@{host}:{port}?database={db}
- mysql
	1. {username}:{password}@{protocol}({host})/{dbname}