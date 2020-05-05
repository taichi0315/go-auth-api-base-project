INSERT INTO users(username) VALUES ("yaga");

INSERT INTO user_auths(user_id,email,hash) VALUES (1,"yaga@example.com","$2a$10$/z1SlkePlRKHfOOxb/w70.7B45svUsrqUq5kAFDUM/E4mjDMWvdwa");

INSERT INTO auth_tokens(user_id,token,expiry) VALUES (1,"token","2020-05-02 21:00:00");

INSERT INTO documents(url,user_id) VALUES
("https://go-tour-jp.appspot.com/welcome/1",1),
("https://soudai.hatenablog.com/entry/2018/05/01/204442",1),
("https://angular.jp/tutorial",1);
