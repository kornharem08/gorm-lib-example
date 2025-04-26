# GORM SQL Server Example

This project demonstrates how to use GORM with SQL Server in Go.

## Project Structure

- `cmd/` - Contains the main application
- `sqlwrap/` - SQL library for interacting with the database
- `docker-compose.yml` - Docker Compose configuration for SQL Server and the Go application
- `init-db.sql` - SQL script to initialize the database

## Running the Application

### Using Docker Compose

1. Make sure you have Docker and Docker Compose installed
2. Run the following command:

```bash
docker-compose up
```

This will:
- Start a SQL Server container
- Initialize the database with sample data
- Build and start the Go application

### Running Locally

1. Start SQL Server (you can use Docker):

```bash
docker-compose up db
```

2. Run the Go application:

```bash
cd cmd
go run main.go
```

## Connection String

The default connection string is:

```
jdbc:sqlserver://localhost:1433;databaseName=master;user=sa;password=YourStrong@Passw0rd;encrypt=true;trustServerCertificate=true

sqlserver://sa:YourStrong@Passw0rd@localhost:1433?database=master&encrypt=true&trustServerCertificate=true
```

You can override it by setting the `DB_CONNECTION_STRING` environment variable. 

```
IF OBJECT_ID('users', 'U') IS NOT NULL
    DROP TABLE users;

IF OBJECT_ID('companies', 'U') IS NOT NULL
    DROP TABLE companies;

-- Create the companies table
CREATE TABLE companies (
    CompanyCode NVARCHAR(50) PRIMARY KEY,
    Name NVARCHAR(255) COLLATE Thai_CI_AS,
    Address NVARCHAR(500) COLLATE Thai_CI_AS
);

-- Insert 10 mock records into companies
INSERT INTO companies (CompanyCode, Name, Address) VALUES
(N'COMP001', N'TechCorp Ltd.', N'123 Sukhumvit Road, Bangkok 10110'),
(N'COMP002', N'บริษัท เทคคอร์ป จำกัด', N'123 ถนนสุขุมวิท, กรุงเทพฯ 10110'),
(N'COMP003', N'Global Innovations Inc.', N'456 Silom Road, Bangkok 10500'),
(N'COMP004', N'บริษัท นวัตกรรมโลก จำกัด', N'456 ถนนสีลม, กรุงเทพฯ 10500'),
(N'COMP005', N'Skyline Solutions', N'789 Ratchadaphisek Road, Bangkok 10400'),
(N'COMP006', N'บริษัท สกายไลน์ โซลูชั่นส์ จำกัด', N'789 ถนนรัชดาภิเษก, กรุงเทพฯ 10400'),
(N'COMP007', N'FutureTech Enterprises', N'101 Rama 9 Road, Bangkok 10320'),
(N'COMP008', N'บริษัท ฟิวเจอร์เทค เอ็นเตอร์ไพรส์ จำกัด', N'101 ถนนพระราม 9, กรุงเทพฯ 10320'),
(N'COMP009', N'Alpha Dynamics', N'321 Phahonyothin Road, Bangkok 10400'),
(N'COMP010', N'บริษัท อัลฟา ไดนามิกส์ จำกัด', N'321 ถนนพหลโยธิน, กรุงเทพฯ 10400');

-- Create the users table with CompanyCode as a foreign key
CREATE TABLE users (
    Contactid NVARCHAR(255) PRIMARY KEY,
    ADUserId NVARCHAR(255),
    EmployeeId NVARCHAR(255),
    FirstName NVARCHAR(255),
    LastName NVARCHAR(255),
    LocalFirstName NVARCHAR(255) COLLATE Thai_CI_AS,
    LocalLastName NVARCHAR(255) COLLATE Thai_CI_AS,
    EmployeeTag NVARCHAR(50),
    CompanyCode NVARCHAR(50),
    CONSTRAINT FK_Users_Companies FOREIGN KEY (CompanyCode) REFERENCES companies(CompanyCode)
);

-- Insert 20 mock records into users with randomly assigned CompanyCode
INSERT INTO users (Contactid, ADUserId, EmployeeId, FirstName, LastName, LocalFirstName, LocalLastName, EmployeeTag, CompanyCode) VALUES
(N'C001', N'user001', N'EMP001', N'John', N'Smith', N'จอห์น', N'สมิธ', N'STAFF', N'COMP001'),
(N'C002', N'user002', N'EMP002', N'Emma', N'Johnson', N'เอ็มมา', N'จอห์นสัน', N'MANAGER', N'COMP001'),
(N'C003', N'user003', N'EMP003', N'Michael', N'Brown', N'ไมเคิล', N'บราวน์', N'STAFF', N'COMP002'),
(N'C004', N'user004', N'EMP004', N'Sarah', N'Davis', N'ซาราห์', N'เดวิส', N'CONTRACTOR', N'COMP002'),
(N'C005', N'user005', N'EMP005', N'David', N'Wilson', N'เดวิด', N'วิลสัน', N'STAFF', N'COMP003'),
(N'C006', N'user006', N'EMP006', N'Lisa', N'Taylor', N'ลิซา', N'เทย์เลอร์', N'MANAGER', N'COMP003'),
(N'C007', N'user007', N'EMP007', N'James', N'Anderson', N'เจมส์', N'แอนเดอร์สัน', N'STAFF', N'COMP004'),
(N'C008', N'user008', N'EMP008', N'Emily', N'Thomas', N'เอมิลี', N'โธมัส', N'CONTRACTOR', N'COMP004'),
(N'C009', N'user009', N'EMP009', N'Robert', N'Jackson', N'โรเบิร์ต', N'แจ็คสัน', N'STAFF', N'COMP005'),
(N'C010', N'user010', N'EMP010', N'Anna', N'White', N'แอนนา', N'ไวท์', N'MANAGER', N'COMP005'),
(N'C011', N'user011', N'EMP011', N'William', N'Harris', N'วิลเลียม', N'แฮร์ริส', N'STAFF', N'COMP006'),
(N'C012', N'user012', N'EMP012', N'Sophie', N'Martin', N'โซฟี', N'มาร์ติน', N'CONTRACTOR', N'COMP006'),
(N'C013', N'user013', N'EMP013', N'Thomas', N'Thompson', N'โธมัส', N'ทอมป์สัน', N'STAFF', N'COMP007'),
(N'C014', N'user014', N'EMP014', N'Olivia', N'Garcia', N'โอลิเวีย', N'การ์เซีย', N'MANAGER', N'COMP007'),
(N'C015', N'user015', N'EMP015', N'Charles', N'Martinez', N'ชาร์ลส', N'มาร์ติเนซ', N'STAFF', N'COMP008'),
(N'C016', N'user016', N'EMP016', N'Isabella', N'Robinson', N'อิซาเบลลา', N'โรบินสัน', N'CONTRACTOR', N'COMP008'),
(N'C017', N'user017', N'EMP017', N'Daniel', N'Clark', N'แดเนียล', N'คลาร์ก', N'STAFF', N'COMP009'),
(N'C018', N'user018', N'EMP018', N'Mia', N'Rodriguez', N'มีอา', N'โรดริเกซ', N'MANAGER', N'COMP009'),
(N'C019', N'user019', N'EMP019', N'Matthew', N'Lewis', N'แมทธิว', N'ลูออิส', N'STAFF', N'COMP010'),
(N'C020', N'user020', N'EMP020', N'Charlotte', N'Lee', N'ชาร์ล็อต', N'ลี', N'CONTRACTOR', N'COMP010');
```