CREATE TABLE users (
  Id INT AUTO_INCREMENT PRIMARY KEY ,
  Username varchar(50) UNIQUE ,
  FirstName varchar (50),
  LastName varchar (50),
  Email varchar (150) not null,
  Password varchar(255) not null,
  Phone varchar(30),
  UserStatus int not null default 0
) ENGINE=INNODB CHARSET=utf8;
