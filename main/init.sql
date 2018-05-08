USE mysql;
DROP DATABASE IF EXISTS shadow;
DROP USER IF EXISTS shadow;
CREATE DATABASE shadow;
CREATE USER shadow
  IDENTIFIED BY 'shadow';
GRANT ALL ON shadow.* TO shadow@'%';

#
# For testing
#

# INSERT INTO user (email, token, token_renew, token_expire) VALUES ( "arief", "token", CURRENT_DATE(), CURRENT_DATE() );
# INSERT INTO user_role (email, role) VALUES ( "arief", "admin" );

# SELECT
#   u.email,
#   u.token,
#   u.token_renew,
#   u.token_expire,
#   r.role
# FROM user u LEFT JOIN user_role r on u.email = r.email
# WHERE token = '32ce79ea-1cf6-4885-52d8-91f0293ee112';