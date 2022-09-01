create table FlagRecord(
no int not null AUTO_INCREMENT,
TeamID int not null,
MachineID int not null,
flag char(32) not null,
Round int not null,
PRIMARY KEY(no)
);
