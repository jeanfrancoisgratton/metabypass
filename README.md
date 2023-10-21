<H1>metabypass</H1>
___

A tool that allows you to publish news media articles on Meta platforms

<H2>How does it work</H2>

Suppose you have your own web server with a published DNS name, or any means to be remotely accessed. Say that you can reach at https://my-server.com/ <br>

Now, all web servers have a WEBROOT (or DOCUMENTROOT, etc). This is where, on the server the html files are hosted.<br>

In this example, say that it is `/var/www/html`<br>

All this tool does then is to create a html file that redirects to whichever remote site you wish to see published on those social media platforms that block other sites.<br><br>

All you need to, then is:
1. navigate to your precious web site, the one you'd wish to see published (say: https://my-news-site.com/my-article.html)<br>
2. now, you need to generate that new page; say that the new (local, generated) page will be named `my-redirect.html`<br>
3. so, you run: `screwUmeta generate /var/www/html/my-redirect.html https://my-news-site.com/my-article.html`
4. now, you publish that news article, on that social media platform with the following URL (according to our examples):<br>
https://my-server.com/my-redirect.html ; it will redirect to https://my-news-site.com/my-article.html