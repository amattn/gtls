[GoTutorial.net][] Link Shortener
=================================


[GoTutorial.net]: http://gotutorial.net


This is a sample project accompanying GoTutorial.net's web app lesson.  The slides that accompany this tutorial can be found at <http://gotutorial.net/golang_lessons.html>.

## Basic Tutorial Usage

	git clone https://github.com/amattn/gtls.git
	cd gtls

	git checkout tags/b01
	git checkout tags/b02
	git checkout tags/b03
	...

after any checkout you can:

	go test 
	go build && ./gtls


Database
--------

Install or update the driver with 

    go get -u github.com/lib/pq

You can install the test db with [macports](http://www.macports.org/install.php)

    # install postgres
    sudo port install postgresql95-server

    # setup db locally
    sudo mkdir -p /gtls/db
    sudo chown postgres:postgres /gtls/db
    sudo su postgres -c '/opt/local/lib/postgresql95/bin/initdb -D /gtls/db' 

In a different terminal window, start like so:

    sudo su postgres -c '/opt/local/lib/postgresql95/bin/pg_ctl -D /gtls/db -l /gtls/db/postgres.log start'

To uninstall simple remove the root gtls directory:

    sudo rm -Rf /gtls

To setup the db:

    sudo su postgres -c '/opt/local/lib/postgresql95/bin/createuser --superuser tutorial -U postgres'
    sudo su postgres -c "/opt/local/lib/postgresql95/bin/psql -c \"ALTER USER tutorial WITH PASSWORD 'changeme';\""
    sudo su postgres -c "/opt/local/lib/postgresql95/bin/psql -c \"CREATE TABLE links (code VARCHAR(64) NOT NULL,url TEXT NOT NULL,PRIMARY KEY(code));\""


