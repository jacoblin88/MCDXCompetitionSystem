create table FlagTable(
TeamID int not null,
MachineID int not null,
flag char(32) not null,
Round int not null,
PRIMARY KEY(TeamID,MachineID)

);
