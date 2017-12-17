create schema Keymanager;
use Keymanager;




create table KeyTable (keyId int(6) unsigned auto_increment primary key, name varchar(128) , size int , owner varchar(30) , 
									exportable bool , deletable bool );
                                    
create table operations ( id int(6)  unsigned auto_increment primary key , 
										encryptDecrypt bool , signVerify bool, hkdf bool , fpe bool , wrap bool);




alter table keyTable add column creationDate date , add column activationDate date , add state int ;


alter table keyTable add column deactivationDate date;


alter table keyTable modify column operations int(6)  unsigned;


alter table keyTable add constraint fk_operations foreign key (operations) references operations (id);


