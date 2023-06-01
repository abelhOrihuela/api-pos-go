server {
       listen 81;
       listen [::]:81;

       server_name *.mialmacen.shop;

       root /var/www/tutorial/;
       index index.html;

       location / {
               try_files $uri $uri/ =404;
       }

    listen [::]:443 ssl ipv6only=on; # managed by Certbot
    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/mialmacen.shop/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/mialmacen.shop/privkey.pem; # managed by Certbot
    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot

}